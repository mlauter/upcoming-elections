package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
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

func addressHandler(w http.ResponseWriter, r *http.Request) *appError {
	buf := &bytes.Buffer{}

	err := parseTemplate("index.html").Execute(buf, map[string]interface{}{
		"Title":          "Elections",
		"States":         states,
		csrf.TemplateTag: csrf.TemplateField(r),
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

func searchHandler(w http.ResponseWriter, r *http.Request) *appError {
	if err := r.ParseForm(); err != nil {
		return &appError{
			Error:   err,
			Message: "Something went wrong",
			Code:    500,
		}
	}

	addr, err := AddressFromPostForm(r.PostForm)
	if err != nil {
		return &appError{
			Error:   err,
			Message: fmt.Sprintf("Invalid address - %s", err.Error()),
			Code:    400,
		}
	}

	res, clientErr := NewTurboVoteClient().GetUpcomingElections(addr)
	if clientErr != nil {
		return &appError{
			Error:   clientErr,
			Message: "Sorry, we were unable to lookup upcoming elections for your address at this time. Please try again.",
			Code:    502,
		}
	}

	buf := &bytes.Buffer{}

	parseErr := parseTemplate("elections.html").Execute(buf, map[string]interface{}{
		"Title":     "Elections",
		"Elections": res,
	})

	if parseErr != nil {
		return &appError{
			Error:   parseErr,
			Message: "Something went wrong",
			Code:    500,
		}
	}

	buf.WriteTo(w)
	return nil
}
