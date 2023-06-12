package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

type Cal int

func (cal *Cal) Square(num int, result *Result) error {
	result.Num = num
	result.Ans = num * num
	return nil
}

func main() {
	//rpc发布 Cal 中满足 RPC 注册条件的方法（Cal.Square）
	rpc.Register(new(Cal))
	rpc.HandleHTTP()
	log.Printf("Serveing PRC server on port %d", 1234)
	if err := http.ListenAndServe("localhost:1234", nil); err != nil {
		log.Fatal("Error serving:", err)
	}

}
