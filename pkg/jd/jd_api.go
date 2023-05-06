package jd

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

func (j *JD) RunAPi() {
	j.UserInfo()
	j.JDBean()
}

func (j *JD) UserInfo() {
	//https://me-api.jd.com/user_new/info/GetJDUserInfoUnion
	u := j.createUrlWithArgs("https://me-api.jd.com/user_new/info/GetJDUserInfoUnion", map[string]string{})
	req, err := j.NewRequestWithHead(http.MethodGet, u, nil, nil)
	if err != nil {
		fmt.Println("JDBean error", err)
	}
	resp, err := j.Client.Do(req)
	if err != nil {
		fmt.Println("err", err)
	}
	defer resp.Body.Close()
	all, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("all", string(all))
}

func (j *JD) JDBean() {
	u := j.createUrlWithArgs("https://api.m.jd.com/client.action?functionId=signBeanIndex&appid=ld", map[string]string{})
	req, err := j.NewRequestWithHead(http.MethodGet, u, nil, nil)
	if err != nil {
		fmt.Println("JDBean error", err)
	}
	resp, err := j.Client.Do(req)
	if err != nil {
		fmt.Println("err", err)
	}
	defer resp.Body.Close()
	all, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(all))
	json := gjson.Parse(string(all))
	daily := json.Get("data.dailyAward")
	if !daily.Exists() {
		daily = json.Get("data.continuityAward")
	}
	fmt.Println("签到领京豆", daily.Get("title").Str, "获得", daily.Get("beanAward.beanCount"), "个京豆")
}