package main

import (
	"io"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")
	data, err := ioutil.ReadFile(file)
	if err != nil {
		output := "File reading error"
		io.WriteString(w, output)
		return
	}
	io.WriteString(w, string(data))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
