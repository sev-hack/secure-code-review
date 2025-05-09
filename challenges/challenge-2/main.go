package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func viewFile(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")

	cleanPath := filepath.Clean(filename)
	fullPath := filepath.Join("data/", cleanPath)
	fmt.Println("The full path is: " + fullPath)

	data, err := os.ReadFile(fullPath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, string(data))
}

func main() {
	http.HandleFunc("/view", viewFile)
	fmt.Println("Starting server at port 8080")
	http.ListenAndServe(":8080", nil)
}
