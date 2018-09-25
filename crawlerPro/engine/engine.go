package engine

import (
	"GoStudy/crawlerPro/fetcher"
	"log"
)

func Run(seed ...Request) {
	var requests []Request
	for _, req := range seed {
		requests = append(requests, req)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		contents, err := fetcher.Fetcher(r.Url)
		if err != nil {
			log.Println()
			continue
		}
		parseResult := r.ParserFunc(contents)
		requests = append(requests, parseResult.Requests...)
	}

}
