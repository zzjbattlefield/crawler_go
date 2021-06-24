package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	content, err := ioutil.ReadFile("cityList_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(content)
	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("Url应该有:%d 实际有:%d", resultSize, len(result.Requests))
	}
	if len(result.Item) != resultSize {
		t.Errorf("Item应该有:%d 实际有:%d", resultSize, len(result.Item))
	}
}
