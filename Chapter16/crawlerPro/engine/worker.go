package engine

import (
	"log"
	"GoStudy/Chapter16/crawlerPro/fetcher"
)

func work(r Request) (ParseResult, error) {
	contents, err := fetcher.Fetcher(r.Url)
	if err != nil {
		log.Printf("Fetcher error fetching url %s %v", r.Url, err)
		return ParseResult{}, nil
	}
	parseResult := r.ParserFunc(contents)
	return parseResult, nil
}