package main

import (
	"fmt"
	"regexp"
)

const text  = "my email is asdasfas@qq.com@dasds"
const text2  = `my email 
is asdasfas@qq.com@dasds

email 2 dsa@a.com

Email d.d@ff.orgdas@..@@..
eda 2das@com.cn.cn.cn

`
func main() {
	// 带两个返回值得 err 为验证正则表达式是否有错误
	//compile, e := regexp.Compile()

	// 包装 Compile 一定正确 错误抛异常
	// "[A-Za-z0-9]+@qq.com"
	//  "" 双引号 如果用 点 \\.  单引号就需要一个 `\.`     [] 里面判断点 不需要双斜杠
	re := regexp.MustCompile("([A-Za-z0-9]+)@([A-Za-z0-9.]+)\\.([A-Za-z0-9]+)")
	s := re.FindString(text)
	fmt.Println(s)

	all := re.FindAllString(text2,-1)
	fmt.Println(all)

	submatch := re.FindAllStringSubmatch(text2, -1)
	fmt.Println(submatch)

}
