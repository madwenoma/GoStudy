package main

import (
	"GoStudy/Chapter15/crawlerPro/engine"
	"GoStudy/Chapter15/crawlerPro/zhenai/parser"
)

func main() {
	req := engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	}
	//engine.Run(req)

	e := engine.ConcurrentEngine{
		Scheduler:   &engine.SimpleScheduler{},
		WorkerCount: 20,
	}
	e.Run(req)
}
