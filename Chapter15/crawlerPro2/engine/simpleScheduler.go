package engine

type SimpleScheduler struct {
	workerChan chan Request
}

func (ss *SimpleScheduler) ConfigureMasterWorkChan(req chan Request) {
	ss.workerChan = req
}

func (ss *SimpleScheduler) Submit(req Request) {
	go func() { ss.workerChan <- req }()
}