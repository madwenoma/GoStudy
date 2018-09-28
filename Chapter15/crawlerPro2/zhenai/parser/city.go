package parser

import (
	"GoStudy/Chapter15/crawlerPro2/engine"
	"regexp"
)

//<a href="http://album.zhenai.com/u/1995815593" target="_blank">小顺儿</a>
var (
	userInfoReg = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlReg  = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult {
	matchCity := userInfoReg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, c := range matchCity {
		name := string(c[2]) //必须提出来 原因？
		// result.Items = append(result.Items, name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(c[1]),
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, name)
			},
		})
	}
	//
	matchs := cityUrlReg.FindAllSubmatch(contents, -1)
	for _, c := range matchs {
		url := string(c[1])
		result.Requests = append(result.Requests, engine.Request{
			Url:        url,
			ParserFunc: ParseCity,
		})
	}
	return result
}
