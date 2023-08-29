package routers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Contact struct {
	Token                       string `uri:"token"`
	ForenameChild               string `form:"forename-child"`
	SurnameChild                string `form:"surname-child"`
	BirthdateChild              string `form:"birthdate-child"`
	EmailChild                  string `form:"email-child"`
	PhonenumberChild            string `form:"phonenumber-child"`
	AddressLineOneChild         string `form:"address-line-one-child"`
	AddressLineTowChild         string `form:"address-line-tow-child"`
	AddressLineThreeChild       string `form:"address-line-three-child"`
	ForenameguardianOne         string `form:"forename-guardian-one"`
	Surnameguardianone          string `form:"surname-guardian-one"`
	EmailguardianOne            string `form:"email-guardian-one"`
	PhonenumberOneguardianOne   string `form:"phonenumber-one-guardian-one"`
	PhonenumberTowguardianOne   string `form:"phonenumber-tow-guardian-one"`
	AddressLineOneguardianOne   string `form:"address-line-one-guardian-one"`
	AddressLineTowguardianOne   string `form:"address-line-tow-guardian-one"`
	AddressLineThreeguardianOne string `form:"address-line-three-guardian-one"`
	ForenameguardianTow         string `form:"forename-guardian-tow"`
	SurnameguardianTow          string `form:"surname-guardian-tow"`
	EmailguardianTow            string `form:"email-guardian-tow"`
	PhonenumberOneguardianTow   string `form:"phonenumber-one-guardian-tow"`
	PhonenumberTowguardianTow   string `form:"phonenumber-tow-guardian-tow"`
	AddressLineOneguardianTow   string `form:"address-line-one-guardian-tow"`
	AddressLneTowguardianTow    string `form:"address-line-tow-guardian-tow"`
	AddressLineThreeguardianTow string `form:"address-line-three-guardian-tow"`
}

func GetEdit(c *gin.Context) {
	c.HTML(http.StatusOK, "edit.tmpl", gin.H{})
}

func PostEdit(c *gin.Context) {
	var contact Contact

	c.Bind(&contact)
	c.BindUri(&contact)

	token := c.Params.ByName("token")
	if token == "" {
		panic("not implementet ")
	}

	log.Println(contact)

	c.Redirect(http.StatusFound, "/edit/"+token)
}
