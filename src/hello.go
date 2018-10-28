package main
import (
  "fmt"
  "log"
  "net/http"
  "io/ioutil"
  "strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
  path := r.URL.Path
  name := strings.Split(path, "/")[2]
  resp, err := http.Get(fmt.Sprintf("http://localhost:8080/name/%s", name))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
  message := fmt.Sprintf("Hello %s", body)
  w.Write([]byte(message))
}

func Log(handler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
    handler.ServeHTTP(w, r)
  })
}

func main() {
  fmt.Println("Hello Service started!")
  http.HandleFunc("/hello/", hello)
  if err := http.ListenAndServe(":8090", Log(http.DefaultServeMux)); err != nil {
    panic(err)
  }
}