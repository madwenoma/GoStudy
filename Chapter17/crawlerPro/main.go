package main

import (
	"GoStudy/Chapter17/crawlerPro/engine"
	"GoStudy/Chapter17/crawlerPro/scheduler"
	"GoStudy/Chapter17/crawlerPro/zhenai/parser"

	itemSaverClient "GoStudy/Chapter17/persist/client"
	workerRpcClient "GoStudy/Chapter17/worker/client"
)

func main() {
	req := engine.Request{
		Url:    "http://www.zhenai.com/zhenghun/shanghai",
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	}
	//engine.Run(req)
	//itemChan, err := persist.ItemSaver("dating_profile")
	itemChan, err := itemSaverClient.ItemSave(":1234")
	if err != nil {
		panic(err)
	}

	//rpc worker
	wProcessor, err := workerRpcClient.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		// Scheduler: &scheduler.SimpleScheduler{},//想用simple要删除itemchan，16-1以后基本废弃
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 50,
		ItemChan:    itemChan,
		//RequestProcessor:engine.Work, //单机版
		RequestProcessor: wProcessor, //rpc版
	}
	e.Run(req)
}
