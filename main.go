package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Home page</h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>404</h1>")
	}
}

func main() {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", mux)
}
