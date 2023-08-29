package main

import (
	"github.com/byteplow/nkajdavk/helper"
	"github.com/byteplow/nkajdavk/models"
	"github.com/byteplow/nkajdavk/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(helper.GetEnv("NKA_MODE", "debug")) // debug or release
	user := helper.GetEnv("NKA_USER", "user")
	password := helper.GetEnv("NKA_PASSWORD", "password") // not hashed :(

	r := routes.SetupRouter(user, password)

	dsn := helper.GetEnv("NKA_DSN", "dev.sqllite") // :memory: or file path
	models.SetupDatabase(dsn)

	addr := helper.GetEnv("NKA_ADDRESS", ":8080")
	r.Run(addr)
}
