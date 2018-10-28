package main
import (
  "log"
  "fmt"
  "net/http"
  "strings"
)

func name(w http.ResponseWriter, r *http.Request) {
  path := r.URL.Path
  name := strings.Split(path, "/")[2]
  message := strings.ToUpper(name)
  w.Write([]byte(message))
}

func Log(handler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
    handler.ServeHTTP(w, r)
  })
}

func main() {
  fmt.Println("Name Service started!")
  http.HandleFunc("/name/", name)
  if err := http.ListenAndServe(":8080", Log(http.DefaultServeMux)); err != nil {
    panic(err)
  }
}