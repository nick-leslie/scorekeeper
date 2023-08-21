package static

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/nick-leslie/scorekeeper/internal/helper"
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
			router.GET(urlPath+name, HandleFuncStatic(readFile, name))
		} else {
			SetupStaticEmbed(router, fileSystem, path+"/"+name, basePath)
		}
	}
	if direrr != nil {
		println(direrr.Error())
		return
	}
}

func HandleFuncStatic(data []byte, name string) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		helper.SetHeaders(context, name)
		context.Data(200, name, data)
	}
	return fn
}
