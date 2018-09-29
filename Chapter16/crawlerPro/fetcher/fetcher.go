package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
)

var rateLimiter = time.Tick(25 * time.Millisecond)

func Fetcher(url string) ([]byte, error) {
	//resp, err := http.Get(url)
	<-rateLimiter
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.80 Safari/537.36 Core/1.47.933.400 QQBrowser/9.4.8699.400")
	req.Close = true//改为短连接，防止出现EOF错误，但实际没有效果，不知道为何
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("wrong status code:%d", resp.StatusCode)
	}

	// utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

//自动检测网页编码
func determineEncoding(reader *bufio.Reader) encoding.Encoding {
	bytes, err := reader.Peek(1024) //接着用？
	if err != nil {
		if err != io.EOF { //产生EOF错误可能原因是，server端关闭了连接，而客户端复用的时候，发现server close了已经。
			log.Printf("determineEncoding error:%v", err)
		}
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "") //返回encoding.Encoding
	return e
}
