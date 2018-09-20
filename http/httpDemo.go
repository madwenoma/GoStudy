package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "https://www.imooc.com/", nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 8_0 like Mac OS X) AppleWebKit/600.1.3 (KHTML, like Gecko) Version/8.0 Mobile/12A4345d Safari/600.1.4")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(123)
	}
	defer resp.Body.Close()
	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(123)
	}
	fmt.Printf("%s", s)
}
