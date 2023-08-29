package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.LoadHTMLGlob("./templates/*")

	r.StaticFS("/static", http.Dir("./static"))

	r.GET("/", GetIndex)
	r.GET("/edit/:token", GetEdit)
	r.POST("/edit/:token", PostEdit)

	return r
}
