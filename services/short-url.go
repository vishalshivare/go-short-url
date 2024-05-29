package services

import (
	"fmt"
	"go-short-url/models"
	"go-short-url/utils"
	"sort"
)

type ShortenURLService struct {
	urlToShort map[string]string
	shortToUrl map[string]string
	count      map[string]int
}

type ShortenURLServicer interface {
	GetOrCreate(domain, url string) (string, error)
	Metrics() []models.RequestURL
	GetLongURL(shortURL string) (string, error)
}

func NewShortenURLServicer() ShortenURLServicer {
	return &ShortenURLService{
		urlToShort: make(map[string]string),
		shortToUrl: make(map[string]string),
		count:      make(map[string]int),
	}
}

func (s *ShortenURLService) GetOrCreate(domain, long_url string) (string, error) {
	if shortURL, exists := s.urlToShort[long_url]; exists {
		s.count[domain]++
		return shortURL, nil
	}

	shortURL := utils.GenerateShortURL(long_url)
	s.shortToUrl[shortURL] = long_url
	s.urlToShort[long_url] = shortURL
	s.count[domain]++

	return shortURL, nil
}

func (s *ShortenURLService) Metrics() []models.RequestURL {
	var domainList []models.RequestURL
	for domain, count := range s.count {
		domainList = append(domainList, models.RequestURL{URL: domain, Hits: count})
	}

	// Sort by count descending
	sort.Slice(domainList, func(i, j int) bool {
		return domainList[i].Hits > domainList[j].Hits
	})

	if len(domainList) <= 3 {
		return domainList
	}

	return domainList[:3]
}

func (s *ShortenURLService) GetLongURL(shortURL string) (string, error) {
	longURL, ok := s.shortToUrl[shortURL]
	if !ok {
		return "", fmt.Errorf("url not found")
	}
	return longURL, nil
}
