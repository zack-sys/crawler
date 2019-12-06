package parser

import (
	"net/http"
	"net/http/httputil"
	"testing"
)

func TestParseCity(t *testing.T) {
	//resp, _ := http.Get("https://album.zhenai.com/u/1084658562")
	//request, _ := http.NewRequest(http.MethodGet, "https://album.zhenai.com/u/1084658562", nil)
	//request, _ := http.NewRequest(http.MethodGet, "http://www.zhenai.com/zhenghun/chengdu", nil)
	request, _ := http.NewRequest(http.MethodGet, "http://www.zhenai.com/zhenghun/aba", nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.75 Safari/537.36")
	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()

	bytes, e := httputil.DumpResponse(response, true)
	if e!=nil{
		panic(e)
	}
	//fmt.Printf("%s \n",bytes)
	_ = ParseCity(bytes)
}