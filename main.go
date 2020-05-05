package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"

	"github.com/daxinc/cowsayweb/cowsay"
	"github.com/gorilla/mux"
)

var healthy = true

func main() {
	port := getPort()
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/health", healthHandler)

	r.NotFoundHandler = http.HandlerFunc(notFound)

	log.Printf("Server listening on %s ... ", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	vars:=mux.Vars(r)
	quote := vars["quote"]
	text := cowsay.Say(quote)
	fmt.Fprintf(w, cowsay.IndexHTML, html.EscapeString(quote), html.EscapeString(text))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, `
	<!doctype html>
	<html lang="en">
	<head>
	<meta charset="utf-8">
	<title>Page Not Found - Cow Say</title>
	</head>
	
	<body>
		<h1>Page Not found.</h1>

		<a href="/">Go to Home Page</a>
	</body>
	</html>`)
}

func getPort() string {
	port := "8080"
	args := os.Args
	if len(args) > 1 {
		port = args[1]
	}

	return port
}