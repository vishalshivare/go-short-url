package utils

import "net/url"

func GetDomain(rawUrl string) (string, error) {
	parsedURL, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	return parsedURL.Hostname(), nil
}
