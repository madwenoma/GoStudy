package engine

import (
	"GoStudy/Chapter15/crawlerPro/fetcher"
	"log"
	"time"
)

//simple engine
//single thread
func Run(seed ...Request) {
	var requests []Request
	for _, req := range seed {
		requests = append(requests, req)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("fetching:%s", r.Url)
		parseResult, err := work(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
		time.Sleep(time.Millisecond * 1200)
	}
}

func work(r Request) (ParseResult, error) {
	contents, err := fetcher.Fetcher(r.Url)
	if err != nil {
		log.Printf("Fetcher error fetching url %s %v", r.Url, err)
		return ParseResult{}, nil
	}
	parseResult := r.ParserFunc(contents)
	return parseResult, nil
}
