package tasks

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/macococo/beego-gamereviews/models"
	"net/http"
)

type ITunesJson struct {
	ResultCount int
	Results     []*models.Game
}

func SearchITunes() error {
	beego.Info("execute search iTunes API.")

	resp, err := http.Get("https://itunes.apple.com/search?term=game&country=jp&lang=ja_jp&limit=200&entity=software")
	defer resp.Body.Close()
	if err != nil {
		beego.Error(err)
	}

	result := ITunesJson{}
	json.NewDecoder(resp.Body).Decode(&result)

	gameManager := models.GameManager{}
	games := gameManager.All()
	gameMap := map[int]*models.Game{}
	for _, game := range games {
		gameMap[game.Id] = game
	}

	for _, game := range result.Results {
		exist := gameMap[game.Id]
		if exist == nil {
			gameManager.Create(game)
		} else {
			game.ReviewCount = exist.ReviewCount
			gameManager.Update(game)
		}
	}

	return nil
}
