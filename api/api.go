package main

import (
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func videoHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	// Handle preflight OPTIONS request
	if r.Method == http.MethodOptions {
		return
	}

	file := r.URL.Query().Get("file")
	if file != "" {
		http.ServeFile(w, r, "videos/"+file)
	} else {
		http.Error(w, "File not found", http.StatusNotFound)
	}
}

func main() {
	http.HandleFunc("/video", videoHandler)
	http.ListenAndServe(":8080", nil)
}
