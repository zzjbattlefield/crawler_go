package engine

type Request struct {
	//从解析器获得的下一个任务的url
	Url string
	//从解析器获得的下一个任务的解析器方法地址
	ParserFunc func(content []byte, url string) ParseResult
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}

//解析器解析的结果
type ParseResult struct {
	//任务列表
	Requests []Request
	//网页获取的数据列表
	Item []Item
}

type Item struct {
	Url     string
	Id      string
	PayLoad interface{} //任意字段
}
