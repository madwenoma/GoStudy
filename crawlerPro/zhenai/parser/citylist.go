package parser

import (
	"GoStudy/crawlerPro/engine"
	"regexp"
)

/**
  使用正则表达式抽取城市名和URL
*/
// <a href="http://www.zhenai.com/zhenghun/zhanjiang class="">湛江</a>
//^表示not ，[^>]表示not左括号，比如：[^>]*一直匹配到>就停止
const cityReqStr = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParserResult {
	cityReq := regexp.MustCompile(cityReqStr)
	matchCity := cityReq.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, c := range matchCity {
		// for _, v := range c {
		// 	fmt.Printf("city:%s", v)
		// }
		result.Items = append(result.Items, c[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(c[1]),
			ParserFunc: engine.NilParser,
		})

		// fmt.Printf("cityName:%s,URL:%s\n", c[2], c[1]) //c[0]是整个匹配串
		// fmt.Println()
	}
	return result
	// fmt.Println(len(matchCity))
}
