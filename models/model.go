package models

import (
	"github.com/astaxie/beego/orm"
)

type ModelManager struct {
	o orm.Ormer
}

func (this *ModelManager) GetOrm() orm.Ormer {
	if this.o == nil {
		this.o = orm.NewOrm()
	}
	return this.o
}

func (this *ModelManager) QueryTable(tableName string) orm.QuerySeter {
	return this.GetOrm().QueryTable(tableName)
}
