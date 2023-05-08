package util

import "regexp"

type RegRep struct {
	Reg, Rep string
}

func RegReplace(str string, m []RegRep) string {
	content := str
	for _, s2 := range m {
		compile := regexp.MustCompile(s2.Reg)
		content = compile.ReplaceAllString(content, s2.Rep)
	}
	return content
}
