package jd

import (
	"bufio"
	"fmt"
	"io"
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
	req.Header.Set("Connection", MapDefaultVal(HeaderMap, "Connection", j.Connection))
	req.Header.Set("Referer", MapDefaultVal(HeaderMap, "Referer", j.Url.Login))
	req.Header.Set("Accept", MapDefaultVal(HeaderMap, "Accept", j.Accept))
	for k, v := range HeaderMap {
		req.Header.Set(k, v)
	}
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
	const str = "./jd/cookies/cookie.txt"
	_, err := os.Stat(str)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("cookie str 不存在")
			return nil
		} else {
			return err
		}
	}
	file, err := os.Open("./jd/cookies/cookie.txt")
	if err != nil {
		return err
	}
	reader := bufio.NewReader(file)
	var cookiesArr = make(map[int][]*http.Cookie)
	var count int
	for {
		lineStr, err := reader.ReadString('\n')
		// net/http cookie验证不能有双引号
		lineStr = strings.Replace(lineStr, "\"", "", -1)
		if err != nil {
			break
		}
		var cookies []*http.Cookie
		for _, i := range strings.Split(strings.TrimSpace(lineStr), ";") {
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
		cookiesArr[count] = cookies
		count++
	}
	err = SaveCookie(cookiesArr)
	if err != nil {
		return err
	}
	return nil
}

func (j *JD) checkCookieIsLogin() bool {
	err := j.LoadCookie()
	if err != nil {
		fmt.Println("load cookie err", err)
		return false
	}
	isLogin, err := j.validateCookies()
	if err != nil {
		fmt.Println("cookie err ", err)
		return false
	}
	if !isLogin {
		fmt.Println("cookie 失效")
		return false
	}
	return true
}
