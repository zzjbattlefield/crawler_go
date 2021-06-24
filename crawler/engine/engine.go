package engine

import (
	"crawler/crawler/fetcher"
	"log"
)

func Run(seed ...Request) {
	var requests []Request
	for _, r := range seed {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("error:Fetching Url: %s ,error: %v", r.Url, err)
			continue
		}
		parserResult := r.ParserFunc(body)
		requests = append(requests, parserResult.Requests...)
		for _, item := range parserResult.Item {
			log.Printf("got Item : %v", item)
		}
	}
}
