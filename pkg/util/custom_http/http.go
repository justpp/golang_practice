package custom_http

import (
	"giao/pkg/util"
	"io"
	"net/http"
)

func Fetch(url string, header map[string]string) (body io.ReadCloser) {
	c := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	// req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	for k, v := range header {
		req.Header.Set(k, v)
	}
	util.CheckErr(err)
	res, err := c.Do(req)
	util.CheckErr(err)

	return res.Body
}
