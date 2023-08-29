package controllers

import (
	"encoding/hex"
	"log"
	"math/rand"
	"net/http"

	"github.com/byteplow/nkajdavk/models"
	"github.com/gin-gonic/gin"
)

func NewForm(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "/form/"+generateToken())
}

func generateToken() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func GetForm(c *gin.Context) {
	var contact models.Contact

	err := models.GetContact(&contact, c.Param("token"))
	if err != nil {
		contact = models.Contact{
			Token: c.Param("token"),
		}
	}

	log.Println(contact)

	c.HTML(http.StatusOK, "form.tmpl", contact)
}

func SubmitForm(c *gin.Context) {
	var contact models.Contact

	c.Bind(&contact)
	c.BindUri(&contact)

	log.Println(contact)
	err := models.CreateOrUpdateContact(&contact)
	if err != nil {
		log.Println(err)
	}

	c.Redirect(http.StatusFound, "/success/"+contact.Token)
}
