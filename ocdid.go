package main

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

const (
	turboVoteBaseURL = "https://api.turbovote.org"
)

var (
	turboVoteTimeout = 5 * time.Second
	re               = regexp.MustCompile(`\s+`)
)

// TurboVoteClient turbovote api client
type TurboVoteClient struct {
	baseURL    string
	httpClient *http.Client
}

// NewTurboVoteClient creates a new, initialized turboVoteClient
func NewTurboVoteClient() *TurboVoteClient {
	return &TurboVoteClient{
		baseURL:    turboVoteBaseURL,
		httpClient: &http.Client{Timeout: turboVoteTimeout},
	}
}

// UpcomingElectionsData represents the response from the turbovote upcoming elections endpoint
type UpcomingElectionsData struct {
}

// GetUpcomingElections gets upcoming elections for given address from the turbovote api
func (tc *TurboVoteClient) GetUpcomingElections(a *Address) (*UpcomingElectionsData, error) {
	ocdids := makeOCDIDs(a)
	url := fmt.Sprintf("%s/elections/upcoming?district-divisions=%s", tc.baseURL, strings.Join(ocdids, ","))

	_, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func makeOCDIDs(a *Address) []string {
	var ids []string

	// state only ocdid
	// state input has been validated against our list of states
	stateID := fmt.Sprintf("ocd-division/country:us/state:%s", strings.ToLower(string(a.State)))

	// city ocdid
	// city has only been validated to be a string, add path escaping to avoid malicious input
	// e.g. city = Cleveland/evil/path...
	city := re.ReplaceAllString(strings.ToLower(a.City), "_")
	cityID := fmt.Sprintf("%s/place:", url.PathEscape(city))

	return append(ids, stateID, cityID)
}
