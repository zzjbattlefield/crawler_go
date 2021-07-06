package persist

import (
	"context"
	"crawler/crawler/engine"
	"crawler/crawler/model"
	"encoding/json"
	"testing"

	"github.com/olivere/elastic"
)

func Test_save(t *testing.T) {
	profile := engine.Item{
		Url: "http://www.baidu.com",
		Id:  "11223344",
		PayLoad: model.Profile{
			Name:       "寂寞成影萌宝",
			Gender:     "女",
			Age:        83,
			Height:     105,
			Weight:     137,
			Income:     "财务自由",
			Marriage:   "离异",
			Education:  "初中",
			Occupation: "金融",
			Hokou:      "南京市",
			Xinzuo:     "狮子座",
			House:      "无房",
			Car:        "无车",
		},
	}
	err := save(profile)
	if err != nil {
		panic(err)
	}
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.58.130:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	rsp, err := client.Get().Index("dating_profile").Id(profile.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	var actual engine.Item
	err = json.Unmarshal(*rsp.Source, &actual)
	if err != nil {
		panic(err)
	}
	if profile != actual {
		t.Errorf("got %v expected %v", actual, profile)
	}
}
