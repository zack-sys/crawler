package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"time"
)

var rarerLimiter = time.Tick(100 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rarerLimiter
	//resp, err := http.Get(url)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.75 Safari/537.36")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("httpstatus error is %s", resp.StatusCode)
	}
	// gbk -> utf-8
	//transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	//自动识别编码 转换为utf-8

	reader := bufio.NewReader(resp.Body)
	r2 := reader
	encoding := determineEncoding(reader)
	utf8_html := transform.NewReader(r2, encoding.NewDecoder())
	return ioutil.ReadAll(utf8_html)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, _ := r.Peek(1024)
	encoding, _, _ := charset.DetermineEncoding(bytes, "")
	return encoding
}
