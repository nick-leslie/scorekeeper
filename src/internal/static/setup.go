package static

import (
	"embed"
	"github.com/gin-gonic/gin"
	"strings"
)

func SetupStaticEmbed(router *gin.Engine, fileSystem *embed.FS, path string, basePath string) {
	dir, direrr := fileSystem.ReadDir(path)
	urlPath := strings.ReplaceAll(path, basePath, "") + "/"
	for i := 0; i < len(dir); i++ {
		file := dir[i]
		name := file.Name()
		if file.IsDir() == false {
			readFile, readFileErr := fileSystem.ReadFile(path + "/" + file.Name())
			if readFileErr != nil {
				println(readFileErr.Error())
				continue
			}
			router.GET(urlPath+name, func(context *gin.Context) {
				data := readFile
				parts := strings.Split(name, ".")
				ext := parts[len(parts)-1]
				switch ext {
				case "html":
					context.Header("Content-Type", "text/html")
				case "css":
					context.Header("Content-Type", "text/css")
				case "js":
					context.Header("Content-Type", "application/javascript")
				}
				context.Data(200, name, data)
			})
		} else {
			SetupStaticEmbed(router, fileSystem, path+"/"+name, basePath)
		}
	}
	if direrr != nil {
		println(direrr.Error())
		return
	}
}
