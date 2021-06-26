package parser

import (
	"crawler/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParserCity(content []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	for _, matche := range matches {
		result.Item = append(result.Item, "User:"+string(matche[2]))
		name := matche[2]
		result.Requests = append(result.Requests, engine.Request{
			Url: string(matche[1]),
			ParserFunc: func(b []byte) engine.ParseResult {
				return ParseProfile(b, string(name))
			},
		})
	}
	return result
}
