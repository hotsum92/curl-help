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

  if r.URL.String() == "/cookie" {
    w.Header().Add("Set-Cookie", "VISIT=TRUE")
    if _, ok := r.Header["Cookie"]; ok {
    // クッキーがあるということは一度来たことがある人
    fmt.Fprintf(w, "<html><body>2 回目以降 </body></html>\n")
    } else {
    fmt.Fprintf(w, "<html><body> 初訪問 </body></html>\n")
    }

    return
  }

  fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}

func main() {
  var httpServer http.Server
  http.HandleFunc("/", handler)
  log.Println("start http listening :18888")
  httpServer.Addr = ":18888"
  log.Println(httpServer.ListenAndServe())
}
