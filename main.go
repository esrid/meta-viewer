package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/keen-c/meta/ui"
	"github.com/keen-c/meta/ui/assets"
)

func main() {
	r := chi.NewRouter()
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(assets.FS))))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		ui.Home().Render(r.Context(), w)
	})

	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
}
