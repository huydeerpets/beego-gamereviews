package controllers

import (
	"github.com/astaxie/beego"
	"github.com/macococo/beego-gamereviews/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	page, err := this.GetInt("page")
	if page <= 0 || err != nil {
		page = 1
	}

	gameManager := models.GameManager{}
	pagination := models.NewPagination(int(page), 12, gameManager.Count())
	games := gameManager.Find(pagination)

	this.Data["Games"] = &games
	this.Data["Pagination"] = &pagination
	this.TplNames = "index.tpl"
}
