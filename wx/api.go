package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"giao/wx/gpt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
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

var userInfo *UserInfoResT

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

type WxReqXml struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	MsgId        int64  `xml:"MsgId"`
}

func ReceiveMsg(c *gin.Context) {
	var wxMsg WxReqXml
	err := c.ShouldBindXML(&wxMsg)
	if err != nil {
		return
	}
	fmt.Println("raw data", wxMsg)
	answer := gpt.Talk(talkSecret, wxMsg.FromUserName, wxMsg.Content)
	WXMsgReply(c, wxMsg.ToUserName, wxMsg.FromUserName, answer)
}

func fetchCode(c *gin.Context) {
	RedirectUri := url.QueryEscape("http://168.138.33.211:8892/get_userinfo")
	oauthUrl := "https://open.weixin.qq.com/connect/oauth2/authorize" +
		"?appid=" + appId +
		"&redirect_uri=" + RedirectUri +
		"&response_type=code" +
		"&scope=snsapi_userinfo" +
		"&state=giao" +
		"#wechat_redirect"
	log.Println("oauthUrl", oauthUrl)
	c.Redirect(http.StatusFound, oauthUrl)
}

func fetchUserAccessToken(code string) *accessTokenRespT {
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

func fetchAccessToken() *accessTokenRespT {
	apiUrl := "https://api.weixin.qq.com/cgi-bin/token" +
		"?appid=" + appId +
		"&secret=" + appSecret +
		"&grant_type=client_credential"
	resp, err := http.Get(apiUrl)
	defer resp.Body.Close()
	if err != nil {
		log.Println("fetchAccessToken err", err)
		return nil
	}

	content, _ := io.ReadAll(resp.Body)
	var respT accessTokenRespT
	json.Unmarshal(content, &respT)
	if respT.Errcode != 0 {
		log.Println("fetch err", respT.Errcode, respT.Errmsg)
		return nil
	}
	return &respT
}

func fetUserInfo(accessToken, openid string) *UserInfoResT {
	apiUrl := "https://api.weixin.qq.com/sns/userinfo" +
		"?access_token=" + accessToken +
		"&openid=" + openid +
		"&lang=zh_CN"
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	var u UserInfoResT

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	json.Unmarshal(all, &u)
	log.Println("fetUserInfo content", string(all))
	log.Println("u", u)
	if u.Errcode != 0 {
		log.Println("fetUserInfo err", u.Errcode, u.Errmsg)
	}
	return &u
}

func getUserInfo(c *gin.Context) {
	code := c.Query("code")
	// state := c.Query("state")
	accessTokenT := fetchUserAccessToken(code)
	log.Println("accessTokenT", accessTokenT)
	userInfo = fetUserInfo(accessTokenT.AccessToken, accessTokenT.Openid)
	if userInfo.Errcode != 0 {
		c.Writer.Write([]byte(userInfo.Errmsg))
		return
	}
	log.Println("userInfo", userInfo)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"userInfo": userInfo,
		"title":    "啦啦啦说了.",
	})
}

type Button struct {
	Type     string   `json:"type"`
	Name     string   `json:"name"`
	Key      string   `json:"key,omitempty"`
	Url      string   `json:"url,omitempty"`
	AppID    string   `json:"appid,omitempty"`
	PagePath string   `json:"pagepath,omitempty"`
	SubBtn   []Button `json:"sub_button,omitempty"`
}

type Menu struct {
	Button []Button `json:"button"`
}

func setMenu(c *gin.Context) {
	accessToken := fetchAccessToken()
	menuBtnMap := Menu{
		Button: []Button{
			{Type: "click", Name: "欸嘿", Key: "V1001_TODAY_MUSIC"},
			{Name: "菜单", SubBtn: []Button{
				{Type: "view", Name: "搜索", Url: "http://www.soso.com/"},
				{Type: "scancode_waitmsg", Name: "扫码带提示", Key: "rselfmenu_0_0", SubBtn: []Button{}},
				{Type: "scancode_push", Name: "扫码推事件", Key: "rselfmenu_0_1", SubBtn: []Button{}}},
			},
		},
	}
	menu, _ := json.Marshal(menuBtnMap)

	apiUrl := "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=" + accessToken.AccessToken
	resp, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(menu))
	defer resp.Body.Close()
	if err != nil {
		log.Println("create menu err", err)
		return
	}

	content, _ := io.ReadAll(resp.Body)
	fmt.Println("content", string(content))
	return
}

func delMenu(c *gin.Context) {
	accessToken := fetchAccessToken()

	apiUrl := "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=" + accessToken.AccessToken
	resp, err := http.Get(apiUrl)
	defer resp.Body.Close()
	if err != nil {
		log.Println("del menu fail err", err)
		return
	}

	content, _ := io.ReadAll(resp.Body)
	fmt.Println("content", string(content))
	return
}

// WXRepTextMsg 微信回复文本消息结构体
type WXRepTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	// 若不标记XMLName, 则解析后的xml名为该结构体的名称
	XMLName xml.Name `xml:"xml"`
}

// WXMsgReply 微信消息回复
func WXMsgReply(c *gin.Context, fromUser, toUser string, resText string) {
	repTextMsg := WXRepTextMsg{
		ToUserName:   toUser,
		FromUserName: fromUser,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      resText,
	}

	msg, err := xml.Marshal(&repTextMsg)
	if err != nil {
		log.Printf("[消息回复] - 将对象进行XML编码出错: %v\n", err)
		return
	}
	_, _ = c.Writer.Write(msg)
}
