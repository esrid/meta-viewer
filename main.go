package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/keen-c/meta/ui"
	"github.com/keen-c/meta/ui/assets"
)

func main() {
	xxx()
	r := chi.NewRouter()
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(assets.FS))))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		ui.Home().Render(r.Context(), w)
	})

	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
}

func xxx() {
	cmd := exec.Command("exiftool", "./oks.html")
	var out, outerr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &outerr

	if err := cmd.Run(); err != nil {
		panic(err)
	}
	outputmap := map[string]string{}
	scanner := bufio.NewScanner(&out)
	for scanner.Scan() {
		astr := strings.Split(scanner.Text(), ":")
		outputmap[astr[0]] = astr[1]
	}
	for key, value := range outputmap {
		fmt.Printf("%s : %s \n", key, value)
	}
}
