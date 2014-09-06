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

func (this *GameReviewManager) All() []*GameReview {
	o := this.GetOrm()

	var games []*GameReview
	_, err := o.QueryTable("t_game_review").All(&games)
	if err != nil {
		beego.Error(err)
	}

	return games
}

func (this *GameReviewManager) Count() int64 {
	o := this.GetOrm()

	cnt, err := o.QueryTable("t_game_review").Count()
	if err != nil {
		beego.Error(err)
	}

	return cnt
}

func (this *GameReviewManager) Find(pagination *Pagination) []*GameReview {
	o := this.GetOrm()

	var games []*GameReview
	_, err := o.QueryTable("t_game_review").Limit(pagination.Length).Offset((pagination.Page - 1) * pagination.Length).All(&games)
	if err != nil {
		beego.Error(err)
	}

	return games
}

func (this *GameReviewManager) Create(game *GameReview) *GameReview {
	o := this.GetOrm()

	_, err := o.Insert(game)
	if err != nil {
		beego.Error(err)
		return nil
	}

	return game
}

func (this *GameReviewManager) Update(game *GameReview) *GameReview {
	o := this.GetOrm()

	_, err := o.Update(game)
	if err != nil {
		beego.Error(err)
		return nil
	}

	return game
}
