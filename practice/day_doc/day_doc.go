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
	domain    string
	docUrl    string
	dir       string
	Urls      [][2]string
	cssJsUrls []string
}

func Gomianshiti() *DayDoc {
	return &DayDoc{
		"https://www.topgoer.cn",
		"https://www.topgoer.cn/docs/gomianshiti/mianshiti",
		"./download/day_doc",
		nil,
		nil,
	}
}

func GolangDesign() *DayDoc {
	return &DayDoc{
		"https://www.topgoer.cn",
		"https://www.topgoer.cn/docs/golang-design-pattern/golang-design-pattern-1cbgha2ltg796",
		"./download/golang_design",
		nil,
		nil,
	}
}

func (d *DayDoc) Run() {
	d.regUrls()
	d.runDownload()
}

func (d *DayDoc) regUrls() {
	resp, err := http.Get(d.docUrl)
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
	// css链接
	compile := regexp.MustCompile(`<link.+?\s*href="(/static/.+?)"[^>]*>`)
	matches := compile.FindAllSubmatch(body, -1)
	for _, match := range matches {
		d.cssJsUrls = append(d.cssJsUrls, string(match[1]))
	}
	// js链接
	compile = regexp.MustCompile(`<script.+?\s*src="(/static/.+?)"[^>]*>`)
	matches = compile.FindAllSubmatch(body, -1)
	for _, match := range matches {
		d.cssJsUrls = append(d.cssJsUrls, string(match[1]))
	}
	// 文件内遗漏的文件链接
	d.cssJsUrls = append(d.cssJsUrls, []string{
		"/static/bootstrap/css/bootstrap.min.css.map",
		"/static/editor.md/fonts/fontawesome-webfont.woff2?v=4.3.0",
		"/static/layer/skin/default/layer.css?v=3.0.3303",
		"/static/editor.md/fonts/fontawesome-webfont.woff?v=4.3.0",
		"/static/font-awesome/fonts/fontawesome-webfont.woff2?v=4.7.0",
		"/static/font-awesome/fonts/fontawesome-webfont.woff?v=4.7.0",
		"/static/font-awesome/fonts/fontawesome-webfont.ttf?v=4.7.0",
		"/static/jstree/3.3.4/themes/default/throbber.gif",
		"/static/jstree/3.3.4/themes/default/32px.png",
	}...)

	// 侧边栏链接
	compile = regexp.MustCompile(`<a.+?\s*href="(https://www.topgoer.cn/docs/.+?)"[^>]*title="(.+?)"[^>]*>`)
	matches = compile.FindAllSubmatch(body, -1)
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
	// css、js
	wg.Add(len(d.cssJsUrls))
	for _, s2 := range d.cssJsUrls {
		split := strings.Split(s2, "/")
		fileDir := fmt.Sprintf("%s%s", d.dir, strings.Join(split[:len(split)-1], "/"))
		if exists, _ := util.IsExists(fileDir); !exists {
			err := os.MkdirAll(fileDir, os.ModePerm)
			if err != nil {
				fmt.Println("create css、js dir err", err)
				return
			}
		}
		filename := strings.Split(split[len(split)-1], "?")[0]
		go func(name string, url string) {
			filename := fmt.Sprintf("%s/%s", fileDir, name)
			url = fmt.Sprintf("%s%s", d.domain, url)
			util.CreateFile(filename, []byte(d.getHtml(url)))
			wg.Done()
		}(filename, s2)
	}
	// 侧边栏链接
	wg.Add(len(d.Urls))
	for _, s2 := range d.Urls {
		go func(name string, url string) {
			util.CreateFile(fmt.Sprintf("%s/%s.html", d.dir, name), []byte(d.getHtml(url)))
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
	if strings.Index(url, "/static/") == -1 {
		for _, i2 := range d.Urls {
			// 替换本地路径
			str = strings.Replace(str, i2[1], fmt.Sprintf("./%s.html", i2[0]), 1)
			// 替换manual-title
			str = strings.Replace(str, `go语言面试题`, "learn learn learn learn ", -1)
		}
		return strings.Replace(str, "/static/", "./static/", -1)
	}
	if strings.Index(url, "/static/") > -1 {
		return strings.Replace(str, `function loadDocument($url, $id, $callback) {`, `
function loadDocument($url, $id, $callback) {
    window.open($url,'_self')
    events.trigger('article.open', {$url: $url, $id: $id});
    return false;
`, 1)
	}
	return str
}
