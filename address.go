package main

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

// Regex for zip and zip+4 codes
var zipRegex = regexp.MustCompile(`^\d{5}(?:-\d{4})?$`)

// Address - a valid address
type Address struct {
	FirstLine  string `schema:"street,required"`
	SecondLine string `schema:"street-2"`
	City       string `schema:"city,required"`
	State      state  `schema:"state,required"`
	Zip        zip    `schema:"zip,required"`
}

// State - a US state
// might be useful at some point to have a more generic Field with Validation
type state string

func (st state) validate() error {
	if !stringInSlice(string(st), states) {
		return fmt.Errorf("%s is not a US state", st)
	}

	return nil
}

type zip string

// Validate returns an error if Zip is invalid
func (z zip) validate() error {
	if matched := zipRegex.MatchString(string(z)); !matched {
		return fmt.Errorf("%s is not a valid zip code", z)
	}

	return nil
}

// AddressFromPostForm creates an address struct from posted data
func AddressFromPostForm(d url.Values) (*Address, error) {
	var addr Address

	// trim space from all inputs before decoding
	for k, v := range d {
		for i, s := range v {
			d[k][i] = strings.TrimSpace(s)
		}
	}

	decoder.IgnoreUnknownKeys(true)
	if err := decoder.Decode(&addr, d); err != nil {
		return nil, err
	}

	// Do some field validation
	if err := addr.State.validate(); err != nil {
		return nil, err
	}

	if err := addr.Zip.validate(); err != nil {
		return nil, err
	}

	return &addr, nil
}
