package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/macococo/beego-gamereviews/models"
	"math/rand"
)

type UserCreateController struct {
	beego.Controller
}

func (this *UserCreateController) Get() {
	userManager := models.UserManager{}
	user := models.User{Type: rand.Intn(3)}
	userManager.Create(&user)

	content, err := json.Marshal(&user)
	if err != nil {
		beego.Error(err)
	}
	this.Ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
	this.Ctx.Output.Body(content)
}
