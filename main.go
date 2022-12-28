package main

import (
	"gee"
	"net/http"
)

/*
*
curl http://127.0.0.1:8081/

	<h1>Hello Gee</h1>

curl http://127.0.0.1:8081/hello\?name\=gyz

	hello gyz, you're at /hello

curl http://127.0.0.1:8081/world

	404 NOT FOUND: /world

curl http://127.0.0.1:8081/login -X POST -d 'username=gyz&password=3344'

	{"password":"3344","username":"gyz"}
*/
func main() {
	web := gee.New()
	web.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>\n")
	})
	web.GET("/hello", func(c *gee.Context) {
		c.TextString(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	web.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	_ = web.Run(":8081")
}
