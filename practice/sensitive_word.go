package practice

import "strings"

func SensitiveWReplace() {
	sensitiveWords := []string{
		"傻逼",
		"傻叉",
		"垃圾",
		"妈的",
		"sb",
	}
	text := "什么垃圾打野，傻逼一样，叫你来开龙不来，sb"

	for _, word := range sensitiveWords {
		replaceChar := ""
		for i, wordLen := 0, len(word); i < wordLen; i++ {
			// 根据敏感词的长度构造和谐字符
			replaceChar += "*"
		}
		text = strings.Replace(text, word, replaceChar, -1)
	}

	println("text -> ", text)
}

func SensitiveRuneReplace() {
	sensitiveWords := []string{
		"傻逼",
		"傻叉",
		"垃圾",
		"妈的",
		"sb",
	}
	text := "什么垃圾打野，傻逼一样，叫你来开龙不来，sb"

	for _, word := range sensitiveWords {
		replaceChar := ""
		for i, wordLen := 0, len([]rune(word)); i < wordLen; i++ {
			// 根据敏感词的长度构造和谐字符
			replaceChar += "*"
		}
		text = strings.Replace(text, word, replaceChar, -1)
	}

	println("text -> ", text)
}
