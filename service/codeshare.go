package service

import (
	"github.com/astaxie/beego/logs"

	"github.com/gongmh/runcode/models"

	"github.com/gongmh/runcode/helpers"

	"github.com/astaxie/beego/orm"
	"github.com/teris-io/shortid"
)

func GenCodeShare(codeType, codeContent string) (shareId string, err error) {
	md5Str := helpers.GenStrMd5sum(codeContent)

	cd := &models.CodeDetail{
		CodeType: codeType,
		Md5sum:   md5Str,
	}
	err = models.GetCodeDetailByMd5(cd)
	if err != nil && err.Error() != models.RecordNotFoundErr {
		logs.Warn("GenCodeShare: GetCodeDetailByMd5 error. md5=%s, err=%v", md5Str, err)
		return "", err
	}

	if err == nil {
		csd := &models.CodeShareDetail{
			DetailId: cd.CodeId,
		}
		err = models.GetShareDetailByDetailID(csd)
		if err != nil {
			logs.Info("GenCodeShare: GetShareDetailByDetailID error. DetailId=%s, err=%v", cd.CodeId, err)
		}
		if csd.ShareId != "" {
			return csd.ShareId, nil
		}
	}

	o := orm.NewOrm()
	//数据库不存在，插入
	if cd.CodeId == "" {
		cd.CodeId = genCodeID(cd)
		cd.Cmd = codeContent
		cd.InsertType = CODE_INSERT_TYPE_BY_SHARE
		_, err = cd.Insert(o)
		if err != nil {
			logs.Info("GenCodeShare: Insert code detail error. Detail=%v, err=%v", cd, err)
			return "", err
		}

	}
	shareId = genShareID()
	sd := &models.CodeShareDetail{
		ShareId:  shareId,
		DetailId: cd.CodeId,
	}
	_, err = sd.Insert(o)
	if err != nil {
		logs.Info("GenCodeShare: Insert share detail error. Detail=%s, err=%v", sd, err)
		return "", err
	}

	return shareId, nil
}

func GetCodeShareDetail(shareID string) (codeTemplate map[string]CodeStruct, codeType string, codeStat string, err error) {
	if len(shareID) == 0 {
		logs.Warning("GetCodeShareDetail: shareID empty. shareID=%s", shareID)
		return nil, "", "", helpers.ShareIDError
	}
	sd := &models.CodeShareDetail{
		ShareId: shareID,
	}

	err = models.GetShareDetailByShareID(sd)
	if err != nil {
		logs.Warning("GetCodeShareDetail: GetShareDetailByShareID error. shareID=%s, err=%v", shareID, err)
		return nil, "", "", helpers.ShareIDError
	}

	shareUpdate := &models.CodeShareDetail{
		Id:        sd.Id,
		ViewCount: sd.ViewCount + 1,
	}
	_, err = shareUpdate.Update(nil, "view_count")
	if err != nil {
		logs.Warning("GetCodeShareDetail: Update share db error. errmsg=[%s], detail=[%#v]", err.Error(), shareUpdate)
	}

	return GetCodeTemplate(sd.DetailId)

}

func genShareID() string {
	shareID, err := shortid.Generate()
	if err != nil {
		return helpers.GetUuidString()
	}

	return shareID
}
