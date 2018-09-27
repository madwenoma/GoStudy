package scheduler

import "GoStudy/Chapter15/crawlerPro2/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request //所有的worker公用一个chan
}

func (ss *SimpleScheduler) WorkerChan() chan engine.Request {
	return ss.workerChan
}
func (ss *SimpleScheduler) WorkerReady(w chan engine.Request) {
}
func (ss *SimpleScheduler) Submit(req engine.Request) {
	go func() { ss.workerChan <- req }()
}

func (ss *SimpleScheduler) Run() {
	ss.workerChan = make(chan engine.Request)
}
