package main

import "log"

type Result struct {
	Num, Ans int
}

type Cal int

func (cal *Cal) Square(num int) *Result {
	return &Result{
		Num: num,
		Ans: num * num,
	}
}

func main() {
	//r := gin.Default()
	//r.GET("/:name", func(c *gin.Context) {
	//	name := c.Param("name")
	//	c.String(200, "Hello，%s", name)
	//})
	//r.Run(":9191")
	cal := new(Cal)
	result := cal.Square(12)
	log.Printf("%d^2 = %d\n", result.Num, result.Ans)
}
