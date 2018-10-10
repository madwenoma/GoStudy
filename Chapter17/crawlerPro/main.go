package main

import (
	"GoStudy/Chapter17/crawlerPro/engine"
	"GoStudy/Chapter17/crawlerPro/scheduler"
	"GoStudy/Chapter17/crawlerPro/zhenai/parser"

	itemSaverClient "GoStudy/Chapter17/persist/client"
	workerRpcClient "GoStudy/Chapter17/worker/client"
	"net/rpc"
	"GoStudy/Chapter17/rpcsupport"
	"log"
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

	hosts := []string{":9000"}
	pool := createProcessPool(hosts)

	//rpc worker
	wProcessor := workerRpcClient.CreateProcessor(pool)

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
func createProcessPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, host := range hosts {
		client, err := rpcsupport.NewClient(host)
		if err != nil {
			log.Printf("rpc new client err:%v", err)
		}
		clients = append(clients, client)
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
