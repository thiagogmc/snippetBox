package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Use the http.NewServeMux() function to initialize a new servemux, then
	//register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	// Use the http.ListenAndServe() function to start a new web server. We pas
	//two parameters: the TCP network address to listen on (in this case ":400
	//and the servemux we just created. If http.ListenAndServe() returns an er
	//we use the log.Fatal() function to log the error message and exit.
	infoLog.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
}
