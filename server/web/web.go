package web

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed assets/*
var UIAssetsFS embed.FS

func StaticFilesHandler() http.Handler {
	staticFiles, err := fs.Sub(UIAssetsFS, "assets")
	if err != nil {
		log.Fatal(err)
	}
	return http.FileServer(http.FS(staticFiles))
}
