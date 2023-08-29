package controllers

import (
	"net/http"

	"github.com/byteplow/nkajdavk/models"
	"github.com/gin-gonic/gin"
)

func GetList(c *gin.Context) {
	var contacts []models.Contact

	err := models.GetContacts(&contacts)
	if err != nil {
		contacts = []models.Contact{}
	}

	c.HTML(http.StatusOK, "list.tmpl", contacts)
}
