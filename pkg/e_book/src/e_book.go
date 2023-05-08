package src

import (
	"fmt"
	"giao/pkg/util"
	"giao/pkg/util/custom_http"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/simplifiedchinese"
	"log"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

type chapter struct {
	url        string
	nextUrl    string
	title      string
	content    string
	retryCount int
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

	// e.host = "https://m2.ddyueshu.com"
	// e.linkCount = 1
	// e.G = 1
	// e.menuMap = make(map[int]*chapter)
	// e.menuMap[1] = &chapter{
	// 	// url: "/wapbook/11082821_703430439.html",
	// 	// url: "/wapbook/11082821_723630156.html",
	// 	// url:   "/wapbook/11082821_757213018.html",
	// 	url:   "/wapbook/11082821_747122134.html",
	// 	title: "第1208章 元央界和二代洛白衣(2/3,求月票）",
	// }

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
		fmt.Println(e.nextUrl)
		e.fetchMenu()
	}
	fmt.Println("章节：", e.linkCount)
}

func (e *EBook) fetchMenu() {
	res := custom_http.Fetch(e.nextUrl, nil)
	defer res.Body.Close()

	utf8, err := charset.NewReader(res.Body, "UTF-8")
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
		err := os.Mkdir(dirName, 0744)
		util.CheckErr(err)
	}
	content := ""

	for i := 0; i < e.linkCount+1; i++ {
		c, ok := e.menuMap[i]
		if ok {
			content = content + "\n\n" + c.title + "\n\n" + c.content
		}
	}

	if e.name == "" {
		e.name = strconv.FormatInt(time.Now().Unix(), 10)
	}
	fileName := dirName + e.name + ".txt"
	fmt.Println("filename", fileName)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	util.CheckErr(err)
	defer file.Close()
	_, err = file.Write([]byte(content))
	util.CheckErr(err)
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
	defer func() {
		a := recover()
		if a != nil {
			fmt.Println("defer err", a, c.url)
		}
	}()
	res := custom_http.Fetch(c.nextUrl, nil)
	currUrl := c.nextUrl
	body := res.Body
	defer body.Close()

	simpl := simplifiedchinese.GB18030.NewDecoder().Reader(body)

	reader, err := goquery.NewDocumentFromReader(simpl)
	util.CheckErr(err)

	reader.Find("*").Each(func(i int, selection *goquery.Selection) {
		// 去除js
		if selection.Is("script") {
			selection.Remove()
		}
	})

	readerHtml, err := reader.Html()
	if err != nil {
		return
	}
	util.DD(readerHtml)

	content, _ := reader.Find("#chaptercontent").Html()
	nextHref, _ := reader.Find("#pb_next").Attr("href")

	c.nextUrl = ""
	str1 := strings.Replace(nextHref, ".html", "", 1)
	str2 := strings.Replace(c.url, ".html", "", 1)

	if strings.HasPrefix(str1, str2) {
		c.nextUrl = e.host + nextHref
	}

	compile := regexp.MustCompile(`<br/>.*?<br/>`)
	matchString := compile.FindAllStringSubmatch(content, -1)

	content = ""
	for _, v := range matchString {
		str := strings.ReplaceAll(v[0], "<br/>", "")
		strReg := []util.RegRep{
			{`第*\(.*?\)页*`, " "},
			{`记住+?.*?\.com`, " "},
			{` `, " "},
		}
		str = util.RegReplace(str, strReg)
		str = strings.TrimSpace(str)
		if len(str) == 0 {
			continue
		}

		if content == "" {
			content = str
		} else {
			content += "\n\n	" + str
		}
	}

	if len(content) == 0 {
		// c.retryCount++
		log.Println("弹广告了", currUrl)
		// c.nextUrl = currUrl
		if c.retryCount > 5 {
			log.Fatalln("tmd试了", c.retryCount, "次了:", currUrl)
		}
		// time.Sleep(time.Second * 6)
		return
	}
	if content[0:1] != " " {
		content = "    " + content
	}
	c.content = c.content + content
}
