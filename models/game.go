package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Game struct {
	Id                             int     `orm:"pk" json:"trackId"`
	Name                           string  `json:"trackName"`
	Version                        string  `orm:"null"`
	Price                          int     `orm:"null"`
	Url                            string  `orm:"null" json:"artworkUrl60"`
	Artwork1Url                    string  `orm:"null" json:"artworkUrl100"`
	Artwork2Url                    string  `orm:"null" json:"artworkUrl512"`
	Artwork3Url                    string  `orm:"null"`
	ArtistId                       int     `orm:"null"`
	ArtistName                     string  `orm:"null"`
	AverageRating                  float32 `orm:"null" json:"averageUserRating"`
	RatingCount                    int     `orm:"null" json:"userRatingCount"`
	AverageRatingForCurrentVersion float32 `orm:"null" json:"averageUserRatingForCurrentVersion"`
	RatingCountForCurrentVersion   int     `orm:"null" json:"userRatingCountForCurrentVersion"`
	JsonText                       string  `orm:"null"`
	ReviewCount                    int     `orm:"default(0)"`
}

type GameManager struct {
	ModelManager
}

func init() {
	orm.RegisterModelWithPrefix("t_", new(Game))
}

func (this *GameManager) All() []*Game {
	o := this.GetOrm()

	var games []*Game
	_, err := o.QueryTable("t_game").All(&games)
	if err != nil {
		beego.Error(err)
	}

	return games
}

func (this *GameManager) Count() int64 {
	o := this.GetOrm()

	cnt, err := o.QueryTable("t_game").Count()
	if err != nil {
		beego.Error(err)
	}

	return cnt
}

func (this *GameManager) Find(pagination *Pagination) []*Game {
	o := this.GetOrm()

	var games []*Game
	_, err := o.QueryTable("t_game").Limit(pagination.Length).Offset((pagination.Page - 1) * pagination.Length).All(&games)
	if err != nil {
		beego.Error(err)
	}

	return games
}

func (this *GameManager) Get(id int) *Game {
	o := this.GetOrm()

	game := Game{Id: id}

	err := o.Read(&game)
	if err != nil {
		beego.Error(err)
		return nil
	}

	return &game
}

func (this *GameManager) Create(game *Game) *Game {
	o := this.GetOrm()

	_, err := o.Insert(game)
	if err != nil {
		beego.Error(err)
		return nil
	}

	return game
}

func (this *GameManager) Update(game *Game) *Game {
	o := this.GetOrm()

	_, err := o.Update(game)
	if err != nil {
		beego.Error(err)
		return nil
	}

	return game
}
