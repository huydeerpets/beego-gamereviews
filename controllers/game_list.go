package controllers

import (
	"github.com/astaxie/beego"
	"github.com/macococo/beego-gamereviews/models"
	"github.com/macococo/beego-gamereviews/modules"
)

type GameListController struct {
	beego.Controller
}

func (this *GameListController) Get() {
	key := "game_list"
	exist := modules.JsonRequestCache(this.Ctx, key)
	if !exist {
		gameManager := models.GameManager{}
		games := gameManager.All()

		modules.WriteInterfaceJson(this.Ctx, key, &games)
	}
}
