package main

import (
	"net/http"
	"path/filepath"
)

func index(w http.ResponseWriter, r *http.Request) {
	if len(props["swagger.path"]) > 0 {
		http.Redirect(w, r, props["swagger.path"], http.StatusMovedPermanently)
	}
}

func icon(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, filepath.Join(props["swagger.path"], "images/mora.ico"), http.StatusMovedPermanently)
}
