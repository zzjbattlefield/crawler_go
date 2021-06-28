package engine

import (
	"crawler/crawler/fetcher"
	"log"
)

type SimpleEngine struct{}

func (s SimpleEngine) Run(seed ...Request) {
	var requests []Request
	// for _, r := range seed {
	// 	requests = append(requests, r)
	// }
	requests = append(requests, seed...)
	//使用for循环不断的从requests队列里取得url来执行
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parserResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parserResult.Requests...)
		for _, item := range parserResult.Item {
			log.Printf("got Item : %v", item)
		}
	}
}

//获取request并解析内容返回
func worker(r Request) (ParseResult, error) {
	//获取url网页内容
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("error:Fetching Url: %s ,error: %v", r.Url, err)
		return ParseResult{}, err
	}
	//通过指定的解析方法去解析url内容
	return r.ParserFunc(body), nil
}
