package githubapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/ryusei-takiya/ossmate/internal/domain/github"
)

// キャッシュデータ構造体
type cachedResult struct {
	data      []github.Repository
	timestamp time.Time
}

var (
	cache         = make(map[string]cachedResult) // クエリごとのキャッシュ
	cacheDuration = 30 * time.Minute
	cacheMutex    sync.Mutex
)

// 最新のリポジトリ一覧を取得
// 定期的にキャッシュを行い、存在しない場合は最新を取得
func FetchPopularRepositories(language string, page int) ([]github.Repository, error) {

	const perPage = 10

	cacheMutex.Lock()

	// キャッシュヒットしてかつ有効期限内ならキャッシュを返す
	result, found := cache[language]
	if found && time.Since(result.timestamp) < cacheDuration {
		cacheMutex.Unlock()
		fmt.Println("キャッシュから返却")

		// リクエストのページで分割する。
		start := (page - 1) * perPage
		end := start + perPage
		if start >= len(result.data) {
			return []github.Repository{}, nil
		}
		if end > len(result.data) {
			end = len(result.data)
		}
		return result.data[start:end], nil
	}
	cacheMutex.Unlock()

	// 初回 or キャッシュ切れ → GitHub APIから取得
	// GithubのAPLにクエリパラメータ（抽出条件）をセット
	q := "stars:>0"
	if language != "" {
		q = "language:" + language
	}
	url := fmt.Sprintf(
		"https://api.github.com/search/repositories?q=%s&sort=stars&order=desc&per_page=%d&page=%d",
		q, 100, 1, // ←ここで最大100件一括取得
	)

	// 取得処理
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	var apiResult github.SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&apiResult); err != nil {
		return nil, err
	}

	// キャッシュ更新
	cacheMutex.Lock()
	cache[language] = cachedResult{
		data:      apiResult.Items,
		timestamp: time.Now(),
	}
	cacheMutex.Unlock()

	start := (page - 1) * perPage
	end := start + perPage
	if start >= len(apiResult.Items) {
		return []github.Repository{}, nil
	}

	if end > len(apiResult.Items) {
		end = len(apiResult.Items)
	}

	fmt.Println("GitHub APIから新規取得")
	return apiResult.Items[start:end], nil
}
