package helper

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func SetHeaders(context *gin.Context, name string) {
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
}
