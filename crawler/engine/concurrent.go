package engine

import "crawler/crawler/model"

//并发引擎
type ConCurrentEngien struct {
	Scheduler   Scheduler
	WorkerCount int
	SaveChan    chan Item
}

//调度器
type Scheduler interface {
	Submit(Request)
	WorkerChannel() chan Request
	ReadyNotifier
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConCurrentEngien) Run(seeds ...Request) {
	//创建接收返回参数的channel
	out := make(chan ParseResult)
	//启动调度器
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChannel(), out, e.Scheduler)
	}
	for _, r := range seeds {
		//传入种子页面
		e.Scheduler.Submit(r)
	}
	for {
		result := <-out
		for _, gotItem := range result.Item {
			if _, ok := gotItem.PayLoad.(model.Profile); ok {
				go func() {
					e.SaveChan <- gotItem
				}()
			}
		}

		for _, request := range result.Requests {
			//url去重
			if isDuplicate(request) {
				// log.Printf("url重复: %s \n", request.Url)
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

//创建worker并执行,通过channel返回ParserResult
func (e *ConCurrentEngien) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	// in := make(chan Request)
	go func() {
		for {
			//将空闲的worker发送给Scheduler加入到worker队列
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				//worker里已经记录日志 如果有错误直接continue
				continue
			}
			out <- result
		}
	}()
}

type emptyStruct struct{}

var visitedUrls = make(map[string]emptyStruct)

//判断url是否重复
func isDuplicate(req Request) bool {
	if _, ok := visitedUrls[req.Url]; !ok {
		//没有见过 保存
		visitedUrls[req.Url] = emptyStruct{}
		return false
	} else {
		return true
	}
}
