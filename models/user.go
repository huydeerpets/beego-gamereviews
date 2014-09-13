package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id   int
	Name string
	Type int
}

func (this *User) TableName() string {
	return "t_user"
}

// func (this *User) TableEngine() string {
// 	return "MyISAM"
// }

type UserManager struct {
	ModelManager
}

func init() {
	orm.RegisterModel(new(User))
}

func (this *UserManager) All() []*User {
	var users []*User
	_, err := this.GetOrm().QueryTable((*User)(nil)).All(&users)
	if err != nil {
		beego.Error(err)
	}

	return users
}

func (this *UserManager) Count(t int64) int64 {
	cnt, err := this.GetOrm().QueryTable((*User)(nil)).
		Filter("type", t).
		Count()
	if err != nil {
		beego.Error(err)
	}

	return cnt
}

func (this *UserManager) Find(t int64, pagination *Pagination) []*User {
	var users []*User
	_, err := this.GetOrm().QueryTable((*User)(nil)).
		Filter("type", t).
		Limit(pagination.Length).
		Offset((pagination.Page - 1) * pagination.Length).
		All(&users)
	if err != nil {
		beego.Error(err)
	}

	return users
}

func (this *UserManager) Get(id int) *User {
	user := User{Id: id}

	err := this.GetOrm().Read(&user)
	if err != nil {
		beego.Error(err)
	}

	return &user
}

func (this *UserManager) Create(user *User) *User {
	_, err := this.GetOrm().Insert(user)
	if err != nil {
		beego.Error(err)
		return nil
	}

	return user
}

func (this *UserManager) Update(user *User) *User {
	_, err := this.GetOrm().Update(user)
	if err != nil {
		beego.Error(err)
		return nil
	}

	return user
}
