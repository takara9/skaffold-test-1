package main
 
import (
  "log"
  "net/http"
)
 
func main() {
  http.Handle("/",http.StripPrefix("/",http.FileServer(http.Dir("/html"))))
  log.Fatal(http.ListenAndServe(":8080", nil))
}
