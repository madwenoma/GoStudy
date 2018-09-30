package engine

import (
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


