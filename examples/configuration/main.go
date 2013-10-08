package main

import (
  "fmt"
  "net/http"
  "github.com/pilu/traffic"
)

func rootHandler(w traffic.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "%s<br />", w.GetVar("foo"))

  // run with TRAFFIC_ENV=production to get the "bar" value
  // from the production section of the config file
  fmt.Fprintf(w, "%s<br />", w.GetVar("bar"))
}

func main() {
  router := traffic.New()
  // Routes
  router.Get("/", rootHandler)
  router.Run()
}
