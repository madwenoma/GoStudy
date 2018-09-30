package main

import (
	"net/http"
	"GoStudy/Chapter17/crawlerPro/frontend/controller"
)

func main() {
	//启动文件服务，使css和js生效
	http.Handle("/", http.FileServer(http.Dir("Chapter17/crawlerPro/frontend/view/")))

	http.Handle("/search", controller.CreateSearchResultHandler(
		"Chapter17/crawlerPro/frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
