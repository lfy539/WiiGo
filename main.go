package main

import (
	"WiiGo/luna"
	"net/http"
)

type student struct {
	Name string
	Age  int
}

func main() {
	r := luna.Default()
	r.GET("/", func(c *luna.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *luna.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})
	r.Run(":9999")

}
