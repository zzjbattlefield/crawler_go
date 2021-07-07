package parser

import (
	"crawler/crawler/engine"
	"regexp"
)

var profileRe = regexp.MustCompile(`<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
var cityUrlRe = regexp.MustCompile(`href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[^"]+)"`)

func ParserCity(content []byte, _ string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	for _, matche := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(matche[1]),
			ParserFunc: ProfileParser(string(matche[2])),
		})
	}

	matches = cityUrlRe.FindAllSubmatch(content, -1)
	for _, matche := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(string(matche[1])),
			ParserFunc: ParserCity,
		})
	}
	return result
}
