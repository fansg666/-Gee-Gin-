package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func indexHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.PATH = %q\n", req.URL.Path)
}
func helloHandle(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
func panxiaolong(w http.ResponseWriter, req *http.Request) {
	//2.解析模板
	t, err := template.ParseFiles("http/base1/baidu.tmpl")
	if err != nil {
		fmt.Println("Parse template failed,err:%v", err)
		return
	}
	//3.渲染模板
	nome := "潘小龙"
	err = t.Execute(w, nome)
	if err != nil {
		fmt.Println("render template wfailed,err:%v", err)
		return
	}
}

func main() {
	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/hello", helloHandle)
	http.HandleFunc("/baidu", panxiaolong)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
