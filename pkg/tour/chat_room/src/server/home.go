package server

import (
	"fmt"
	"net/http"
	"text/template"
)

func homeHandleFunc(w http.ResponseWriter, r *http.Request) {
	files, err := template.ParseFiles(rootDir + "/src/template/home.html")
	if err != nil {
		_, _ = fmt.Fprint(w, "模板解析错误")
		return
	}
	err = files.Execute(w, nil)
	if err != nil {
		_, _ = fmt.Fprint(w, "模板执行错误")
		return
	}
}
