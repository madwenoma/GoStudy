package scheduler

import "GoStudy/Chapter15/crawlerPro2/engine"

type QueueScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request //chan也是一等公民 chan chan
}

func (s *QueueScheduler) Submit(req engine.Request) {
	s.requestChan <- req
}

func (s *QueueScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

//每个worker都make一个channel
func (s *QueueScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueueScheduler) Run() {
	s.requestChan = make(chan engine.Request)
	s.workerChan = make(chan chan engine.Request)
	go func() {
		var requestsQue []engine.Request
		var workerQue []chan engine.Request

		for {
			var activeReq engine.Request
			var activeWorker chan engine.Request
			if len(requestsQue) > 0 && len(workerQue) > 0 {
				activeReq = requestsQue[0]
				activeWorker = workerQue[0]
			}
			select {
			case r := <-s.requestChan:
				requestsQue = append(requestsQue, r)
			case w := <-s.workerChan:
				workerQue = append(workerQue, w)
			case activeWorker <- activeReq: //activeWorker是一个chan，所以可以将activeReq送进去
				requestsQue = requestsQue[1:]
				workerQue = workerQue[1:]
			}
		}
	}()
}
