package httpinterface

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryusei-takiya/ossmate/internal/usecase/fetch_trending"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/api/trending", func(c *gin.Context) {
		language := c.Query("language")

		repos, err := fetch_trending.FetchTrendingRepos(language)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch repositories"})
			return
		}
		c.JSON(http.StatusOK, repos)
	})
}
