package day_doc

import (
	"fmt"
	"giao/util"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
)

type DayDoc struct {
	dir  string
	Urls [][2]string
}

func (d *DayDoc) Run() {
	d.dir = "./download/day_doc"
	d.regUrls()
	d.runDownload()
}

func (d *DayDoc) regUrls() {
	mainUrl := "https://www.topgoer.cn/docs/gomianshiti/mianshiti"
	resp, err := http.Get(mainUrl)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("get err", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadAll err", err)
		return
	}
	compile := regexp.MustCompile(`<a.+?\s*href="(https://www.topgoer.cn/docs/gomianshiti/.+?)"[^>]*title="(.+?)"[^>]*>`)
	matches := compile.FindAllSubmatch(body, -1)
	fmt.Println(len(matches))
	for _, match := range matches {
		d.Urls = append(d.Urls, [2]string{string(match[2]), string(match[1])})
	}
}

func (d *DayDoc) runDownload() {
	// 指定一个空文件夹
	if exists, _ := util.IsExists(d.dir); exists {
		err := os.RemoveAll(d.dir)
		if err != nil {
			return
		}
	}
	err := os.MkdirAll(d.dir, os.ModePerm)
	if err != nil {
		fmt.Println("create dir err", err)
		return
	}
	wg := sync.WaitGroup{}
	wg.Add(len(d.Urls))
	for _, s2 := range d.Urls {
		go func(title string, url string) {
			util.CreateFile(fmt.Sprintf("%s/%s.html", d.dir, title), []byte(d.getHtml(url)))
			wg.Done()
		}(s2[0], s2[1])
	}
	wg.Wait()
}

func (d *DayDoc) getHtml(url string) string {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("get file err", err)
		return ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get file err", err)
		return ""
	}
	str := string(body)
	for _, i2 := range d.Urls {
		str = strings.Replace(str, i2[1], fmt.Sprintf("./%s.html", i2[0]), -1)
	}
	return strings.Replace(str, "/static/", "https://www.topgoer.cn/static/", -1)
}
