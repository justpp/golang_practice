package server

import (
	"encoding/json"
	"fmt"
	"giao/pkg/tour/chat_room/src/logic"
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

func userListHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	userList := logic.Broadcaster.GetUserList()

	b, err := json.Marshal(userList)
	if err != nil {
		fmt.Fprint(w, `[]`)
	} else {
		fmt.Fprint(w, string(b))
	}
}
