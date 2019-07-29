package engine

//分布式添加接口
type Parser interface {
	Parse(contents []byte, url string) ParserResult
	Serialize() (name string, args interface{})
}

type ParseFunc func(contents []byte, url string) ParserResult

type Request struct {
	Url string
	Parser Parser//更改为分布式，将ParseFunc换为Parser的interface
}

type ParserResult struct {
	Requests []Request
	Items 	[]Item
}

//单独创建由parser赋予的字段
type Item struct {
	Url 	string//查询到内容后，可以去访问
	Id 	 	string//去重不完善，多次运行会有重复的人出现，分配不同的Id
	Type 	string//类似于database的表名
	Payload interface{}
}

//分布式后更改nilParser的写法
type NilParser struct {
}

func (NilParser) Parse(_ []byte, _ string) ParserResult {
	return ParserResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

//我们有那么多Parse函数，将函数包装成Parser的类型
type FuncParser struct {
	parser ParseFunc
	name string //放一个函数的名字
}
//工厂函数FuncParser
func NewFuncParser(p ParseFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:  name,
	}
}

//让FuncParser去实现Parser这个Interface
func (f *FuncParser) Parse(contents []byte, url string) ParserResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}