package main

import (
	"fmt"
	"regexp"
)

const txt = `my email is kkk@gmail.com
this1234 email is : killgov@163.com  my second email is 28223231@sina.com.cn
last email someone@qq.com
`

func main() {
	// re, err := regexp.Compile("kkk@gmail.com")
	/*关于正则
	  .任何字符
	  *0个或多个
	  +1个或多个
	*/

	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`) //程序里写的，默认为一定是匹配的,不需要处理错误
	match := re.FindAllString(txt, -1)                                   //找到所有匹配正则的部分
	fmt.Println(match)

	re2 := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`) //程序里写的，默认为一定是匹配的,不需要处理错误
	matchAndSplit := re2.FindAllStringSubmatch(txt, -1)
	for _, v := range matchAndSplit {
		fmt.Println(v)
	}
	// fmt.Println(matchAndSplit)
}
