package main

import (
	"GoStudy/Chapter17/rpcsupport"
	"GoStudy/Chapter17/worker"
	"flag"
	"log"
)
// 创建启动配置参数,第三个参数就是我们的解析参数解析语句,使用go run xxxx.go --help即可查询
var port = flag.Int("port", 0, "The port for me to Listen on")

func main() {
	//log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{}))
	log.Fatal(rpcsupport.ServeRpc(":9000", worker.CrawlService{}))
}
