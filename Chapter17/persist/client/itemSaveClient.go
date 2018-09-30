package client

import (
	"log"
	"GoStudy/Chapter17/crawlerPro/engine"
	"GoStudy/Chapter17/rpcsupport"
)

func ItemSave(host string) (chan engine.Item, error) {
	rpcClient, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("saver recieved:%d %v \n", itemCount, item)
			var result string
			err = rpcClient.Call("ItemSaveService.Save", item, &result)
			if err != nil {
				log.Printf("Item Save error:saving item %v,error:%v", item, err)
			}
			itemCount++
		}
	}()
	return out, err
}
