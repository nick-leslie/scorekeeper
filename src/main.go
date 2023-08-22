package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/nick-leslie/scorekeeper/internal/api"
	"github.com/nick-leslie/scorekeeper/internal/dbMan"
	"github.com/nick-leslie/scorekeeper/internal/static"
)

var buildDir = "frontend/build"
var DBPath = "../score.db"

//go:embed all:frontend/build
var files embed.FS

func main() {

	router := gin.Default()
	_, dberr := dbMan.New(DBPath)
	if dberr != nil {
		return
	}
	static.SetupStaticEmbed(router, &files, buildDir, buildDir)
	api.SetupApi(router)
	routerErr := router.Run(":3000")
	if routerErr != nil {
		return
	} // listen and serve on 0.0.0.0:3000 (for windows "localhost:3000")
}
