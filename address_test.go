package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestAddressFromPostForm(t *testing.T) {
	tests := []struct {
		Input  url.Values
		Expect *Address
	}{
		{
			url.Values{
				"street":   []string{"123 4th St"},
				"street-2": []string{""},
				"city":     []string{"Brooklyn"},
				"state":    []string{"NY"},
				"zip":      []string{"11201"},
			},
			&Address{
				FirstLine:  "123 4th St",
				SecondLine: "",
				City:       "Brooklyn",
				State:      state("NY"),
				Zip:        zip("11201"),
			},
		},
		{
			url.Values{
				"street":   []string{"123 4th St"},
				"street-2": []string{"Apt 3R"},
				"city":     []string{"Los Angeles"},
				"state":    []string{"CA"},
				"zip":      []string{"90210-1234"},
			},
			&Address{
				FirstLine:  "123 4th St",
				SecondLine: "Apt 3R",
				City:       "Los Angeles",
				State:      state("CA"),
				Zip:        zip("90210-1234"),
			},
		},
	}

	for _, tc := range tests {
		actual, err := AddressFromPostForm(tc.Input)
		if err != nil {
			t.Errorf("Encountered unexpected exception: %s", err)
		}
		if !reflect.DeepEqual(tc.Expect, actual) {
			t.Errorf("expected %v, got %v", tc.Expect, actual)
		}
	}
}

func TestAddressFromPostForm_Invalid(t *testing.T) {
	tests := []url.Values{
		url.Values{
			"street":   []string{"123 4th St"},
			"street-2": []string{""},
			"city":     []string{"Brooklyn"},
			"state":    []string{"notastate"},
			"zip":      []string{"11201"},
		},
		url.Values{
			"street":   []string{"123 4th St"},
			"street-2": []string{""},
			"city":     []string{"Brooklyn"},
			"state":    []string{"NY"},
			"zip":      []string{"notazip"},
		},
		url.Values{
			"street":   []string{""}, // missing first line
			"street-2": []string{"aaa"},
			"city":     []string{"Brooklyn"},
			"state":    []string{"NY"},
			"zip":      []string{"11201"},
		},
		url.Values{
			"street":   []string{"aaa"},
			"street-2": []string{"aaa"},
			"city":     []string{""}, // missing city
			"state":    []string{"NY"},
			"zip":      []string{"11201"},
		},
	}

	for _, tc := range tests {
		if _, err := AddressFromPostForm(tc); err == nil {
			// todo, would be more helpful to specify the error
			t.Errorf("Expected error")
		}

	}
}
