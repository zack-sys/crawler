package parser

import (
	"bufio"
	"crawler/crawler/fetcher"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)
func TestParseCityList(t *testing.T) {
	contents, e := ioutil.ReadFile("citylist_test_data.html")
	if e!=nil{
		panic(e)
	}
	//fmt.Printf("%s\n",contents)
	parseResult := ParseCityList(contents)
	//fmt.Println(parseResult)
	//for _,item := range parseResult.Items{
	//	log.Printf("Got item %s",item)
	//}

	const resultSize = 470
	if len(parseResult.Requests) != resultSize{
		t.Errorf(" Requests error resultSize = %d len(parseResult.Requests) = %d",resultSize,len(parseResult.Requests))
	}
	if len(parseResult.Items) != resultSize{
		t.Errorf(" Requests error resultSize = %d len(parseResult.Items) = %d",resultSize,len(parseResult.Items))
	}
}
func TestLoadParseCityList(t *testing.T) {
	contents, e := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	if e!=nil{
		panic(e)
	}
	//fmt.Printf("%s\n",contents)
	parseResult := ParseCityList(contents)

	for _,item := range parseResult.Items{
		log.Printf("Got item %s",item)
	}
	fmt.Println(len(parseResult.Requests))

	file, e := os.Create("citylist_test_data.html")
	if e!=nil{
		panic(e)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	fmt.Fprintln(writer,string(contents))
}