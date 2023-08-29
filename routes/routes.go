package routes

import (
	"net/http"

	"github.com/byteplow/nkajdavk/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("./templates/*")

	r.StaticFS("/static", http.Dir("./static"))

	form := r.Group("/form")
	{
		form.GET("/", controllers.NewForm)
		form.GET("/:token", controllers.GetForm)
		form.POST("/:token", controllers.SubmitForm)
	}

	r.GET("/success/:token", controllers.GetSuccess)

	r.GET("/", controllers.GetIndex)

	return r
}
