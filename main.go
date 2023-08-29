package main

import (
	"log"

	"github.com/byteplow/nkajdavk/helper"
	"github.com/byteplow/nkajdavk/models"
	"github.com/byteplow/nkajdavk/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println(1)

	gin.SetMode(helper.GetEnv("NKA_MODE", "debug")) // debug or release
	r := routes.SetupRouter()

	dsn := helper.GetEnv("NKA_DSN", ":memory:") // :memory: or file path
	models.SetupDatabase(dsn)

	addr := helper.GetEnv("NKA_ADDRESS", ":8080")
	r.Run(addr)
}
