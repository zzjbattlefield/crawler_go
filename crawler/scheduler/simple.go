package scheduler

import "crawler/crawler/engine"

//调度器
type SimpleScheduler struct {
	workerChan chan engine.Request
}

//直接将收到的url放入channel
func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		s.workerChan <- request
	}()
}

func (s *SimpleScheduler) WorkerChannel() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}
