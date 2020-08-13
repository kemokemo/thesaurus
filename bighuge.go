package thesaurus

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

const baseURL = "https://words.bighugelabs.com/api/2/"

// BigHuge is the struct to use the API of the 'Big Huge Thesaurus'.
type BigHuge struct {
	APIKey string
}

type synonyms struct {
	Noun *words `json:"noun"`
	Verb *words `json:"verb"`
}

type words struct {
	Syn []string `json:"syn"`
}

// Synonyms returns the synonyms of the 'term' from the 'Big Huge Thesaurus'.
// If there is no synonyms, this function returns the error with the message of "http stasus is '404 Not Found'".
func (b *BigHuge) Synonyms(term string) ([]string, error) {
	var syns []string
	if b.APIKey == "" {
		return syns, fmt.Errorf("bighuge: the API Key is empty")
	}

	URL, err := url.Parse(baseURL)
	if err != nil {
		return syns, fmt.Errorf("bighuge: failed to load the base URL: %v", err)
	}
	URL.Path = path.Join(URL.Path, b.APIKey, term, "json")

	response, err := http.Get(URL.String())
	if err != nil {
		return syns, fmt.Errorf("bighuge: failed to get the synonyms of the %q: %v", term, err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return syns, fmt.Errorf("bighuge: http status is '%v'", response.Status)
	}

	var data synonyms
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return syns, fmt.Errorf("bighuge: failed to decode the json: %v", err)
	}

	if data.Noun != nil {
		syns = append(syns, data.Noun.Syn...)
	}
	if data.Verb != nil {
		syns = append(syns, data.Verb.Syn...)
	}

	return syns, nil
}
