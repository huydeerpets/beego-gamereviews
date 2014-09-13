package routers

import (
	"github.com/astaxie/beego"
	"github.com/macococo/beego-gamereviews/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/user/list", &controllers.UserListController{})
	beego.Router("/api/user/create", &controllers.UserCreateController{})
	beego.Router("/api/game/list", &controllers.GameListController{})
	beego.Router("/api/game/review", &controllers.GameReviewController{})
}
