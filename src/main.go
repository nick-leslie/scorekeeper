package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
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
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"test": 1,
		})
	})
	err := router.Run(":3000")
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	fmt.Println("Hello, world.")
}
