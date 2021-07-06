package parser

import (
	"crawler/crawler/model"
	"io/ioutil"
	"os"
	"testing"
)

func TestParseProfile(t *testing.T) {
	handler, err := os.Open("profile_test_data.html")
	if err != nil {
		t.Error("fail to open test file: ", err)
	}
	content, err := ioutil.ReadAll(handler)
	if err != nil {
		t.Error("fail to read test file: ", err)
	}
	result := ParseProfile(content, "寂寞成影萌宝")
	if len(result.Item) != 1 {
		t.Errorf("Result 最少也有一个内容,但是有%v", result.Item)
	}
	profile := result.Item[0]

	expected := model.Profile{
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
	}
	if profile != expected {
		t.Errorf("expected %+v, but was %+v", expected, profile)
	}
}
