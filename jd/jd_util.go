package jd

import (
	"fmt"
	"giao/util"
	"github.com/tidwall/gjson"
	"io"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type JD struct {
	QrCodeURl    string // "https://qr.m.jd.com/show" ?appid=133&size=300&t=
	CheckSanUrl  string // "https://qr.m.jd.com/check"
	CheckTickUrl string // "https://passport.jd.com/uc/qrCodeTicketValidation?"
	Referer      string // "https://passport.jd.com/new/login.aspx"
	UserAgent    string // "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36"
	Connection   string // "keep-alive"
	Accept       string // "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3"
	JdCookie     []*http.Cookie
	Client       *http.Client
	Ticket       string
}

// JDLogin
func JDLogin() {
	j := JDInit()
	err := j.GetQrCode()
	if err != nil {
		fmt.Println("err", err)
		return
	}
	err = j.CheckScan()
	if err != nil {
		return
	}
}

func JDInit() *JD {
	j := &JD{
		QrCodeURl:   "https://qr.m.jd.com/show",
		CheckSanUrl: "https://qr.m.jd.com/check",
		Referer:     "https://passport.jd.com/new/login.aspx",
		UserAgent:   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36",
		Connection:  "keep-alive",
		Accept:      "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3",
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil
	}
	j.Client = &http.Client{Jar: jar}
	return j
}

func (j *JD) GetQrCode() error {
	args := url.Values{}
	args.Add("appid", "133")
	args.Add("size", "300")
	args.Add("t", strconv.FormatInt(time.Now().Unix()*1e3, 10))
	u := j.QrCodeURl + "?" + args.Encode()
	req, err := j.NewRequest(http.MethodGet, u)
	if err != nil {
		fmt.Println("new url err", err)
		return err
	}
	resp, _ := j.Client.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		util.DownLoadImg(resp.Body, "./qr_code.png")
	}
	j.JdCookie = resp.Cookies()
	return nil
}

func (j *JD) CheckScan() error {
	for i := 0; i < 85; i++ {
		args := url.Values{}
		args.Add("appid", "133")
		args.Add("callback", fmt.Sprintf("jQuery%v", rand.Intn(9999999-1000000)+1000000))

		var token string
		for _, v := range j.JdCookie {
			if v.Name == "wlfstk_smdl" {
				token = v.Value
				break
			}
		}
		if token == "" {
			fmt.Println("获取token失败")
			return nil
		}
		args.Add("token", token)
		args.Add("_", strconv.FormatInt(1e3*time.Now().Unix(), 10))
		u := j.CheckSanUrl + "?" + args.Encode()
		req, err := j.NewRequest("GET", u)
		if err != nil {
			return err
		}
		resp, _ := j.Client.Do(req)
		if err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
		all, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		fmt.Println("check res", string(all))
		respMsg := string(all)

		n1 := strings.Index(respMsg, "(")
		n2 := strings.Index(respMsg, ")")

		json := gjson.Parse(string(all[n1+1 : n2]))
		fmt.Println("json", json)

		j.Ticket = json.Get("ticket").Str
		resp.Body.Close()
	}
	fmt.Println("超时了")
	return nil
}

func (j *JD) NewRequest(Method, URL string) (*http.Request, error) {
	req, err := http.NewRequest(Method, URL, nil)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return req, err
	}
	req.Header.Set("User-Agent", j.UserAgent)
	req.Header.Set("Connection", j.Connection)
	req.Header.Set("Referer", j.Referer)
	req.Header.Set("Accept", j.Accept)
	return req, nil
}
