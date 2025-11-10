package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

type staticHandler struct {
	Port         int     `json:"port"`
	Path         string  `json:"path"`
	IndexFile    *string `json:"indexFile"`
	NotfoundFile *string `json:"notfoundFile"`
}

func newStaticHandler(configPath string) *staticHandler {
	fileContent, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	handler := new(staticHandler)
	json.Unmarshal(fileContent, &handler)
	return handler
}

func (h staticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(h.Path, r.URL.Path)

	stat, err := os.Stat(path)
	if err == nil && !stat.IsDir() {
		http.FileServer(http.Dir(path)).ServeHTTP(w, r)
	}

	redirectPath := filepath.Join(h.Path, *h.IndexFile)
	if h.NotfoundFile != nil {
		redirectPath = filepath.Join(h.Path, *h.NotfoundFile)
	}

	http.ServeFile(w, r, redirectPath)
}

func (h staticHandler) GetServerAddress() string {
	return fmt.Sprintf(":%d", h.Port)
}

func main() {
	staticHandler := newStaticHandler("config.json")

	http.Handle("/", staticHandler)

	addr := staticHandler.GetServerAddress()
	println("server start on address: " + addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
