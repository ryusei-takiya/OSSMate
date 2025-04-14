package githubapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ryusei-takiya/ossmate/internal/domain/github"
)

func FetchPopularRepositories(language string, page int) ([]github.Repository, error) {
	// GithubのAPLにクエリパラメータ（抽出条件）をセット
	q := "stars:>0"
	if language != "" {
		q = "language:" + language
	}
	url := fmt.Sprintf(
		"https://api.github.com/search/repositories?q=%s&sort=stars&order=desc&per_page=%d&page=%d",
		q, 10, page,
	)

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
