package logic

import (
	"fmt"
	"net/http"
)

const (
	// 500 MG
	maxSize int64 = 500 * 1024 * 1024
)

func FileUpload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(maxSize); err != nil {
		fmt.Printf("multipart : %s", err)
		http.Error(w, "Status Internal Server Error", http.StatusInternalServerError)
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Printf("file : %s", err)
		http.Error(w, "Status Internal Server Error", http.StatusInternalServerError)
		return
	}
	println("and here")
	fmt.Printf("%+v", file)

	defer file.Close()
}
