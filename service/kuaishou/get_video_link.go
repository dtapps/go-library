package kuaishou

import (
	"errors"
	"log"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

type Response struct {
	Link    string
	Cookies string
	Err     error
}

func (c *Client) GetVideoLink(url string) Response {

	result := Response{}

	jar, err := cookiejar.New(&cookiejar.Options{})
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Jar: jar,
	}
	request, err := http.NewRequest("GET", url, nil)

	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	// request.Header.Add("Host", "c.kuaishou.com")
	// request.Header.Add("Host", "v.kuaishouapp.com")

	result.Err = err
	var cookieList []string

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {

		if len(via) > 10 {
			return errors.New("too many redirects")
		}

		result.Link = req.URL.String()

		return nil
	}

	response, _ := client.Do(request)

	cookies := jar.Cookies(request.URL)
	for _, cookie := range cookies {
		cookieList = append(cookieList, cookie.String())
	}

	result.Cookies = strings.Join(cookieList, ";")

	defer response.Body.Close()

	return result
}
