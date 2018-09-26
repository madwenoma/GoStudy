package main

import (
	"fmt"
	"GoStudy/retriever/mock"
	"GoStudy/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "imooc",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
	//func xxx(){}//组合接口里也可以定义其他函数
}

const url = "http://www.baidu.com"

func session(rp RetrieverPoster) string {
	rp.Post(url, map[string]string{"contents": "another faked baidu.com"})
	return rp.Get(url)
}

func main() {
	var r Retriever
	//等号右侧生成 copy给r
	r = mock.Retriever{"content of baidu"}
	fmt.Println(download(r))

	inspect(r)
	//&表示是指针给r 上面则是copy给r
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		Timeout:   time.Minute,
	}
	//fmt.Println(download(r))
	inspect(r)
	//type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.UserAgent)

	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Content)
	} else {
		fmt.Println("not a mock retriever")
	}

	rp := &mock.Retriever{"some content"}
	session(rp)
	fmt.Println(rp.Content)
}

func inspect(r Retriever) {
	fmt.Println("type switch")
	fmt.Printf("%T %v\n", r, r) //type value

	switch v := r.(type) { //检测变量类型
	case mock.Retriever:
		fmt.Println("Content:", v.Content)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
