package controllers

import (
	"github.com/astaxie/beego"
	"github.com/macococo/beego-gamereviews/models"
	"github.com/macococo/beego-gamereviews/modules"
	"strconv"
)

type UserListController struct {
	beego.Controller
}

type UserListResponse struct {
	Pagination *models.Pagination
	Users      []*models.User
}

func (this *UserListController) Get() {
	page, err := this.GetInt("page")
	if page <= 0 || err != nil {
		page = 1
	}
	t, err := this.GetInt("type")
	if t < 0 || t > 2 || err != nil {
		t = 1
	}

	key := "user_list_" + strconv.FormatInt(t, 10) + "_" + strconv.FormatInt(page, 10)
	exist := modules.JsonRequestCache(this.Ctx, key)
	if !exist {
		userManager := models.UserManager{}
		pagination := models.NewPagination(int(page), 10, userManager.Count(t))
		users := userManager.Find(t, pagination)

		response := UserListResponse{Pagination: pagination, Users: users}
		modules.WriteInterfaceJson(this.Ctx, key, &response)
	}
}
