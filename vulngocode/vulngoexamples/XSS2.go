package main

import (
	"net/http"
	"text/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
	test := r.URL.Query().Get("test")
	tmpl := template.New("hello")
	tmpl, _ = tmpl.Parse(`{{define "T"}}{{.}}{{end}}`)
	tmpl.ExecuteTemplate(w, "T", test)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
