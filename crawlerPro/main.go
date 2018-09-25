package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("err http, %d", resp.StatusCode)
		return
	}

	// utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("content is %s", all)
	printCityList(all)
}

func printCityList(contents []byte) {
	// <a href="http://www.zhenai.com/zhenghun/zhanjiang class="">湛江</a>
	//^表示not ，[^>]表示not左括号，比如：[^>]*一直匹配到>就停止
	cityReg := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`)
	matchCity := cityReg.FindAllSubmatch(contents, -1)
	for _, c := range matchCity {
		// for _, v := range c {
		// 	fmt.Printf("city:%s", v)
		// }

		fmt.Printf("cityName:%s,URL:%s\n", c[2], c[1])
		fmt.Println()
	}
	fmt.Println(len(matchCity))
}

func determineEncoding(reader io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(reader).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
