package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"html/template"
	"net/http"
	"os"
)

func web(w http.ResponseWriter, r *http.Request) {
	path, err := os.Getwd()
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/%s", path, "lijie.io/websocket_test/demo1/index.html"))

	if err != nil {
		w.Write([]byte("模版加载错误"))
	}

	tmpl.Execute(w, nil)
}

func Echo(conn *websocket.Conn) {
	var err error
	for {
		var reply string

		err = websocket.Message.Receive(conn, &reply)
		if err != nil {
			fmt.Println("失败")
			break
		}

		err = websocket.Message.Send(conn, fmt.Sprintf("H5：%s；GO:收到", reply))
		if err != nil {
			fmt.Println("失败")
			break
		}
	}
}

func main() {
	http.HandleFunc("/web", web)

	http.Handle("/websocket", websocket.Handler(Echo))

	http.ListenAndServe("0.0.0.0:3003", nil)
}
