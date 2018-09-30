package main

import (
	"GoStudy/Chapter17/rpcsupport"
	"GoStudy/Chapter17/persist"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func main() {
	//todo: 一般来说serveRpc方法是不会出错的，出错肯定是大错(系统出错)，所以可以直接错误退出就可以了

	// 常用的错误写法
	//err := serveRpc(":1234", "dating_profile")
	//if err != nil {
	//	panic(err)
	//}

	// 偷懒的错误写法
	//log.Fatal(serveRpc(fmt.Sprintf(":%d", *port),  config.ElasticIndex))
	log.Fatal(serveRpc(":1234", "dating_profile"))

}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetURL("http://100.100.16.55:9200"),
		elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host, &persist.ItemSaveService{
		Client: client,
		Index:  index,
	})
}
