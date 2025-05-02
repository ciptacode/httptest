package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/yuin/goldmark"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == "/" {
		md, err := os.ReadFile("README.md")
		if err != nil {
			http.Error(w, "README.md not found", http.StatusInternalServerError)
			return
		}

		var buf bytes.Buffer
		if err := goldmark.Convert(md, &buf); err != nil {
			http.Error(w, "Failed to render markdown", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write(buf.Bytes())
		return
	}

	statusStr := path[1:]
	statusCode, err := strconv.Atoi(statusStr)
	if err != nil || statusCode < 100 || statusCode > 599 {
		http.Error(w, "Invalid status code", http.StatusBadRequest)
		return
	}

	sleepStr := r.URL.Query().Get("sleep")
	if sleepStr != "" {
		if sleepMs, err := strconv.Atoi(sleepStr); err == nil && sleepMs > 0 {
			time.Sleep(time.Duration(sleepMs) * time.Millisecond)
		}
	}

	statusText := http.StatusText(statusCode)
	if statusText == "" {
		statusText = "Unknown Status"
	}

	w.WriteHeader(statusCode)
	fmt.Fprintf(w, "%d %s\n", statusCode, statusText)
}

func main() {
	port := flag.String("p", "8080", "Port to listen on")
	flag.Parse()

	http.HandleFunc("/", handler)

	log.Printf("httptest running at http://localhost:%s", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
