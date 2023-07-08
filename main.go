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
	Sample Go App - Version 1
  </body>
</html>`

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, INDEX)
	})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
