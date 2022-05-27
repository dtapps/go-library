package ip2region

import (
	"io/ioutil"
	"net/http"
)

func getOnline() ([]byte, error) {
	resp, err := http.Get("https://ghproxy.com/?q=https://github.com/lionsoul2014/ip2region/blob/master/data/ip2region.db?raw=true")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
