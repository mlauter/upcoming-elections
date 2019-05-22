/*
This file contains structs representing json responses the turbovote api
Structs generated with the help of https://github.com/mholt/json-to-go

*/
package main

import (
	"time"
)

// Represents a response from the turbovote upcoming elections api
type UpcomingElectionsData []struct {
	DistrictDivisions []struct {
		OcdID                           string `json:"ocd-id"`
		VoterRegistrationAuthorityLevel string `json:"voter-registration-authority-level"`
		ElectionAuthorityLevel          string `json:"election-authority-level"`
		VotingMethods                   []struct {
			Primary        bool   `json:"primary"`
			Type           string `json:"type"`
			ExcuseRequired bool   `json:"excuse-required"`
			Instructions   struct {
				VotingID string `json:"voting-id"`
			} `json:"instructions,omitempty"`
			BallotRequestDeadlineReceived time.Time `json:"ballot-request-deadline-received,omitempty"`
			AcceptableForms               []struct {
				Name string `json:"name"`
			} `json:"acceptable-forms,omitempty"`
		} `json:"voting-methods"`
		VoterRegistrationMethods []struct {
			// seems like Instructions always shows up either with registration or Idnumber+Signature
			// depending on whether it has the DeadlinePostmarked or DeadlineOnline field
			// but trying to keep things a bit flexible here
			Instructions struct {
				Registration string `json:"registration,omitempty"`
				Idnumber     string `json:"idnumber,omitempty"`
				Signature    string `json:"signature,omitempty"`
			} `json:"instructions,omitempty"`
			Type               string    `json:"type"`
			SupportsIframe     bool      `json:"supports-iframe,omitempty"`
			DeadlineOnline     time.Time `json:"deadline-online,omitempty"`
			URL                string    `json:"url,omitempty"`
			DeadlinePostmarked time.Time `json:"deadline-postmarked,omitempty"`
			AcceptableForms    []struct {
				Name string `json:"name"`
			} `json:"acceptable-forms,omitempty"`
		} `json:"voter-registration-methods"`
		PrimaryVotingMethodSource string `json:"primary-voting-method-source"`
	} `json:"district-divisions"`
	Website                  string    `json:"website,omitempty"`
	PollingPlaceURL          string    `json:"polling-place-url"`
	Date                     time.Time `json:"date"`
	Population               int       `json:"population,omitempty"`
	PollingPlaceURLShortened string    `json:"polling-place-url-shortened"`
	Description              string    `json:"description"`
	ID                       string    `json:"id"`
}
