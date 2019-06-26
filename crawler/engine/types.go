package engine

type ParserFunc func(contents []byte, url string) ParserResult

type Request struct {
	Url string
	ParserFunc func([]byte, string) ParserResult
}

type ParserResult struct {
	Requests []Request
	Items 	[]Item
}

func NilParse([]byte) ParserResult {
	return ParserResult{}
}

//单独创建由parser赋予的字段
type Item struct {
	Url 	string//查询到内容后，可以去访问
	Id 	 	string//去重不完善，多次运行会有重复的人出现，分配不同的Id
	Type 	string//类似于database的表名
	Payload interface{}
}

