package main

import (
	"crawler/crawler/engine"
	"crawler/crawler/scheduler"
	"crawler/crawler/zhenai/parser"
	_ "golang.org/x/net/html"
	_ "golang.org/x/text/encoding/simplifiedchinese"
	_ "golang.org/x/text/transform"
)

//判断错误
func PrintErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	//	URL:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	e := engine.ConcurrentEngine{
		Scheduler:    &scheduler.SimpleScheduler{},
		WorkereCount: 100,
	}
	e.Run(engine.Request{
		URL:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	//parser.ParseCityList(html)
}
