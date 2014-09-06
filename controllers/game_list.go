package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/macococo/beego-gamereviews/models"
	"github.com/macococo/beego-gamereviews/modules"
)

type GameListController struct {
	beego.Controller
}

func (this *GameListController) Get() {
	exist := modules.JsonRequestCache(this.Ctx, "game_list")
	if !exist {
		gameManager := models.GameManager{}
		games := gameManager.All()

		content, err := json.Marshal(&games)
		if err != nil {
			beego.Error(err)
		}

		modules.AppCache.Put("game_list", string(content), 3600)

		this.Ctx.Output.Body(content)
	}
}
