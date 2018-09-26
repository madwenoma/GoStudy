package main

import (
	"GoStudy/crawlerPro/engine"
	"GoStudy/crawlerPro/zhenai/parser"
)

func main() {
	req := engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	}
	engine.Run(req)
}
