package parser

import (
	"crawler/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-zA-Z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(content []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllStringSubmatch(string(content), -1)
	result := engine.ParseResult{}
	for _, matche := range matches {
		//返回城市名
		result.Item = append(result.Item, matche[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:        matche[1],
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
