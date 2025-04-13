package fetch_trending

import (
	"github.com/ryusei-takiya/ossmate/internal/domain/github"
	"github.com/ryusei-takiya/ossmate/internal/infrastructure/githubapi"
)

func FetchTrendingRepos(language string) ([]github.Repository, error) {
	return githubapi.FetchPopularRepositories(language)
}
