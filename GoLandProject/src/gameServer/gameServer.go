package main

import (
	"fmt"
	"net/http"
)

type Engine struct{}

//func (e Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
//	//TODO implement me
//	//panic("implement me")
//	switch request.URL.Path {
//	case "/":
//		fmt.Fprintf(writer, "my Handler")
//	case "/hello":
//		for k, v := range request.Header {
//			fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
//		}
//	default:
//		fmt.Fprintf(writer, "URL.PATH = %q\n", request.URL.Path)
//	}
//}

//func main() {
//	//open Two Route
//	//http.HandleFunc("/", indexHandler)
//	//http.HandleFunc("/hello", helloHandler)
//	engine := new(Engine)
//	log.Fatal(http.ListenAndServe("localhost:8080", engine))
//}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.PATH = %q\n", req.URL)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
