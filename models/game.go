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

func (this *GameManager) TableName() string {
	const TABLE_NAME = "t_game"
	return TABLE_NAME
}

func (this *GameManager) All() []*Game {
	var games []*Game
	_, err := this.QueryTable(this.TableName()).All(&games)
	if err != nil {
		beego.Error(err)
	}

	return games
}

func (this *GameManager) Count() int64 {
	cnt, err := this.QueryTable(this.TableName()).Count()
	if err != nil {
		beego.Error(err)
	}

	return cnt
}

func (this *GameManager) Find(pagination *Pagination) []*Game {
	var games []*Game
	_, err := this.QueryTable(this.TableName()).Limit(pagination.Length).Offset((pagination.Page - 1) * pagination.Length).All(&games)
	if err != nil {
		beego.Error(err)
	}

	return games
}

func (this *GameManager) Get(id int) *Game {
	game := Game{Id: id}

	err := this.GetOrm().Read(&game)
	if err != nil {
		beego.Error(err)
	}

	return &game
}

func (this *GameManager) Create(game *Game) *Game {
	_, err := this.GetOrm().Insert(game)
	if err != nil {
		beego.Error(err)
		return nil
	}

	return game
}

func (this *GameManager) Update(game *Game) *Game {
	_, err := this.GetOrm().Update(game)
	if err != nil {
		beego.Error(err)
		return nil
	}

	return game
}
