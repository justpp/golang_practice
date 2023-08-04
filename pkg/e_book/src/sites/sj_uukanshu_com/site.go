package sj_uukanshu_com

import (
	"giao/pkg/e_book/src"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
)

//https://sj.uukanshu.com/read.aspx?tid=135864

type Site struct {
}

func (s *Site) GetHost() string {
	return "sj.uukanshu.com"
}

func (s *Site) GetBookName(reader *goquery.Document) string {
	return reader.Find(".bookname").Text()
}

func (s *Site) GetMenuNextPageUrl(reader *goquery.Document) string {
	val, exists := reader.Find(".CurrentPage + a").Attr("href")
	if !exists {
		return ""
	}
	return "/" + val
}

func (s *Site) GetMenuMap(e *src.EBook, reader *goquery.Document) {
	chapterList := reader.Find("#chapterList")
	chapterList.Find("a").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		title := selection.Text()

		e.LinkCount += 1
		count := e.LinkCount
		e.AddMenuMap(count, &src.Chapter{
			Url:   "/" + href,
			Title: title,
		})
	})
}

func (s *Site) GetChapterContent(c *src.Chapter, reader *goquery.Document, e *src.EBook) {

	content, _ := reader.Find("#bookContent").RemoveClass("box").Html()
	compile := regexp.MustCompile(`<p>(.*?)\s*?</p>`)
	content = compile.ReplaceAllString(content, "$1 \n	")

	compile = regexp.MustCompile(`<div class="box"[\s\S]*?</div>`)
	content = compile.ReplaceAllString(content, "")

	compile = regexp.MustCompile(`<br(/)*>`)
	content = compile.ReplaceAllString(content, "\n")
	c.NextUrl = ""

	if c.Content == "" {
		c.Content = strings.TrimSuffix(content, "\n")
	} else {
		c.Content = c.Content + strings.TrimSpace(content)
	}
}
