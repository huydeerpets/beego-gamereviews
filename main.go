package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/macococo/beego-gamereviews/routers"
	"github.com/macococo/beego-gamereviews/tasks"
)

func init() {
	initDb()
	initTask()
}

func initDb() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "gamereviews:gamereviews@tcp(127.0.0.1:3306)/gamereviews?charset=utf8")
	orm.SetMaxIdleConns("default", 100)
	orm.SetMaxOpenConns("default", 100)

	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		panic(err)
	}
}

func initTask() {
	task := toolbox.NewTask("SearchITunes", "0/60 * * * * *", tasks.SearchITunes)
	toolbox.AddTask("SearchITunes", task)
	toolbox.StartTask()
}

func main() {
	beego.Run()
}
