package sites

import (
	"fmt"
	"giao/pkg/e_book/src"
	"giao/pkg/e_book/src/sites/m2_ddyueshu_com"
	"giao/pkg/e_book/src/sites/sj_uukanshu_com"
	"giao/pkg/e_book/src/sites/www_vbiquge_co"
	"giao/pkg/util"
	"regexp"
)

type SiteCenter struct {
	m map[string]src.SiteT
}

func (sc *SiteCenter) Add(s src.SiteT) {
	if s != nil {
		sc.m[s.GetHost()] = s
	}
}

func (sc *SiteCenter) InitImport() *SiteCenter {
	sc.Add(&www_vbiquge_co.Site{})
	sc.Add(&sj_uukanshu_com.Site{})
	sc.Add(&m2_ddyueshu_com.Site{})
	return sc
}

// GetSite url  like http://www.vbiquge.co/58_58437/
func (sc *SiteCenter) GetSite(url string) src.SiteT {

	compile, err := regexp.Compile("://(.+?)/")
	util.CheckErr(err)
	sub := compile.FindStringSubmatch(url)

	var host string
	if len(sub) != 2 {
		panic("url 错误")
	}
	host = sub[1]
	site, ok := sc.m[host]
	if ok {
		return site
	}
	fmt.Println("site", sc.m, host)
	panic("未找到该站点配置")
}

func New() *SiteCenter {
	return &SiteCenter{m: make(map[string]src.SiteT)}
}
