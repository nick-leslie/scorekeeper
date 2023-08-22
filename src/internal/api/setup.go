package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nick-leslie/scorekeeper/internal/api/playerapi"
)

func SetupApi(router *gin.Engine) {
	router.GET("/api/players", playerapi.GetPlayerRequest)
	router.POST("/api/players", playerapi.AddPlayerRequest)
}
