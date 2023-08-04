package custom_http

import (
	"giao/pkg/util"
	"net/http"
)

func Fetch(url string, header map[string]string) *http.Response {
	c := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	//req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.183")
	for k, v := range header {
		req.Header.Set(k, v)
	}
	util.CheckErr(err)
	res, err := c.Do(req)
	util.CheckErr(err)
	return res
}
