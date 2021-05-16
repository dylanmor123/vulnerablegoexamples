package main

import (
	"bufio"
	"fmt"
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

	reader := bufio.NewReader(f)
	b := make([]byte, 3)
	var data string
	for {
		n, err := reader.Read(b)
		if err != nil {
			fmt.Println("", err)
			break
		}
		data = data + string(b[0:n])
	}
	io.WriteString(w, data)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
