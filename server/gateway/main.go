package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		var backend *url.URL
		var err error

		if strings.HasPrefix(path, "/v1/cars") {
			backend, err = url.Parse("http://localhost:4000")
		} else if strings.HasPrefix(path, "/v1/order") {
			backend, err = url.Parse("http://localhost:4001")
		} else if strings.HasPrefix(path, "/v1/users") {
			backend, err = url.Parse("http://localhost:8080")
		} else {
			http.NotFound(w, r)
			return
		}

		if err != nil {
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(backend)

		r.URL.Host = backend.Host
		r.URL.Scheme = backend.Scheme
		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
		r.Host = backend.Host

		proxy.ServeHTTP(w, r)
	})

	log.Println("API Gateway started at :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
