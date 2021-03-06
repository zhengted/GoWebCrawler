package main

import (
	"GoWebCrawler/crawler/Scheduler"
	"GoWebCrawler/crawler/engine"
	"GoWebCrawler/crawler/persist"
	"GoWebCrawler/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &Scheduler.SimpleScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url:    "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})
	//engine.SimpleEngine{}.Run(engine.Request{
	//		Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
	//		ParserFunc: parser.ParseCityList,
	//	})
}

// 城市列表解析器Seed

// 城市解析器
// 输入：UTF-8文本
// 输出：request列表{URL和对应的parser} item列表

// 用户解析器
