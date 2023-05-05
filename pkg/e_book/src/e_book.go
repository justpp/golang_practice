package src

import (
	"fmt"
	"giao/pkg/util"
	"giao/pkg/util/custom_http"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync/atomic"
	"time"
)

type chapter struct {
	url     string
	nextUrl string
	title   string
	content string
}

type EBook struct {
	host      string
	nextUrl   string
	name      string
	G         int // 限制协程数
	linkCount int
	menuMap   map[int]*chapter
}

func (e *EBook) Run(EbookUrl string) {
	start := time.Now()
	e.FetchMenuList(EbookUrl)

	e.goFetchData()
	fmt.Println("内容已获取完成")

	e.download()

	fmt.Println("花费：", time.Now().Sub(start).Seconds())
}

func (e *EBook) FetchMenuList(EbookUrl string) {
	u, _ := url.Parse(EbookUrl)
	e.host = u.Scheme + "://" + u.Hostname()
	e.nextUrl = u.String()

	e.linkCount = 0
	e.menuMap = make(map[int]*chapter)

	for e.nextUrl != "" {
		e.fetchMenu()
	}
	fmt.Println("章节：", e.linkCount)
}

func (e *EBook) fetchMenu() {
	body := custom_http.Fetch(e.nextUrl, nil)
	defer body.Close()

	utf8, err := charset.NewReader(body, "UTF-8")
	util.CheckErr(err)
	e.nextUrl = ""

	reader, err := goquery.NewDocumentFromReader(utf8)
	util.CheckErr(err)

	if e.name == "" {
		bookName, _ := reader.Find("meta[property=\"og:novel:book_name\"]").Attr("content")
		e.name = bookName
	}

	nextPage, _ := reader.Find(".listpage .right>a.onclick").Attr("href")
	if nextPage != "" {
		e.nextUrl = e.host + nextPage
	}

	chapterList := reader.Find(".book_last").Last()
	chapterList.Find("a").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		title, _ := selection.Html()
		e.linkCount += 1
		count := e.linkCount
		e.menuMap[count] = &chapter{
			url:   href,
			title: title,
		}
	})
}

func (e *EBook) download() {
	dirName := "./download/"
	if ok, _ := util.IsExists(dirName); !ok {
		_ = os.Mkdir(dirName, 0644)
	}
	content := ""

	for i := 0; i < e.linkCount+1; i++ {
		c, ok := e.menuMap[i]
		if ok {
			// content = content + "\n\n" + c.title + "\n\n" + c.content
			content = content + "\n\n" + "\n\n" + c.content
		}
	}

	content = strings.Replace(content, "<br/><br/>", "\n", -1)

	fileName := dirName + e.name + ".txt"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer file.Close()
	util.CheckErr(err)
	_, _ = file.Write([]byte(content))
}

func (e *EBook) goFetchData() {
	g := e.G
	if g == 0 {
		g = 20
	}
	fmt.Println("开启协程:", g)

	var opts int32 = 0

	gLimit := util.NewGLimit(g)
	gLimit.Run(e.linkCount, func(i int) {
		atomic.AddInt32(&opts, 1)
		c, ok := e.menuMap[i+1]
		if ok {
			e.fetchContent(c)
		}

		fmt.Println("获取进度:", (float32(opts)/float32(e.linkCount))*100)
	})
}

func (e *EBook) fetchContent(c *chapter) {
	c.nextUrl = e.host + c.url
	for c.nextUrl != "" {
		c.fetchPage(e)
	}
}

func (c *chapter) fetchPage(e *EBook) {
	body := custom_http.Fetch(c.nextUrl, nil)
	defer body.Close()

	utf8, err := charset.NewReader(body, "UTF-8")
	util.CheckErr(err)

	reader, err := goquery.NewDocumentFromReader(utf8)
	util.CheckErr(err)

	reader.Find("*").Each(func(i int, selection *goquery.Selection) {
		// 去除js
		if selection.Is("script") {
			selection.Remove()
		}
	})

	// html, err := reader.Find("#chaptercontent").Html()
	// if err != nil {
	// 	return
	// }
	// util.DD(html)

	content, _ := reader.Find("#chaptercontent").Html()
	nextHref, _ := reader.Find("#pb_next").Attr("href")

	c.nextUrl = ""
	str1 := strings.Replace(nextHref, ".html", "", 1)
	str2 := strings.Replace(c.url, ".html", "", 1)

	if strings.HasPrefix(str1, str2) {
		c.nextUrl = e.host + nextHref
	}
	compile, _ := regexp.Compile(`第\(.*?\)页`)
	util.CheckErr(err)
	all := compile.ReplaceAllString(content, "\n")

	compile, _ = regexp.Compile(`：\w.+?.com`)
	util.CheckErr(err)
	all = compile.ReplaceAllString(all, "\n")

	content = strings.ReplaceAll(all, "<br/>", "\n")
	if c.content == "" {
		c.content = strings.TrimSuffix(content, "\n")
	} else {
		c.content = c.content + strings.TrimSpace(content)
	}
}
