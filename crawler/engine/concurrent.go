package engine

import "log"

//并发引擎
type ConCurrentEngien struct {
	Scheduler   Scheduler
	WorkerCount int
}

//调度器
type Scheduler interface {
	Submit(Request)
	ConfiguereMasterWorkerChan(chan Request)
}

func (e *ConCurrentEngien) Run(seeds ...Request) {
	//先创建并发需要的channel
	in := make(chan Request)
	out := make(chan ParseResult)
	//给Scheduler的channel赋值
	e.Scheduler.ConfiguereMasterWorkerChan(in)
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(in, out)
	}
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	count := 0
	for {
		result := <-out
		for _, gotItem := range result.Item {
			log.Printf("got Item #%d : %v", count, gotItem)
			count++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}

}

//创建worker并执行,通过channel返回ParserResult
func (e *ConCurrentEngien) createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
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
