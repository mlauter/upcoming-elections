package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/api/civicinfo/v2"
	"google.golang.org/api/option"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	turboVoteBaseURL = "https://api.turbovote.org"
)

// todo make configurable
var (
	turboVoteTimeout = 5 * time.Second

	_ OCDIDGetter = (*CivicInfoClient)(nil) // ensure civic info client fulfills the interface
)

// TurboVoteClient turbovote api client
type TurboVoteClient struct {
	baseURL    string
	httpClient *http.Client
}

// NewTurboVoteClient creates a new, initialized TurboVoteClient
func NewTurboVoteClient() *TurboVoteClient {
	return &TurboVoteClient{
		baseURL:    turboVoteBaseURL,
		httpClient: &http.Client{Timeout: turboVoteTimeout},
	}
}

// GetUpcomingElections gets upcoming elections for given address from the turbovote api
func (tc *TurboVoteClient) GetUpcomingElections(a *Address, o OCDIDGetter) (*UpcomingElectionsData, error) {
	ocdids, _ := o.GetOCDIDs(a)
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
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &elections, nil
}

// OCDIDGetter wrapper around google civic info service for fetching ocdids
type OCDIDGetter interface {
	GetOCDIDs(a *Address) ([]string, error)
}

// CivicInfoClient receiver for GetOCDID method
type CivicInfoClient struct {
}

// GetOCDIDs gets ocdids for a given address from the google civic info api
func (c *CivicInfoClient) GetOCDIDs(a *Address) ([]string, error) {
	var ids []string
	addrString := a.ToString()

	ctx := context.Background()
	civicinfoService, err := civicinfo.NewService(ctx, option.WithAPIKey(civicDataAPIKey))
	if err != nil {
		return ids, err
	}

	req := &civicinfo.RepresentativeInfoRequest{}
	call := civicinfoService.Representatives.RepresentativeInfoByAddress(req)
	res, _ := call.Address(addrString).Context(ctx).Do()

	for k := range res.Divisions {
		ids = append(ids, k)
	}

	return ids, nil
}
