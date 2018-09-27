package main

import (
	"GoStudy/Chapter15/crawlerPro2/engine"
	"GoStudy/Chapter15/crawlerPro2/scheduler"
	"GoStudy/Chapter15/crawlerPro2/zhenai/parser"
)

func main() {
	req := engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	}
	//engine.Run(req)

	e := engine.ConcurrentEngine{
		// Scheduler: &scheduler.SimpleScheduler{},
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 20,
	}
	e.Run(req)
}
