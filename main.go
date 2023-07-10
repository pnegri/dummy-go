package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const INDEX = `<!DOCTYPE html>
<html>
  <head>
    <title>GitOps Testing</title>
  </head>
  <body>
	Sample Go App - Version 6
  </body>
</html>`

func getenv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}

func main() {
	fmt.Println("Version 6")
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, INDEX)
	})

	log.Fatal(http.ListenAndServe(":"+getenv("PORT","8080"), router))
}
