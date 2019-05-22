package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

// GetUpcomingElections gets upcoming elections for given address from the turbovote api
func (tc *TurboVoteClient) GetUpcomingElections(a *Address) (*UpcomingElectionsData, error) {
	ocdids := makeOCDIDs(a)
	url := fmt.Sprintf("%s/elections/upcoming?district-divisions=%s", tc.baseURL, strings.Join(ocdids, ","))

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// add json header
	req.Header.Add("Accept", "application/json")

	// make the request
	res, getErr := tc.httpClient.Do(req)
	if getErr != nil {
		return nil, getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	elections := UpcomingElectionsData{}
	jsonErr := json.Unmarshal(body, &elections)
	if err != nil {
		return nil, jsonErr
	}

	return &elections, nil
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
	cityID := fmt.Sprintf("%s/place:%s", stateID, url.PathEscape(city))

	return append(ids, stateID, cityID)
}
