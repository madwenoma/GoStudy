package main

import (
	"goStudy/crawlerPro/engine"
	"goStudy/crawlerPro/zhenai/parser"
)

func main() {
	req := engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	}
	engine.Run(req)
}
