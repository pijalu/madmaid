package main

import (
	"embed"
	"flag"
	"io/fs"
	"log"
	"net/http"
)

//go:embed all:dist
var webAssets embed.FS

func main() {
	addr := flag.String("addr", "127.0.0.1:8080", "address to listen on (use 0.0.0.0:8080 to expose on your network)")
	flag.Parse()
	content, err := fs.Sub(webAssets, "dist")
	if err != nil {
		log.Fatal(err)
	}
	files := http.FileServer(http.FS(content))
	http.Handle("/", files)
	log.Printf("Mermaid Studio listening at http://%s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
