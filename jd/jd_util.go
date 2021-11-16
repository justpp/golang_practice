package jd

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
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

func CookieStr2Json() error {
	const str = "./cookies/cookie.txt"
	_, err := os.Stat(str)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("cookie str 不存在")
			return nil
		} else {
			return err
		}
	}
	cookiesByte, err := ioutil.ReadFile(str)
	if err != nil {
		return err
	}
	var cookies []*http.Cookie

	for _, i := range strings.Split(string(cookiesByte), ";") {
		s := strings.Split(i, "=")
		name := s[0]
		val := s[1]
		cookies = append(cookies, &http.Cookie{
			Name:   name,
			Value:  val,
			Domain: ".jd.com",
			Path:   "/",
		})
	}
	err = SaveCookie(cookies)
	if err != nil {
		return err
	}
	return nil
}
