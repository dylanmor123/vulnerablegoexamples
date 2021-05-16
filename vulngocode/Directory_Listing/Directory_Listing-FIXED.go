package main

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./tmp/static")})
	mux.Handle("tmp/static", http.NotFoundHandler())
	mux.Handle("/tmp/static/", http.StripPrefix("/tmp/static", fileServer))

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		//line below needed for cross-platform windows to unix
		index = strings.Replace(index, "\\", "/", -1)
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}
