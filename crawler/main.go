package main

import (
	"crawler/crawler/engine"
	"crawler/crawler/persist"
	"crawler/crawler/scheduler"
	"crawler/crawler/zhenai/parser"
)

func main() {
	itemChannel, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConCurrentEngien{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
		SaveChan:    itemChannel,
	}
	e.Run(engine.Request{
		Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun/",
		ParserFunc: parser.ParserCity,
	})
}
