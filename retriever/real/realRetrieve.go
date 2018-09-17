package real

import (
	"time"
	"net/http"
	"net/http/httputil"
	"fmt"
)

type Retriever struct {
	UserAgent string
	Timeout   time.Duration
}
//如果用指针，注意*是加在结构体上而不是变量上
//传入的时候应该是用&取得
func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url);
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()

	if err != nil {
		panic(err)
	}
	return string(result)
}
//实现Stringer接口 类似java重写toString
func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever <content is %s>", r.UserAgent)
}