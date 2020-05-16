package models

import "github.com/astaxie/beego/orm"

type CodeShareDetail struct {
	Id        int64
	ShareId   string
	DetailId  string
	ViewCount int64
}

func init() {
	orm.RegisterModelWithPrefix("user_", new(CodeShareDetail))
}

func (csd *CodeShareDetail) Insert(o orm.Ormer) (int64, error) {
	if o == nil {
		o = orm.NewOrm()
	}
	return o.Insert(csd)
}

func (csd *CodeShareDetail) Update(o orm.Ormer, cols ...string) (int64, error) {
	if o == nil {
		o = orm.NewOrm()
	}
	return o.Update(csd, cols...)
}

func GetShareDetailByShareID(csd *CodeShareDetail) (err error) {
	return orm.NewOrm().Read(csd, "share_id")
}

func GetShareDetailByDetailID(csd *CodeShareDetail) (err error) {
	return orm.NewOrm().Read(csd, "detail_id")
}
