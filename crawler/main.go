package main

import (
	"crawler/crawler/engine"
	"crawler/crawler/scheduler"
	"crawler/crawler/zhenai/parser"
)

func main() {
	e := engine.ConCurrentEngien{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun/",
		ParserFunc: parser.ParseCityList,
	})
}
