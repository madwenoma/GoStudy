package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, req := range seeds {
		e.Scheduler.Submit(req)//将request发送给scheduler的管道，即in管道
	}

	for {
		result := <-out
		for _, item := range result.Items {
			fmt.Println(item)
		}

		for _, req := range result.Requests {
			e.Scheduler.Submit(req)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult)  {
	go func() {
		for {
			req := <- in//接受scheduler in管道发来的request
			parseResult,err := worker(req)
			if err != nil {
				continue
			}
			out <- parseResult//处理完毕，向输出管道发送result
		}
	}()

}
