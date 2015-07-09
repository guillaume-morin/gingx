package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	//"strings"
	//"path/filepath"
)

func main() {
	http.HandleFunc("/", http_handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func http_handler(w http.ResponseWriter, r *http.Request) {
	document_root := "/home/gmorin/git/gingx"
	target_file := fmt.Sprint(document_root, r.URL.Path)
	fmt.Fprintln(w, fmt.Sprint("Requested file : ", target_file))
	if file_stat, err := os.Stat(target_file); err == nil {
		fmt.Fprintln(w, file_stat)
	} else {
		fmt.Fprintln(w, "404 File not found")
	}
	// FileMode
	//files, _ := filepath.Glob("*")
	//fmt.Fprintln(w, files)
}
