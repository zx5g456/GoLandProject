package main

import (
	"fmt"
	"geeServer"
	"html/template"
	"log"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}
func main() {
	//html使用1
	r := geeServer.New()
	r.Use(geeServer.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")
	r.GET("/ball", func(c *geeServer.Context) {
		c.HTML(http.StatusOK, "index.html", geeServer.H{
			"title": "ballGame",
		})
	})

	r.Run(":8080")

}

func onlyForV2() geeServer.HandlerFunc {
	return func(c *geeServer.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		//c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
