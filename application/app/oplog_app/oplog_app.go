package oplog_app

import (
	"asm_platform/application/dto"
	"asm_platform/application/vo"
	"asm_platform/infrastructure/pkg/constants"
	constapicode "asm_platform/infrastructure/pkg/constants/api_code"
	"asm_platform/infrastructure/pkg/slog"
	utils_tool "asm_platform/infrastructure/pkg/tool/utils"
	"asm_platform/infrastructure/repo"
	"encoding/json"
	"time"
)

type OpLogApp struct {
}

func NewOpLogApp() OpLogAppInterface {
	return &OpLogApp{}
}

// oplogApp implements OpLogAppInterface
var _ OpLogAppInterface = &OpLogApp{}

// SaveOpLog
// @Description:     保存文档信息
// @params opDto     日志DTO
// @return SocError  错误码
//
func (o OpLogApp) SaveOpLog(opDto *dto.OpLogDTO) constapicode.SocError {
	// 生成编号
	opDto.ID = utils_tool.GenerateUniqueId()

	// 处理response结果
	var res2 dto.ResponseKV
	if opDto.ResponseResult != "" && opDto.RequestUri != constants.CaptChaApiPath {
		jsonErr := json.Unmarshal([]byte(opDto.ResponseResult), &res2)
		if jsonErr != nil {
			slog.Errorf("[asm op log create]:format response  json fail,err: %v", jsonErr)
			// 解析失败，不是一个完整的json
			res2 = dto.ResponseKV{
				Code:    0,
				Message: "",
			}
		}
		opDto.ResponseKV = res2
	}
	// 处理login request args结果，登录密码格式需要隐藏；
	if !constants.FilterLoginApiPath.Contains(opDto.RequestUri) {
		if opDto.RequestArgs != "" {
			var login dto.LoginDTO
			jsonErr := json.Unmarshal([]byte(opDto.RequestArgs), &login)
			if jsonErr != nil {
				slog.Errorf("[asm op log create]:format request args json fail,err: %v", jsonErr)
			} else {
				login.Password = "**********"
				a, _ := json.Marshal(login)
				opDto.RequestArgs = string(a)
			}
		}
	}

	// 组装日志 dto to entity
	opLogInfo := opDto.OpLogDtoTOEntity()

	// 保存操作日志
	var indexName = getIndexName(constants.AsmOpLogIndex, utils_tool.FormatYearMonth(time.Now()))
	opLogRepo := repo.NewOpLogRepository(indexName)
	err := opLogRepo.SaveOpLog(opLogInfo)
	if err != nil {
		slog.Errorf("[soc op log create]:format json fail,err: %v", indexName)
		return constapicode.DocumentSaveFail
	}
	return constapicode.Success
}

// SearchOpLogList
// @Description: 查询文档信息
// @params opLogQuery 日志查询DTO
// @return results    操作日志实体
// @return code       错误码
//
func (o OpLogApp) SearchOpLogList(opLogQuery *dto.OpLogQueryDTO) (results []vo.OpLogVo, totalCount int64, code constapicode.SocError) {
	// op query
	opQuery := opLogQuery.OpLogQueryDtoTOEntity()
	var indexName = getIndexName(constants.AsmOpLogIndex, utils_tool.FormatYearMonth(time.Now()))
	opLogRepo := repo.NewOpLogRepository(indexName)
	opLogList, totalCount, err := opLogRepo.FindOpLogList(opQuery)
	if err != nil {
		return nil, 0, constapicode.ESSearchError
	}
	if len(opLogList) > 0 {
		for _, item := range opLogList {
			opLogVo := item.OpLogToVo()
			results = append(results, opLogVo)
		}
	}
	return results, totalCount, constapicode.Success
}

// EditOpLog
// @Description:     修改文档信息
// @params opDto     日志DTO
// @return SocError  错误码
//
func (o OpLogApp) EditOpLog(oplog *dto.OpLogDTO) constapicode.SocError {
	//TODO implement me
	panic("implement me")
}

// DeleteOpLogById
// @Description:     根据ID删除操作日志
// @params id        日志文档ID
// @return SocError  错误码
//
func (o OpLogApp) DeleteOpLogById(id string) constapicode.SocError {
	var indexName = getIndexName(constants.AsmOpLogIndex, utils_tool.FormatYearMonth(time.Now()))
	opLogRepo := repo.NewOpLogRepository(indexName)
	err := opLogRepo.DeleteOpLogById(id)
	if err != nil {
		return constapicode.DocumentDeleteFail
	}
	return constapicode.Success
}

// DeleteOpLogByQuery
// @Description:      根据查询条件删除操作日志
// @params opLogQuery 查询条件
// @return SocError   错误码
//
func (o OpLogApp) DeleteOpLogByQuery(opLogQuery *dto.OpLogQueryDTO) constapicode.SocError {
	//TODO implement me
	panic("implement me")
}
