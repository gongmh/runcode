package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

const RecordNotFoundErr = "<QuerySeter> no row found"

//code run result
const (
	RUNCODE_RESULT_INIT = 0
	RUNCODE_RESULT_SUCC = 1
	RUNCODE_RESULT_ERR  = 2
)

//code detail orm
type CodeDetail struct {
	Id            int64
	CodeId        string
	Cmd           string
	CodeType      string
	InsertType    string
	RunResult     int
	RunResultInfo string
	Md5sum        string
	ClientIp      string
	UserAgent     string
}

func init() {
	orm.RegisterModelWithPrefix("user_", new(CodeDetail))
}

func (cd *CodeDetail) Insert(o orm.Ormer) (int64, error) {
	if o == nil {
		o = orm.NewOrm()
	}
	return o.Insert(cd)
}

func (cd *CodeDetail) Update(o orm.Ormer, cols ...string) (int64, error) {
	if o == nil {
		o = orm.NewOrm()
	}
	return o.Update(cd, cols...)
}

func GetCodeDetailByMd5(cd *CodeDetail) (err error) {
	return orm.NewOrm().Read(cd, "code_type", "md5sum")
}

func GetCodeDetailByCodeID(cd *CodeDetail) (err error) {
	return orm.NewOrm().Read(cd, "code_id")
}

type CountStruct struct {
	CodeType string
	Count    int64
}

func GetCodeCount() (cs []CountStruct, err error) {
	_, err = orm.NewOrm().Raw("SELECT code_type,count(*) AS \"count\" FROM user_code_detail group by code_type;").QueryRows(&cs)
	if err != nil {
		logs.Warning("[Warn] read db codetype error.")
		return nil, err
	}

	return cs, nil
}
