package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/csrf"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var (
	env = os.Getenv("ENV") // set to development if running locally

	host         = flag.String("host", "0.0.0.0", "host to listen on")
	port         = flag.Int("port", 80, "port to listen on")
	readTimeout  = flag.Duration("read_timeout", 5*time.Second, "read timeout")
	writeTimeout = flag.Duration("write_timeout", 30*time.Second, "write timeout")

	civicDataAPIKey = os.Getenv("CIVIC_DATA_API_KEY")

	decoder = schema.NewDecoder()
)

func main() {
	flag.Parse()

	r := mux.NewRouter()

	// Handler for static assets
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.Handle("/", handlers.LoggingHandler(os.Stdout, errorHandler(addressHandler))).Methods("GET")

	// POST requests without a valid csrf token will return HTTP 403
	r.Handle("/search", handlers.LoggingHandler(os.Stdout, errorHandler(searchHandler))).Methods("POST")

	var opts []csrf.Option
	if env == "development" {
		// Set to false in development, otherwise csrf token
		// will not be sent over plain HTTP
		opts = append(opts, csrf.Secure(false))
	}

	CSRF := csrf.Protect([]byte("32-byte-long-auth-key"), opts...)

	srv := &http.Server{
		Handler:      CSRF(r),
		Addr:         fmt.Sprintf("%s:%d", *host, *port),
		WriteTimeout: *writeTimeout,
		ReadTimeout:  *readTimeout,
	}

	log.Fatal(srv.ListenAndServe())
}
