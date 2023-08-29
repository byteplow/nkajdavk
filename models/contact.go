package models

type Contact struct {
	Token                       string `uri:"token" gorm:"<-;primaryKey"`
	ForenameChild               string `form:"ForenameChild"`
	SurnameChild                string `form:"SurnameChild"`
	BirthdateChild              string `form:"BirthdateChild"`
	EmailChild                  string `form:"EmailChild"`
	EmailListChild              bool   `from:"EmailListChild"`
	PhonenumberChild            string `form:"PhonenumberChild"`
	AddressLineOneChild         string `form:"AddressLineOneChild"`
	AddressLineTowChild         string `form:"AddressLineTowChild"`
	AddressLineThreeChild       string `form:"AddressLineThreeChild"`
	MembershipNumberChild       string `from:"MembershipNumberChild"`
	ForenameGuardianOne         string `form:"ForenameGuardianOne"`
	SurnameGuardianone          string `form:"SurnameGuardianone"`
	EmailGuardianOne            string `form:"EmailGuardianOne"`
	EmailListGuardianOne        bool   `from:"EmailListGuardianOne"`
	PhonenumberOneGuardianOne   string `form:"PhonenumberOneGuardianOne"`
	PhonenumberTowGuardianOne   string `form:"PhonenumberTowGuardianOne"`
	AddressLineOneGuardianOne   string `form:"AddressLineOneGuardianOne"`
	AddressLineTowGuardianOne   string `form:"AddressLineTowGuardianOne"`
	AddressLineThreeGuardianOne string `form:"AddressLineThreeGuardianOne"`
	ForenameGuardianTow         string `form:"AddressLineThreeGuardianOne"`
	SurnameGuardianTow          string `form:"SurnameGuardianTow"`
	EmailGuardianTow            string `form:"EmailGuardianTow"`
	EmailListGuardianTow        bool   `from:"EmailListGuardianTow"`
	PhonenumberOneGuardianTow   string `form:"PhonenumberOneGuardianTow"`
	PhonenumberTowGuardianTow   string `form:"PhonenumberTowGuardianTow"`
	AddressLineOneGuardianTow   string `form:"AddressLineOneGuardianTow"`
	AddressLineTowGuardianTow   string `form:"AddressLineTowGuardianTow"`
	AddressLineThreeGuardianTow string `form:"AddressLineThreeGuardianTow"`
}

func CreateOrUpdateContact(contact *Contact) (err error) {
	if err = DB.Debug().Save(contact).Error; err != nil {
		return err
	}

	return nil
}

func GetContact(contact *Contact, token string) (err error) {
	if err = DB.Debug().First(contact, "token = ?", token).Error; err != nil {
		return err
	}

	return nil
}
