package jd

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (j *JD) GetItemDetailPage() error {
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
	fmt.Println("item info", string(all))
	return nil
}
