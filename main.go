package main

import (
	"fmt"
	"gee"
	"net/http"
)

/*
*
curl http://127.0.0.1:8081/

	URL.Path = "/"

curl http://127.0.0.1:8081/hello

	header["User-Agent"] = ["curl/7.79.1"]
	header["Accept"] = ["* /*"]

curl http://127.0.0.1:8081/world

	404 NOT FOUND: /world
*/
func main() {
	web := gee.New()
	web.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})
	web.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "header[%q] = %q\n", k, v)
		}
	})
	web.Run(":8081")
}
