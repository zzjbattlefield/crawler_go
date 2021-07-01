package fetcher

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var rateLimit = time.Tick(50 * time.Microsecond)

//通过url获取网页内容并返回
func Fetch(url string) ([]byte, error) {
	<-rateLimit
	rsp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()
	if rsp.StatusCode == http.StatusOK {
		// log.Println("fetchUrl:", url)
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
