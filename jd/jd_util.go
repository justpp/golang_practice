package jd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (j *JD) NewRequestWithHead(Method, URL string, HeaderMap map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(Method, URL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", MapDefaultVal(HeaderMap, "User-Agent", j.UserAgent))
	req.Header.Set("Connection", MapDefaultVal(HeaderMap, "User-Agent", j.Connection))
	req.Header.Set("Referer", MapDefaultVal(HeaderMap, "User-Agent", j.Url.Login))
	req.Header.Set("Accept", MapDefaultVal(HeaderMap, "User-Agent", j.Accept))
	return req, nil
}

func (j JD) createUrlWithArgs(u string, argsMap map[string]string) string {
	args := url.Values{}
	for i, v := range argsMap {
		args.Add(i, v)
	}
	return u + args.Encode()
}

func MapDefaultVal(m map[string]string, k, defaultVal string) string {
	v, ok := m[k]
	if !ok {
		return defaultVal
	}
	return v
}

func (j *JD) GetItemDetailPage() error {
	u := j.createUrlWithArgs("https://cart.jd.com/gate.action?",
		map[string]string{
			"pid":    "1179553",
			"pcount": "1",
			"ptype":  "1",
		})
	req, err := j.NewRequestWithHead(http.MethodGet, u, map[string]string{"Referer": j.Url.CenterList})
	if err != nil {
		return err
	}
	resp, _ := j.Client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	all, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("item info", all)
	fmt.Println("item state", resp.Status)
	return nil
}
