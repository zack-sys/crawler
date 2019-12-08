package engine

import (
	"context"
	"crawler/crawler/model"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"testing"
)

func Test_save(t *testing.T) {
	// {半路夫妻 1535762338 34 高中及以下 阿坝 离异 165cm 3001-5000元 3001 5000}
	profile := model.Profile{
		Name:      "半路夫妻",
		Id:        "1535762338997",
		Age:       34,
		Education: "高中及以下",
		Local:     "阿坝",
		Marriage:  "离异",
		Height:    "165cm",
		Wage:      "3001-5000元",
		WageMin:   3001,
		WageMax:   5000,
	}
	save(profile)

}
func Test_get(t *testing.T) {
	client, e := elastic.NewClient(elastic.SetSniff(false))
	if e != nil {
		panic(e)
	}
	result, e := client.Get().
		Index("dating_profile").
		Type("zhenai").
		Id("1535762338").
		Do(context.Background())
	if e != nil {
		panic(e)
	}
	fmt.Printf("%s\n", result.Source)

	var pro model.Profile
	e = json.Unmarshal(result.Source, &pro)
	if e != nil {
		panic(e)
	}
	fmt.Printf("%+v\n", pro)

}
