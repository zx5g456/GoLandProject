package geeServer

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		//start timer
		t := time.Now()
		//Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
