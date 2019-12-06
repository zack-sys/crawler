package engine

import (
	"crawler/crawler/fetcher"
	"log"
)

func Run(seeds ...Request){
	var requests []Request
	for _,r := range seeds{
		requests = append(requests, r)
	}

	for len(requests)>0 {
		r :=requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s",r.URL)
		html, err := fetcher.Fetch(r.URL)
		if err!=nil{
			log.Printf("Fetcher: error getching url %s : %v",r.URL,err)
			continue
		}
		parseResult := r.ParserFunc(html)
		// 吧一个切片加入另一个切片 ...
		requests = append(requests,parseResult.Requests...)

		for _,item := range parseResult.Items{
			log.Printf("Got item %s",item)
		}
	}
}
