package main

import (
  "fmt"
  "net/http"
  "github.com/pilu/traffic"
)

type PingMiddleware struct {}

// If the request path is "/ping", it writes PONG in the response and returns without calling the next middleware
// Otherwise it sets the variable "PING" with PONG as value and calls the next  middleware.
// The next middleware can
func (c *PingMiddleware) ServeHTTP(w traffic.ResponseWriter, r *http.Request, next traffic.NextMiddlewareFunc) (traffic.ResponseWriter, *http.Request) {
  if r.URL.Path == "/ping" {
    fmt.Fprint(w, "pong\n")

    return w, r
  }

  if nextMiddleware := next(); nextMiddleware != nil {
    w.SetVar("ping", "pong")
    w, r = nextMiddleware.ServeHTTP(w, r, next)
  }

  return w, r
}

func root(w traffic.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Router var foo: %v.\n", w.GetVar("foo"))
  fmt.Fprintf(w, "Middleware var ping: %v\n", w.GetVar("ping"))
}

func main() {
  t := traffic.New()
  // Add PingMiddleware
  t.AddMiddleware(&PingMiddleware{})
  // Set router var "foo"
  t.SetVar("foo", "bar")
  // Add root handler
  t.Get("/", root)

  t.Run()
}
