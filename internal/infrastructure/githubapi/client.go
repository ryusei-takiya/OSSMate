package githubapi

import (
	"encoding/json"
	"net/http"

	"github.com/ryusei-takiya/ossmate/internal/domain/github"
)

func FetchPopularRepositories(language string) ([]github.Repository, error) {
	// 例: /api/trending?language=go
	q := "stars:>1000"
	if language != "" {
		q += "+language:" + language
	}
	url := "https://api.github.com/search/repositories?q=" + q + "&sort=stars&order=desc&per_page=20"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	var result github.SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Items, nil
}
