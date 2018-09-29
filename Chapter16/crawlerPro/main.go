package main

import (
	"GoStudy/Chapter16/crawlerPro/engine"
	"GoStudy/Chapter16/crawlerPro/persist"
	"GoStudy/Chapter16/crawlerPro/scheduler"
	"GoStudy/Chapter16/crawlerPro/zhenai/parser"
)

func main() {
	req := engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCityList,
	}
	//engine.Run(req)

	e := engine.ConcurrentEngine{
		// Scheduler: &scheduler.SimpleScheduler{},//想用simple要删除itemchan，16-1以后基本废弃
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 50,
		ItemChan:    persist.ItemSaver(),
	}
	e.Run(req)
}
