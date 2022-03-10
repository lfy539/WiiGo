package luna

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(ctx *Context) {
		// start timer
		t := time.Now()
		//Process request
		ctx.Next()
		// Cal resolution time
		log.Printf("[%d] %s in %v", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))

	}
}
