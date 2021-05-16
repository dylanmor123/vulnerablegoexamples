package main

import (
	"bufio"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")
	f, err := os.Open(file)
	if err != nil {
		println("")
	}
	var data string
	s := bufio.NewScanner(f)
	for s.Scan() {
		data = data + s.Text() + "\n"
	}
	io.WriteString(w, data)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
