package main

import (
	"fmt"
	"net/http"
)

func main() {
	setuproutes()
	fmt.Println("Started Server...!!!!")
	http.ListenAndServe(":8080", nil)
}

func setuproutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})

}
