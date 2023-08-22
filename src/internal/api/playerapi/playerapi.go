package playerapi

import (
	"github.com/gin-gonic/gin"
	"github.com/nick-leslie/scorekeeper/internal/dbMan"
)

func AddPlayerRequest(context *gin.Context) {
	var newPlayer dbMan.Player
	err := context.BindJSON(&newPlayer)
	if err != nil {
		return
	}
	rowsAdded, dbErr := dbMan.AddPlayerToDb(newPlayer.Name)
	if dbErr != nil {
		return
	}
	context.JSON(200, gin.H{
		"rows added": rowsAdded,
	})
}
func GetPlayerRequest(context *gin.Context) {
	players := dbMan.GetPlayers()
	context.JSON(200, players)
}
