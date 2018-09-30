package main

import (
	"GoStudy/Chapter17/crawlerPro/engine"
	"GoStudy/Chapter17/crawlerPro/scheduler"
	"GoStudy/Chapter17/crawlerPro/zhenai/parser"

	"GoStudy/Chapter17/persist/client"
)

func main() {
	req := engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCityList,
	}
	//engine.Run(req)
	//itemChan, err := persist.ItemSaver("dating_profile")
	itemChan, err := client.ItemSave(":1234")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		// Scheduler: &scheduler.SimpleScheduler{},//想用simple要删除itemchan，16-1以后基本废弃
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 50,
		ItemChan:    itemChan,
	}
	e.Run(req)
}
