package engine

import "fmt"

//https://www.bilibili.com/video/av24365381/?p=66
//并发爬虫

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	ReadyNotify
	Run()
}

type ReadyNotify interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	// in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, req := range seeds {
		e.Scheduler.Submit(req) //将request发送给scheduler的管道，即in管道
	}

	itemCount := 0
	//用一个for循环接收out收到的数据
	for {
		result := <-out
		for _, item := range result.Items {
			itemCount++
			fmt.Println(itemCount, item)
		}

		for _, req := range result.Requests {
			e.Scheduler.Submit(req)
		}
	}
}

//
func createWorker(in chan Request, out chan ParseResult, r ReadyNotify) {
	go func() {
		for {
			r.WorkerReady(in) //worker告诉scheduler，他已经就绪了
			req := <-in
			parseResult, err := work(req)
			if err != nil {
				continue
			}
			out <- parseResult //处理完毕，向输出管道发送result
		}
	}()

}
