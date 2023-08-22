package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/nick-leslie/scorekeeper/internal/api"
	"github.com/nick-leslie/scorekeeper/internal/static"
)

var buildDir = "frontend/build"

//go:generate npm i
//go:generate npm run build
//go:embed all:frontend/build
var files embed.FS

func main() {

	router := gin.Default()
	static.SetupStaticEmbed(router, &files, buildDir, buildDir)
	api.SetupApi(router)
	err := router.Run(":3000")
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:3000 (for windows "localhost:3000")
}
