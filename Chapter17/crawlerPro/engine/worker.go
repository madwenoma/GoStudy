package engine

import (
	"log"
	"GoStudy/Chapter17/crawlerPro/fetcher"
)

func Work(r Request) (ParseResult, error) {
	contents, err := fetcher.Fetcher(r.Url)
	if err != nil {
		log.Printf("Fetcher error fetching url %s %v", r.Url, err)
		return ParseResult{}, nil
	}
	parseResult := r.Parser.Parse(contents, r.Url)
	return parseResult, nil
}
