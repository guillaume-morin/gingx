package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", http_handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func http_handler(w http.ResponseWriter, r *http.Request) {
	document_root := filepath.Dir(os.Args[0])
	if len(os.Args) > 1 {
		document_root = os.Args[1]
	}
	target_file := fmt.Sprint(document_root, r.URL.Path)
	if file_stat, err := os.Stat(target_file); err == nil {
		if file_stat.IsDir() {
			files, _ := filepath.Glob(fmt.Sprint(target_file, "/*"))
			fmt.Fprintln(w, fmt.Sprint("Content of directory : ", r.URL.Path))
			for _, element := range files {
				fmt.Fprintln(w, element)
			}
		} else {
			file, _ := os.Open(target_file)
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				fmt.Fprintln(w, scanner.Text())
			}
		}
	} else {
		http.Error(w, fmt.Sprint("No such file or directory : ", r.URL.Path), http.StatusNotFound)
	}
}
