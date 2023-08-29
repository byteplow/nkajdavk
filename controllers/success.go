package controllers

import (
	"log"
	"net/http"

	"github.com/byteplow/nkajdavk/models"
	"github.com/gin-gonic/gin"
)

func GetSuccess(c *gin.Context) {
	var contact models.Contact

	err := models.GetContact(&contact, c.Param("token"))
	if err != nil {
		contact = models.Contact{
			Token: c.Param("token"),
		}
	}

	log.Println(contact)

	c.HTML(http.StatusOK, "success.tmpl", contact)
}
