package models

import (
	"fmt"
	"testing"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func Init() {
	err := orm.RegisterDataBase("default", "sqlite3", "../databases/data/tools.db")
	fmt.Println(err)
}

func TestGetCodeDetailByMd5(t *testing.T) {
	Init()
	cd := &CodeDetail{
		CodeType: "golang",
		Md5sum:   "43d49a38cd2933ce5eeee5fb741e2e471",
	}
	err := GetCodeDetailByMd5(cd)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v\n", cd)
}
