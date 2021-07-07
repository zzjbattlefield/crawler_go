package engine

import (
	"crawler/crawler/fetcher"
	"log"
)

//获取request并解析内容返回
func worker(r Request) (ParseResult, error) {
	//获取url网页内容
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("error:Fetching Url: %s ,error: %v", r.Url, err)
		return ParseResult{}, err
	}
	//通过指定的解析方法去解析url内容
	return r.ParserFunc(body, r.Url), nil
}
