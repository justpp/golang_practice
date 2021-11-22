package jd

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

func (j *JD) JDBean() error {
	u := j.createUrlWithArgs("https://api.m.jd.com/client.action?functionId=signBeanIndex&appid=ld", map[string]string{})
	req, err := j.NewRequestWithHead(http.MethodGet, u, nil, nil)
	if err != nil {
		return err
	}

	resp, err := j.Client.Do(req)
	if err != nil {
		fmt.Println("err", err)
		return err
	}
	defer resp.Body.Close()
	all, _ := ioutil.ReadAll(resp.Body)
	json := gjson.Parse(string(all))
	daily := json.Get("data.dailyAward")
	fmt.Println("签到领京豆", daily.Get("title").Str, "获得", daily.Get("beanAward.beanCount"), "个京豆")
	return nil
}
