package parser

import (
	"crawler/crawler/engine"
	"regexp"
)

const cityListRe  = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" [^>]*>([^<]+)</a>`

//解析城市列表
func ParseCityList(contents []byte) engine.ParseResult{
	//  [^>] 只到遇见 > 就结束 * 0个或多个  + 一个或多个
	reg := cityListRe
	re := regexp.MustCompile(reg)
	submatch := re.FindAllSubmatch(contents, -1)

	result:=engine.ParseResult{}

	for i,s := range submatch{
		result.Items = append(result.Items, "City: "+string(s[2]))
		result.Requests = append(result.Requests, engine.Request{
			URL:        string(s[1]),
			ParserFunc: ParseCity,
		})
		if i == 1{
			break
		}
	}
	return result
}