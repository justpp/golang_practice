package jd

import (
	"io"
	"net/http"
	"net/url"
)

func (j *JD) NewRequestWithHead(Method, URL string, HeaderMap map[string]string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(Method, URL, body)
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
