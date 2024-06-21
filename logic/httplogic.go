package logic

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/keen-c/meta/ui"
)

const (
	// 500 MG
	maxSize int64 = 500 * 1024 * 1024
)

func FileUpload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(maxSize); err != nil {
		http.Error(w, "Status Internal Server Error", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Status Internal Server Error", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	tf, err := os.CreateTemp("", header.Filename)

	defer tf.Close()

	filebytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Status Internal Server Error", http.StatusInternalServerError)
		return
	}

	if _, err := tf.Write(filebytes); err != nil {
		http.Error(w, "Status Internal Server Error", http.StatusInternalServerError)
		return
	}

	data, err := READATA(tf.Name())
	if err != nil {
		http.Error(w, "Status Internal Server Error", http.StatusInternalServerError)
		return
	}

	ui.DataList(data).Render(r.Context(), w)
}

func Options(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "test")
	fmt.Println(param)
}
