package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.HandlerFunc(makePage))
	http.ListenAndServe(":8782", nil)
}

func makePage(resw http.ResponseWriter, req *http.Request) {
	content, err := ioutil.ReadFile("indexhtml.txt")
	if err != nil {
		fmt.Println("テンプレートファイル読み込み失敗", err)
		os.Exit(1)
	} else {
		templatestr := string(content)
		templ := template.Must(template.New("sendname").Parse(templatestr))
		templ.Execute(resw, req.FormValue("yourname"))
	}
}
