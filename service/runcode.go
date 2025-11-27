package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/astaxie/beego/logs"

	"github.com/gongmh/runcode/models"

	"github.com/gongmh/runcode/helpers"

	"github.com/astaxie/beego"
	"github.com/codeskyblue/go-sh"
)

func GetCodeTemplate(codeID string) (codeTemplate map[string]CodeStruct, codeType string, codeStat string, err error) {
	//1. 配置默认模板代码
	codeType = "golang"
	codeTemplate = codeShowMap

	//2. 统计代码执行信息
	countInfo, err := models.GetCodeCount()
	if err != nil {
		return
	}
	countJson, _ := json.Marshal(countInfo)
	codeStat = string(countJson)

	if len(codeID) == 0 {
		return codeTemplate, codeType, codeStat, nil
	}

	//3. 根据用户cookie的codeID获取之前保存代码
	cd := &models.CodeDetail{
		CodeId: codeID,
	}
	err = models.GetCodeDetailByCodeID(cd)
	if err != nil {
		logs.Warning("GetCodeTemplate: GetCodeDetailByCodeID error. cd=%v, err=%v", cd, err)
		return
	}

	if len(cd.Cmd) == 0 || len(cd.CodeType) == 0 {
		logs.Warning("GetCodeTemplate: code detail empty. cd=%v", cd)
		return
	}

	codeType = cd.CodeType

	item := codeTemplate[codeType]
	item.DemoCode = cd.Cmd
	codeTemplate[codeType] = item

	return
}

func CodeRun(codeType, codeContent, clientIp, userAgent string) (result string, codeID string, err error) {
	md5Str := helpers.GenStrMd5sum(codeContent)

	cd := &models.CodeDetail{
		CodeType:  codeType,
		Cmd:       codeContent,
		CodeId:    codeID,
		Md5sum:    md5Str,
		ClientIp:  clientIp,
		UserAgent: userAgent,
	}

	err = models.GetCodeDetailByMd5(cd)
	if err != nil && err.Error() != models.RecordNotFoundErr {
		logs.Warning("CodeRun: GetCodeDetailByMd5 error. cd=%v, err=%v", cd, err)
		return
	}

	if err == nil && cd.RunResult == models.RUNCODE_RESULT_SUCC && cd.RunResultInfo != "" {
		logs.Info("CodeRun: code exist. cd=%v", cd.CodeId)
		return cd.RunResultInfo, cd.CodeId, nil
	}

	runResult, runErr := runCodeInDocker(codeType, md5Str, codeContent)
	if runErr != nil {
		logs.Warning("CodeRun: runCodeInDocker error. runResult=%v, err=%v", runResult, runErr)
		cd.RunResult = models.RUNCODE_RESULT_ERR
		result = runErr.Error()
	} else {
		cd.RunResult = models.RUNCODE_RESULT_SUCC
		result = runResult
	}

	cd.RunResultInfo = result

	//已存在，更新
	if cd.CodeId != "" {
		_, err = cd.Update(nil, "run_result", "run_result_info")
		if err != nil {
			logs.Warning("CodeRun: update code detail error. codeID=%v, err=%v", cd.CodeId, err)
			return "", "", err
		}
		return result, codeID, nil
	}

	cd.CodeId = genCodeID(cd)
	cd.InsertType = CODE_INSERT_TYPE_BY_RUN
	_, err = cd.Insert(nil)
	if err != nil {
		logs.Warning("CodeRun: Insert code detail error. codeID=%v, err=%v", cd.CodeId, err)
		return "", "", err
	}

	return result, codeID, err
}

func runCodeInDocker(codeType, md5, codeContent string) (result string, err error) {
	cs, ok := codeShowMap[codeType]
	if !ok {
		logs.Warning("runCodeInDocker: code type error. codeType=%v", codeType)
		return "", helpers.CodeTypeError
	}

	//1. 写入到本地文件
	filePath := beego.AppConfig.String("runcodepath")
	fileName := filePath + codeType + md5 + cs.postfix
	err = ioutil.WriteFile(fileName, []byte(codeContent), 0666)
	if err != nil {
		logs.Warning("runCodeInDocker: WriteFile error. err=%v", err)
		return "", err
	}

	//2. 执行代码
	session := sh.NewSession()
	cmd := fmt.Sprintf(cs.dockerCmd, fileName)
	s, err := session.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		logs.Warning("runCodeInDocker: Command error. cmd=%s, err=%v", cmd, err)
		//return "", err
	}
	return string(s), nil
}

func genCodeID(codeDetail *models.CodeDetail) string {
	return fmt.Sprintf("%s_%s", codeDetail.CodeType, codeDetail.Md5sum)
}
