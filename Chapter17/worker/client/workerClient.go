package client

import (
	"GoStudy/Chapter17/crawlerPro/engine"
	"GoStudy/Chapter17/rpcsupport"
	"fmt"
	"GoStudy/Chapter17/crawlerPro/config"
	rpcconfig "GoStudy/Chapter17/config"
	"GoStudy/Chapter17/worker"
)

//包装了rpc的调用，及序列化和调用成功后反序列化的过程
//golang不负责序列化，相当于只提供了dubbo的netty通信部分
func CreateProcessor() (engine.Processor, error) {
	//创建rpc client
	client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil {
		return nil, err
	}
	//返回接口实现（这里的实现，类似dubbo框架的部分）
	//1.将engine request序列化并传给rpc 2.调用rpc， 3.将结果反序列化
	return func(request engine.Request) (engine.ParseResult, error) {
		engineRequest := worker.SerializedRequest(request)
		var sResult worker.ParseResult //worker的request和parseResult相当于接口契约里的api
		err := client.Call(rpcconfig.CrawlServiceRpc, engineRequest, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}, nil
}
