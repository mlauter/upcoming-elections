package main

import (
	"bytes"
	"log"
	"net/http"
)

type appHandler func(w http.ResponseWriter, r *http.Request) *appError

type appError struct {
	Error   error
	Message string
	Code    int
}

func errorHandler(f appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if e := f(w, r); e != nil {
			log.Printf("Handler error: status code: %d, message: %s, underlying err: %s",
				e.Code, e.Message, e.Error)

			http.Error(w, e.Message, e.Code)
		}
	}
}

type addressData struct {
	Title  string
	States []string
}

func addressHandler(w http.ResponseWriter, r *http.Request) *appError {
	buf := &bytes.Buffer{}
	err := parseTemplate("index.html").Execute(buf, &addressData{
		Title:  "Elections",
		States: states,
	})

	if err != nil {
		return &appError{
			Error:   err,
			Message: "Something went wrong",
			Code:    500,
		}
	}

	buf.WriteTo(w)
	return nil
}
