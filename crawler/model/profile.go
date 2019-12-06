package model

//成都 | 36岁 | 大学本科 | 离异 | 170cm | 12001-20000元

type Profile struct {

	Name string

	Id string

	Age int
	//教育
	Education string
	//位置
	Local string
	//婚姻
	Marriage string
	//身高
	Height string
	//薪资
	Wage string
	WageMin int
	WageMax int
}

//func (p Profile)String() string{
//
//	return "{Id:"+p.Id+",Name:"+p.Name+",Age:"+strconv.Itoa(p.Age)+",Education:"+p.Education+
//		",Local:"+p.Local+",Marriage:"+p.Marriage+",Height:"+p.Height+",Wage:"+p.Wage+",WageMin:"+strconv.Itoa(p.WageMin)+
//		",WageMax:"+strconv.Itoa(p.WageMax)+"}"
//}
