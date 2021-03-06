package models

import (
	"errors"
	"strconv"
	"time"
	"gopkg.in/go-playground/validator.v9"
	"strings"
	"reflect"
)

var (
	Phonebooks map[string]*Phonebook
)

type Phonebook struct {
	PhonebookId   string
	Name      	string `json:"name" validate:"required"`
	PhoneNumber string `json:"phonenumber" validate:"required"`
	Address string `json:"address" validate:"required"`
}
var validate *validator.Validate

func init() {
	Phonebooks = make(map[string]*Phonebook)
	Phonebooks["hjkhsbnmn123"] = &Phonebook{"hjkhsbnmn123", "Jeff", "021-8909", "Jakarta Indonesia"}
	Phonebooks["hjkhsbnmn124"] = &Phonebook{"hjkhsbnmn124", "Bezoz", "021-1309", "Bandung Indonesia"}
}

//Model for add phonebook
func AddPhonebook(phonebook Phonebook) (PhonebookId string) {

	//Validation for new phonebook request, with rules define in struct
	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	err := validate.Struct(phonebook)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return err.Field() + " " + err.Tag()
		}
	}

	phonebook.PhonebookId = "astaxie" + strconv.FormatInt(time.Now().UnixNano(), 10)
	Phonebooks[phonebook.PhonebookId] = &phonebook
	return phonebook.PhonebookId
}

//Model for get single phonebook with specific PhonebookID
func GetPhonebook(PhonebookId string) (phonebook *Phonebook, err error) {
	if v, ok := Phonebooks[PhonebookId]; ok {
		return v, nil
	}
	return nil, errors.New("PhonebookId Not Exist")
}

//Model for get all phonebook
func GetAllPhonebook() map[string]*Phonebook {
	return Phonebooks
}

//Model for update phonebook name with specific PhonebookId
func UpdatePhonebook(PhonebookId string, Name string) (err error) {
	if v, ok := Phonebooks[PhonebookId]; ok {
		v.Name = Name
		return nil
	}
	return errors.New("PhonebookId Not Exist")
}

//Model for delete phonebook with specific PhonebookId
func DeletePhonebook(PhonebookId string) {
	delete(Phonebooks, PhonebookId)
}

