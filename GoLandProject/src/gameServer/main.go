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
	//分组控制
	//r := geeServer.New()
	//r.GET("/index", func(c *geeServer.Context) {
	//	c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	//})
	//v1 := r.Group("/v1")
	//{
	//	v1.GET("/", func(c *geeServer.Context) {
	//		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	//	})
	//
	//	v1.GET("/hello", func(c *geeServer.Context) {
	//		// expect /hello?name=geektutu
	//		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	//	})
	//}
	//v2 := r.Group("/v2")
	//{
	//	v2.GET("/hello/:name", func(c *geeServer.Context) {
	//		// expect /hello/geektutu
	//		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	//	})
	//	v2.POST("/login", func(c *geeServer.Context) {
	//		c.JSON(http.StatusOK, geeServer.H{
	//			"username": c.PostForm("username"),
	//			"password": c.PostForm("password"),
	//		})
	//	})
	//
	//}

	//中间件使用
	//r := geeServer.New()
	////全局中间件
	//r.Use(geeServer.Logger())
	//r.GET("/", func(c *geeServer.Context) {
	//	c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	//})
	//v2 := r.Group("/v2")
	//v2.Use(onlyForV2()) // v2 group middleware
	//{
	//	v2.GET("/hello/:name", func(c *geeServer.Context) {
	//		// expect /hello/geektutu
	//		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	//	})
	//}
	//r.POST("/login",func(c *geeServer.Context))

	//html使用1
	r := geeServer.New()
	r.Use(geeServer.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")
	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/", func(c *geeServer.Context) {
		c.HTML(http.StatusOK, "css.html", nil)
	})
	r.GET("/stu", func(c *geeServer.Context) {
		c.HTML(http.StatusOK, "arr.html", geeServer.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})
	r.GET("/date", func(c *geeServer.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", geeServer.H{
			"title": "gee",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})
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
