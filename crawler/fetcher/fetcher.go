package fetcher

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	rsp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()
	if rsp.StatusCode == http.StatusOK {
		log.Println("fetchUrl:", url)
		all, err := ioutil.ReadAll(rsp.Body)
		if err != nil {
			return nil, err
		}
		return all, nil
	} else {
		errMsg := fmt.Sprintf("httpStatusError:%d", rsp.StatusCode)
		return nil, errors.New(errMsg)
	}
}
