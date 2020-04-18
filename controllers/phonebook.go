package controllers

import (
	"phonebook/models"
	"encoding/json"
	
	"github.com/astaxie/beego"
)

// Operations about phonebook
type PhonebookController struct {
	beego.Controller
}

// @Title Create
// @Description create phonebook
// @Param	body		body 	models.Phonebook	true		"The phonebook content"
// @Success 200 {string} models.Phonebook.Id
// @Failure 403 body is empty
// @router / [post]
func (o *PhonebookController) Post() {

	var ob models.Phonebook
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	phonebookId := models.AddPhonebook(ob)
	o.Data["json"] = map[string]string{"PhonebookId": phonebookId}
	o.ServeJSON()
}

// @Title Get
// @Description find object by phonebookId
// @Param	phonebookId		path 	string	true		"the phonebookId you want to get"
// @Success 200 {phonebook} models.Phonebook
// @Failure 403 :phonebookId is empty
// @router /:phonebookId [get]
func (o *PhonebookController) Get() {
	phonebookId := o.Ctx.Input.Param(":phonebookId")
	if phonebookId != "" {
		ob, err := models.GetPhonebook(phonebookId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all phonebooks
// @Success 200 {phonebook} models.Phonebook
// @Failure 403 :phonebookId is empty
// @router / [get]
func (o *PhonebookController) GetAll() {
	obs := models.GetAllPhonebook()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title Update
// @Description update the phonebook
// @Param	phonebookId		path 	string	true		"The phonebookid you want to update"
// @Param	body		body 	models.Phonebook	true		"The body"
// @Success 200 {object} models.Phonebook
// @Failure 403 :phonebookId is empty
// @router /:phonebookId [put]
func (o *PhonebookController) Put() {
	phonebookId := o.Ctx.Input.Param(":phonebookId")
	var ob models.Phonebook
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	err := models.UpdatePhonebook(phonebookId, ob.Name)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the phonebook
// @Param	phonebookId		path 	string	true		"The phonebookId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 phonebookId is empty
// @router /:phonebookId [delete]
func (o *PhonebookController) Delete() {
	phonebookId := o.Ctx.Input.Param(":phonebookId")
	models.DeletePhonebook(phonebookId)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}

