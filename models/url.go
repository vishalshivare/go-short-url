package models

import "fmt"

type RequestURL struct {
	URL  string `json:"url"`
	Hits int    `json:"hits"`
}

func (r *RequestURL) Validate() error {
	if r == nil {
		return fmt.Errorf("missing required fields")
	}
	if len(r.URL) == 0 {
		return fmt.Errorf("field 'URL' can not be empty")
	}
	return nil
}

type ListURLs struct {
	URLs []RequestURL `json:"urls"`
}
