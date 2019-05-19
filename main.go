package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	host         = flag.String("host", "0.0.0.0", "host to listen on")
	port         = flag.Int("port", 80, "port to listen on")
	readTimeout  = flag.Duration("read_timeout", 15*time.Second, "read timeout")
	writeTimeout = flag.Duration("write_timeout", 15*time.Second, "write timeout")
)

func main() {
	flag.Parse()

	r := mux.NewRouter()

	// Handler for static assets
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.Handle("/", handlers.LoggingHandler(os.Stdout, errorHandler(addressHandler))).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%d", *host, *port),
		WriteTimeout: *writeTimeout,
		ReadTimeout:  *readTimeout,
	}

	log.Fatal(srv.ListenAndServe())
}
