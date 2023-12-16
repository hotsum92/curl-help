package main

import (
  "fmt"
  "log"
  "net/http"
  "net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
  dump, err := httputil.DumpRequest(r, true)

  if err != nil {
    http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
    return
  }
  fmt.Println(string(dump))

  if r.URL.String() == "/old-place" {
    w.Header().Set("Location", "/new-place")
    w.WriteHeader(http.StatusMovedPermanently)
    fmt.Fprintf(w, "<html><body>old place</body></html>\n")
    return
  }

  if r.URL.String() == "/new-place" {
    fmt.Fprintf(w, "<html><body>new place</body></html>\n")
    return
  }

  http.NotFound(w, r)
}

func main() {
  var httpServer http.Server
  http.HandleFunc("/", handler)
  log.Println("start http listening :18888")
  httpServer.Addr = ":18888"
  log.Println(httpServer.ListenAndServe())
}
