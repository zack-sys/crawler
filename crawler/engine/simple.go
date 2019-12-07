package engine

import (
	"crawler/crawler/fetcher"
	"log"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}

		// 吧一个切片加入另一个切片 ...
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %s", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.URL)
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("ParserFunc: panic getching url %s : %v", r.URL, err)
		}
	}()
	html, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("Fetcher: error getching url %s : %v", r.URL, err)
		return ParseResult{}, nil
	}
	parserFunc := r.ParserFunc(html)
	return parserFunc, nil
}
