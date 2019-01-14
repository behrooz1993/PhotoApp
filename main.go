package main

import (
  "fmt"
  "net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  if r.URL.Path == "/" {
    fmt.Fprint(w, "<h1>Home page</h1>")
  }else {
    fmt.Fprint(w, "<h1>Some other page</h1>")
  } 
}

func main() {
  http.HandleFunc("/", handlerFunc)
  http.ListenAndServe(":3000",nil)
}
