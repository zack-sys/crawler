package parser

import (
	"crawler/crawler/engine"
	"crawler/crawler/model"
	"regexp"
	"strconv"
	"strings"
)

//const Re = `< .+ class="info" .+ \| 36岁 \| 大学本科 \| 离异 \| 170cm \| 12001-20000元</div> <div class="actions" `
//const Re = `<div class="des f-cl".+ [^</div>]+?</div>`
const Re = `<div class="des f-cl".+>(.+ [^\|</div>]*)</div>`
const ReName  = `<h1 class="nickName" .+[^>]+>(.+[^<]+)</h1>`
const ReId  = `<div class="id" .+[^>]+>ID：([0-9]+)</div>`
var re = regexp.MustCompile(Re)
var reName = regexp.MustCompile(ReName)
var reId = regexp.MustCompile(ReId)
func ParseProfile(contents []byte) engine.ParseResult{
	profile := model.Profile{}
	submatch := re.FindSubmatch(contents)
	submatchName := reName.FindSubmatch(contents)
	submatchId := reId.FindSubmatch(contents)
	//fmt.Printf("%s\n",submatch)
	//fmt.Printf("%s\n",submatchName)
	//fmt.Printf("%s\n",submatchId)
	//fmt.Println(string(submatchName[1]))
	//fmt.Println(string(submatchId[1]))
	//fmt.Println(string(submatch[1]))
	split := strings.Split(strings.ReplaceAll(string(submatch[1])," ",""), "|")
	profile.Name = string(submatchName[1])
	profile.Id = string(submatchId[1])
	profile.Local = split[0]
	profile.Age,_ = strconv.Atoi(strings.ReplaceAll(split[1],"岁",""))
	profile.Education = split[2]
	profile.Marriage = split[3]
	profile.Height = split[4]
	profile.Wage = split[5]
	all := strings.ReplaceAll(split[5], "元", "")
	i := strings.Split(all, "-")
	if len(i) == 2{
		profile.WageMin, _ = strconv.Atoi(i[0])
		profile.WageMax, _ = strconv.Atoi(i[1])
	}

	//for _,s:=range split{
	//	fmt.Println(s)
	//}
	//fmt.Println(profile)
	result := engine.ParseResult{
		Requests: nil,
		Items:    []interface{}{profile},
	}
	return result
}
