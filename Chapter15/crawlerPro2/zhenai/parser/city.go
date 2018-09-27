package parser

import (
	"GoStudy/Chapter15/crawlerPro/engine"
	"regexp"
)

//<a href="http://album.zhenai.com/u/1995815593" target="_blank">小顺儿</a>
const userInfoReg = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	cityReq := regexp.MustCompile(userInfoReg)
	matchCity := cityReq.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, c := range matchCity {
		name := string(c[2])
		result.Items = append(result.Items, name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(c[1]),
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, name)
			},
		})

	}
	return result
}
