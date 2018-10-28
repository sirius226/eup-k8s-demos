package main
import (
	"os"
	"log"
	"fmt"
	"net/http"
)

func secret(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	if r.FormValue("secret") == (os.Getenv("SECRET")) {
		w.Write([]byte("Good credential"))
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Bad credential"))
	}
}

func Log(handler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
    handler.ServeHTTP(w, r)
  })
}

func main() {
  fmt.Println("Secret Service started!")
  http.HandleFunc("/auth", secret)
  if err := http.ListenAndServe(":8080", Log(http.DefaultServeMux)); err != nil {
    panic(err)
  }
}