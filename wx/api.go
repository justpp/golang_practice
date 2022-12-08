package wx

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"net/url"
)

type accessTokenRespT struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
	// 错误参数
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type UserInfoResT struct {
	Openid     string      `json:"openid"`
	Nickname   interface{} `json:"nickname"`
	Sex        int         `json:"sex"`
	Province   string      `json:"province"`
	City       string      `json:"city"`
	Country    string      `json:"country"`
	Headimgurl string      `json:"headimgurl"`
	Privilege  []string    `json:"privilege"`
	Unionid    string      `json:"unionid"`
	Errcode    int         `json:"errcode"`
	Errmsg     string      `json:"errmsg"`
}

var userInfo UserInfoResT

// CheckSignature 接口服务器验证
func CheckSignature(c *gin.Context) {
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	signature := c.Query("signature")
	echostr := c.Query("echostr")
	signatureGen := makeSignature(timestamp, nonce)
	if signatureGen != signature {
		log.Println("微信公众号接入失败")
		return
	}
	log.Println("微信公众号接入成功")
	_, _ = c.Writer.WriteString(echostr)
}

func fetchCode(c *gin.Context) {
	RedirectUri := url.QueryEscape("http://152.69.198.203/get_userinfo")
	oauthUrl := "https://open.weixin.qq.com/connect/oauth2/authorize" +
		"?appid=" + appId +
		"&redirect_uri=" + RedirectUri +
		"&response_type=code" +
		"&scope=snsapi_base" +
		"&state=giao" +
		"#wechat_redirect"
	c.Redirect(http.StatusFound, oauthUrl)
}

func fetchAccessToken(code string) *accessTokenRespT {
	apiUrl := "https://api.weixin.qq.com/sns/oauth2/access_token" +
		"?appid=" + appId +
		"&secret=" + appSecret +
		"&code=" + code +
		"&grant_type=authorization_code"
	resp, err := http.Get(apiUrl)
	defer resp.Body.Close()
	if err != nil {
		log.Println("fetchAccessToken err", err)
		return nil
	}
	log.Println("code", code)

	content, _ := io.ReadAll(resp.Body)
	var respT accessTokenRespT
	json.Unmarshal(content, &respT)
	if respT.Errcode != 0 {
		log.Println("fetch err", respT.Errcode, respT.Errmsg)
		return nil
	}
	return &respT
}

func fetUserInfo(accessToken, openid string) UserInfoResT {
	apiUrl := "https://api.weixin.qq.com/sns/userinfo" +
		"?access_token=" + accessToken +
		"&openid=" + openid +
		"&lang=zh_CN"
	resp, err := http.Get(apiUrl)
	if err != nil {
		return UserInfoResT{}
	}
	defer resp.Body.Close()
	var u UserInfoResT

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return u
	}
	json.Unmarshal(all, u)
	if u.Errcode != 0 {
		log.Println("fetUserInfo err", u.Errcode, u.Errmsg)
	}
	return u
}

func getUserInfo(c *gin.Context) {
	code := c.Query("code")
	//state := c.Query("state")
	accessTokenT := fetchAccessToken(code)
	log.Println("accessTokenT", accessTokenT)
	userInfo = fetUserInfo(accessTokenT.AccessToken, accessTokenT.Openid)
	if userInfo.Errcode != 0 {
		c.Writer.Write([]byte(userInfo.Errmsg))
		return
	}
	log.Println(userInfo)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"userInfo": userInfo,
		"title":    "啦啦啦说了",
	})
}
