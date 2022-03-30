package jd

import (
	"errors"
	"fmt"
	"giao/practice"
	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type JD struct {
	Url         Url
	UserAgent   string // "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36"
	Connection  string // "keep-alive"
	Accept      string // "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3"
	JdCookie    []*http.Cookie
	JdCookieMap map[int][]*http.Cookie
	Client      *http.Client
	Ticket      string
	IsLogin     bool
}

type Url struct {
	Login        string // "https://passport.jd.com/new/login.aspx"
	QrCode       string // "https://qr.m.jd.com/show" ?appid=133&size=300&t=
	CheckSan     string // "https://qr.m.jd.com/check"
	CheckTick    string // "https://passport.jd.com/uc/qrCodeTicketValidation?"
	ValidateTick string // https://passport.jd.com/uc/qrCodeTicketValidation?
	GetUserInfo  string // "https://passport.jd.com/user/petName/getUserInfoForMiniJd.action?"
	CenterList   string // "https://order.jd.com/center/list.action"
	Jd           string // "https://www.jd.com"
}

func Login() {
	j := JdInit()

	err := j.LoadCookie()
	if err != nil {
		fmt.Println("load cookie err", err)
		return
	}
	for _, cookies := range j.JdCookieMap {
		j.JdCookie = cookies
		u, _ := url.Parse(j.Url.Jd)
		j.Client.Jar.SetCookies(u, cookies)
		isLogin, err := j.validateCookies()
		if err != nil {
			fmt.Println("err ", err)
			return
		}
		if !isLogin {
			fmt.Println("cookie 失效")
			j.QrCodeLogin()
		}
		err = j.JDBean()
		if err != nil {
			return
		}
	}
	//fmt.Println("签到结束 按回车键退出")
	//b := make([]byte, 1)
	//_, err = os.Stdin.Read(b)
	//if err != nil {
	//	return
	//}
	fmt.Println("签到结束")
}

func JdInit() *JD {
	jdUrl := Url{
		"https://passport.jd.com/new/login.aspx",
		"https://qr.m.jd.com/show",
		"https://qr.m.jd.com/check",
		"https://passport.jd.com/uc/qrCodeTicketValidation?",
		"https://passport.jd.com/uc/qrCodeTicketValidation?",
		"https://passport.jd.com/user/petName/getUserInfoForMiniJd.action?",
		"https://order.jd.com/center/list.action",
		"https://www.jd.com",
	}
	j := &JD{
		Url:        jdUrl,
		UserAgent:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36",
		Connection: "keep-alive",
		Accept:     "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3",
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil
	}
	j.Client = &http.Client{Jar: jar}
	return j
}

func (j *JD) QrCodeLogin() {
	err := j.GetQrCode()
	if err != nil {
		fmt.Println("err", err)
		return
	}
	err = j.CheckScan()
	if err != nil {
		return
	}
	ok, err := j.ValidateQrCodeTick(j.Ticket)
	if err != nil {
		fmt.Println("validate err", err)
		return
	}
	if !ok {
		fmt.Println("扫描未通过")
	}
	info, err := j.GetUserInfo(SaveCookie)
	if err != nil {
		return
	}
	fmt.Println("info", info)
}

func (j *JD) GetQrCode() error {
	args := url.Values{}
	args.Add("appid", "133")
	args.Add("size", "300")
	args.Add("t", strconv.FormatInt(time.Now().Unix()*1e3, 10))
	u := j.Url.QrCode + "?" + args.Encode()
	req, err := j.NewRequestWithHead(http.MethodGet, u, nil, nil)
	if err != nil {
		fmt.Println("new url err", err)
		return err
	}
	resp, _ := j.Client.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		practice.DownLoadImg(resp.Body, "./qr_code.png")
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
		u := j.Url.CheckSan + "?" + args.Encode()
		req, err := j.NewRequestWithHead("GET", u, map[string]string{}, nil)
		if err != nil {
			return err
		}
		resp, _ := j.Client.Do(req)
		defer resp.Body.Close()

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
		if j.Ticket != "" {
			break
		}
	}
	return nil
}

func (j *JD) ValidateQrCodeTick(tick string) (bool, error) {
	u := j.createUrlWithArgs(j.Url.ValidateTick, map[string]string{"t": tick})
	req, err := j.NewRequestWithHead(http.MethodGet, u, map[string]string{"Referer": "https://passport.jd.com/uc/login?ltype=logout"}, nil)
	if err != nil {
		return false, err
	}
	resp, _ := j.Client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return false, err
	}

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	fmt.Println("check res", string(all))
	return true, nil
}

func (j *JD) GetUserInfo(SaveCookie func(cookies map[int][]*http.Cookie) error) (string, error) {
	u := j.createUrlWithArgs(j.Url.GetUserInfo, map[string]string{
		"appid":    "133",
		"callback": fmt.Sprintf("jQuery%v", rand.Intn(9999999-1000000)+1000000),
		"_":        fmt.Sprintf("%v", time.Now().Unix()*1e3),
	})
	req, err := j.NewRequestWithHead(http.MethodGet, u, map[string]string{"Referer": j.Url.CenterList}, nil)
	if err != nil {
		return "", err
	}
	resp, err := j.Client.Do(req)
	if err != nil {
		return "", err
	}
	all, _ := ioutil.ReadAll(resp.Body)
	ret := gjson.Parse(string(all[14 : len(all)-1]))
	fmt.Println("get userinfo ", ret)
	m := make(map[int][]*http.Cookie)
	m[0] = req.Cookies()
	err = SaveCookie(m)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return "", nil
}

func SaveCookie(cookies map[int][]*http.Cookie) error {
	_, err := os.Stat("./jd/cookies")
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir("./jd/cookies", os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	cookiesFile := path.Join("./jd/cookies", fmt.Sprintf("%s.json", "cookie"))
	f, err := os.Create(cookiesFile)
	if err != nil {
		return err
	}
	defer f.Close()
	cookiesByte, err := jsoniter.Marshal(cookies)
	if err != nil {
		return err
	}
	_, err = f.Write(cookiesByte)
	if err != nil {
		return err
	}
	return nil
}

func (j *JD) LoadCookie() error {
	var cookies map[int][]*http.Cookie
	wd, _ := os.Getwd()
	cookieDir := fmt.Sprintf("%v/jd/cookies", wd)
	_, err := os.Stat(cookieDir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(cookieDir, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	// 从cookie.txt中读取cookie处理后存入cookie.json
	err = CookieStr2Json()
	if err != nil {
		return err
	}
	cookiesFile := path.Join(cookieDir, fmt.Sprintf("%s.json", "cookie"))
	cookiesByte, err := ioutil.ReadFile(cookiesFile)
	if err != nil {
		return err
	}
	err = jsoniter.Unmarshal(cookiesByte, &cookies)
	if err != nil {
		return err
	}
	j.JdCookieMap = cookies
	return nil
}

func (j *JD) validateCookies() (bool, error) {
	var infoUrl string
	var nickNamePath string
	if j.JdCookie[0].Name == "wlfstk_smdl" {
		infoUrl = j.Url.GetUserInfo
		nickNamePath = "nickName"
	} else {
		infoUrl = "https://me-api.jd.com/user_new/info/GetJDUserInfoUnion?"
		nickNamePath = "data.userInfo.baseInfo.nickname"
	}
	u := j.createUrlWithArgs(infoUrl, nil)
	req, err := j.NewRequestWithHead(http.MethodGet, u, map[string]string{"Referer": j.Url.Login}, nil)
	if err != nil {
		return false, err
	}
	j.Client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	resp, err := j.Client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return false, err
	}
	all, _ := ioutil.ReadAll(resp.Body)
	json := gjson.Parse(string(all))
	if json.Get("msg").Str == "not login" {
		return false, errors.New("cookie 失效，未登录")
	}
	nickName := json.Get(nickNamePath).Str
	if nickName != "" {
		fmt.Println(nickName, "已登录")
	}
	if resp.StatusCode != http.StatusOK {
		return false, nil
	}
	return true, nil
}
