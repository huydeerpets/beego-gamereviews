package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type GameReview struct {
	Id         int
	GameId     int
	ReviewText string
}

type GameReviewManager struct {
	ModelManager
}

func init() {
	orm.RegisterModelWithPrefix("t_", new(GameReview))
}

func (this *GameReviewManager) TableName() string {
	const TABLE_NAME = "t_game_review"
	return TABLE_NAME
}

func (this *GameReviewManager) All() []*GameReview {
	var gameReviews []*GameReview
	_, err := this.QueryTable(this.TableName()).All(&gameReviews)
	if err != nil {
		beego.Error(err)
	}

	return gameReviews
}

func (this *GameReviewManager) Count() int64 {
	cnt, err := this.QueryTable(this.TableName()).Count()
	if err != nil {
		beego.Error(err)
	}

	return cnt
}

func (this *GameReviewManager) Find(pagination *Pagination) []*GameReview {
	var gameReviews []*GameReview
	_, err := this.QueryTable(this.TableName()).Limit(pagination.Length).Offset((pagination.Page - 1) * pagination.Length).All(&gameReviews)
	if err != nil {
		beego.Error(err)
	}

	return gameReviews
}

func (this *GameReviewManager) Get(id int) *GameReview {
	gameReview := GameReview{Id: id}

	err := this.GetOrm().Read(&gameReview)
	if err != nil {
		beego.Error(err)
	}

	return &gameReview
}

func (this *GameReviewManager) Create(gameReview *GameReview) *GameReview {
	_, err := this.GetOrm().Insert(gameReview)
	if err != nil {
		beego.Error(err)
		return nil
	}

	return gameReview
}

func (this *GameReviewManager) Update(gameReview *GameReview) *GameReview {
	_, err := this.GetOrm().Update(gameReview)
	if err != nil {
		beego.Error(err)
		return nil
	}

	return gameReview
}
