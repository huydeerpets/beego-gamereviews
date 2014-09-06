package controllers

import (
	"github.com/astaxie/beego"
	"github.com/macococo/beego-gamereviews/models"
)

type GameReviewController struct {
	beego.Controller
}

func (this *GameReviewController) Post() {
	text := this.GetString("text")
	gameId, err := this.GetInt("gameId")
	if err != nil {
		beego.Error(err)
	}

	gameManager := models.GameManager{}
	gameReviewManager := models.GameReviewManager{}

	review := models.GameReview{GameId: int(gameId), ReviewText: text}
	gameReviewManager.Create(&review)

	game := gameManager.Get(review.GameId)
	game.ReviewCount += 1
	gameManager.Update(game)

	this.Data["json"] = &review
	this.ServeJson()
}
