package parser

import (
	"crawler/crawler/engine"
	"regexp"
)

const cityListRes = `<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^</a>]*)</a>`

//const cityListRes  = `<a href="http://album.zhenai.com/u/[0-9]+" target="_blank">心给懂的人</a>`

//解析某个城市的....解析出每个用户
func ParseCity(contents []byte) engine.ParseResult {
	//  [^>] 只到遇见 > 就结束 * 0个或多个  + 一个或多个
	reg := cityListRes
	re := regexp.MustCompile(reg)
	submatch := re.FindAllSubmatch(contents, -1)
	//fmt.Printf("%s\n",submatch)
	//fmt.Println(len(submatch))
	result := engine.ParseResult{}

	for _, s := range submatch {
		//result.Items = append(result.Items, "User: "+string(s[2]))
		result.Requests = append(result.Requests, engine.Request{
			URL:        string(s[1]),
			ParserFunc: ParseProfile,
		})
	}
	return result
}
