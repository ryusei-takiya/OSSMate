package main

import (
	"github.com/gin-gonic/gin"
	httpinterface "github.com/ryusei-takiya/ossmate/internal/interface/http"
)

func main() {

	router := gin.Default()
	httpinterface.RegisterRoutes(router)

	// フロントエンド読み込み
	router.Static("/web", "../../web")

	router.Run(":8080")

}
