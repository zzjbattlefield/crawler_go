package scheduler

import "crawler/crawler/engine"

type QueueScheduler struct {
	ReqeustChan chan engine.Request      //接受外部requests的Channel
	WorkerChan  chan chan engine.Request //接收外部worker的channel
}

func (s *QueueScheduler) Submit(r engine.Request) {
	s.ReqeustChan <- r
}

//将worker发送到调度器内部的workerChannel 之后会将它加入到woker队列里
func (s *QueueScheduler) WorkerReady(w chan engine.Request) {
	s.WorkerChan <- w
}

func (s *QueueScheduler) ConfiguereMasterWorkerChan(chan engine.Request) {

}

//启动队列调度器
func (s *QueueScheduler) Run() {
	//启动时创建好调度器内部的channel
	s.WorkerChan = make(chan chan engine.Request)
	s.ReqeustChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activityRequest engine.Request
			var activityWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				//当两个队列都不为空的时候从队列获取值
				activityRequest = requestQ[0]
				activityWorker = workerQ[0]
			}
			select {
			case r := <-s.ReqeustChan:
				requestQ = append(requestQ, r)
			case w := <-s.WorkerChan:
				workerQ = append(workerQ, w)
			case activityWorker <- activityRequest:
				//当收到request时才从队列里去除值
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
