package main

import (
	"io"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
)

func main() {
	log.Println("Starting Kubernetes Application.")
	ws := new(restful.WebService)
	ws.Route(ws.GET("/hello").To(hello))
	ws.Route(ws.GET("/health").To(health))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "hello")
}

func health(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "OK")
}
