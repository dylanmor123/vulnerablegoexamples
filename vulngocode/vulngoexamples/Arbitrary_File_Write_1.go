package main

import (
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")
	data := r.URL.Query().Get("data")
	filedata := []byte(data)
	ioutil.WriteFile(file, filedata, 0644)
	println("")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
