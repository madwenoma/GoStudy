package engine

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type Request struct {
	Url    string
	Parser Parser
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Id      string
	Url     string
	Type    string
	Payload interface{}
}

type NilParser struct{}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

//

type ParserFunc func(contents []byte, url string) ParseResult

//对ParserFunc的包装 实现了Parser
//适用于只有name参数的地方，如city和citylist
///profile不适用，所以自己定义，具体见profile.go
type FuncParser struct {
	parser ParserFunc
	name   string
}

func (p *FuncParser) Parse(contents []byte, url string) ParseResult {
	return p.parser(contents, url)
}

func (p *FuncParser) Serialize() (name string, args interface{}) {
	return p.name, nil
}

//工厂模式
func NewFuncParser(p ParserFunc, n string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   n,
	}
}
