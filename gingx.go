package main

import (
	"fmt"
	"log"
	"net/http"
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
	// FileMode
	//files, _ := filepath.Glob("*")
	//fmt.Fprintln(w, files)
}
