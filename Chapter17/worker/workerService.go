package worker

import (
	"GoStudy/Chapter17/crawlerPro/engine"
	"log"
)

type CrawlService struct {
}

func (c CrawlService) Process(req Request, result *ParseResult) error {
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}
	//真正实现rpc server端的work逻辑
	log.Printf("rpc worker process:%s", req.Parser.Name)
	engineResult, err := engine.Work(engineReq)
	if err != nil {
		return err
	}
	*result = SerializeResult(engineResult)
	return nil
}
