package httpinterface

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ryusei-takiya/ossmate/internal/usecase/fetch_trending"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/api/trending", func(c *gin.Context) {
		// 各種パラメータを取得
		language := c.Query("language")
		pageStr := c.DefaultQuery("page", "1")
		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}
		// GithubAPI呼び出し
		repos, err := fetch_trending.FetchTrendingRepos(language, page)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch repositories"})
			return
		}
		c.JSON(http.StatusOK, repos)
	})
}
