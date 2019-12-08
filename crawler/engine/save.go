package engine

import (
	"context"
	"crawler/crawler/model"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

func Save() chan interface{} {
	out := make(chan interface{})
	go func() {
		count := 0
		for {
			item := <-out
			log.Printf("%d Got item: %v", count, item)
			save(item)
			count++
		}
	}()
	return out
}
func save(item interface{}) {

	client, e := elastic.NewClient(elastic.SetSniff(false))
	if e != nil {
		panic(e)
	}
	response, e := client.Index().
		Index("dating_profile").
		Type("zhenai").
		Id(item.(model.Profile).Id).
		BodyJson(item).
		Do(context.Background())
	if e != nil {
		panic(e)
	}
	fmt.Printf("%+v", response)

	x := []model.Profile{}
	x = append(x)
}
