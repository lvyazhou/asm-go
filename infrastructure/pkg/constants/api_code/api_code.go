package constapicode

// SocError 错误定义
type SocError int32

const (
	Success          SocError = 200
	NotAuth          SocError = 400
	ErrorOnLogin     SocError = 401
	ErrorVerifyToken SocError = 402

	DBTxBeginError      SocError = 501
	DBTxCommitError     SocError = 502
	DBTxRollbackError   SocError = 503
	DBTxSavePointError  SocError = 504
	DBTxRollbackToError SocError = 505

	ErrorNoData        SocError = 10000
	ErrorReq           SocError = 90000
	ErrorOnDataFind    SocError = 10002
	ErrorOnDataSave    SocError = 10004
	ErrorOnDataDelete  SocError = 10006
	ErrorOnShallowCopy SocError = 10008
	ErrorGenerateToken SocError = 10010
	ErrorParserToken   SocError = 10012
	LogOutError        SocError = 10013
	PleaseRepeatLogin  SocError = 10014
	DataConvertError   SocError = 10015
	LookupError        SocError = 10016
	DataRepeat         SocError = 10017

	UserExist        SocError = 10100
	UserNotExist     SocError = 10101
	OldPWDIncorrect  SocError = 10102
	ChangePWDError   SocError = 10103
	NoChangeOtherPWD SocError = 10104

	RoleExist    SocError = 10200
	RoleNotExist SocError = 10201
	RoleDisable  SocError = 10202
	RoleHasDel   SocError = 10203

	UserRoleExist    SocError = 10300
	UserRoleNotExist SocError = 10302
	MenuExist        SocError = 10400
	MenuNotExist     SocError = 10401

	// es common
	ESSearchError      SocError = 10801
	ESCreateIndexError SocError = 10802
	ESIndexExist       SocError = 10803
	DocumentSaveFail   SocError = 10804
	DocumentDelFail    SocError = 10805
	DocumentNotFind    SocError = 10806
	DocumentUpdateFail SocError = 10807
	DocumentDeleteFail SocError = 10808

	// common
	DateConvertError SocError = 10900

	// 操作日志
	OpLogCreateError SocError = 11500
)

func (e SocError) String() string {
	switch e {
	case Success:
		return "OK"
	case NotAuth:
		return "认证失败"
	case ErrorReq:
		return "请求参数错误"
	case ErrorOnLogin:
		return "用户名或密码错误"
	case ErrorVerifyToken:
		return "验证码错误"
	case ErrorNoData:
		return "无数据"
	case ErrorOnDataFind:
		return "数据查询错误"
	case ErrorOnDataSave:
		return "保存数据错误"
	case ErrorOnDataDelete:
		return "删除数据错误"
	case ErrorOnShallowCopy:
		return "数据处理错误"
	case ErrorGenerateToken:
		return "生成JWT Token失败"
	case UserExist:
		return "用户已经存在"
	case UserNotExist:
		return "用户不存在"
	case OldPWDIncorrect:
		return "当前密码不正确"
	case ChangePWDError:
		return "修改密码失败"
	case NoChangeOtherPWD:
		return "不能修改非本人密码"
	case RoleExist:
		return "角色已经存在"
	case RoleNotExist:
		return "角色不存在"
	case RoleDisable:
		return "角色已经禁用"
	case RoleHasDel:
		return "角色被删除"
	case UserRoleExist:
		return "用户已拥有该角色"
	case UserRoleNotExist:
		return "用户角色不存在"
	case MenuExist:
		return "菜单已经存在"
	case MenuNotExist:
		return "菜单不存在"
	case ErrorParserToken:
		return "Token失效"
	case LogOutError:
		return "退出失败"
	case PleaseRepeatLogin:
		return "请重新登录"
	case DataConvertError:
		return "数据转换失败"
	case LookupError:
		return "查看失败"
	case DataRepeat:
		return "数据重复,保存失败"
	case ESSearchError:
		return "查询失败"
	case ESCreateIndexError:
		return "创建索引失败"
	case ESIndexExist:
		return "索引已经存在,创建失败"
	case DocumentSaveFail:
		return "保存失败"
	case DocumentDelFail:
		return "删除失败"
	case DocumentNotFind:
		return "数据不存在"
	case DocumentUpdateFail:
		return "数据更新失败"
	case DocumentDeleteFail:
		return "数据删除失败"
	case DateConvertError:
		return "日期转换错误"
	case DBTxBeginError:
		return "事务开启失败"
	case DBTxCommitError:
		return "事务提交失败"
	case DBTxRollbackError:
		return "事务回滚失败"
	case DBTxSavePointError:
		return "事务保存点失败"
	case DBTxRollbackToError:
		return "事务回滚至保存点失败"
	case OpLogCreateError:
		return "创建操作日志失败"
	default:
		return "未知错误"
	}
}

func (e SocError) Val() (val int32) {
	switch e {
	case Success:
		val = 200
	case NotAuth:
		val = 400
	case ErrorOnLogin:
		val = 401
	case ErrorVerifyToken:
		val = 402
	case ErrorReq:
		val = 10000
	case ErrorNoData:
		val = 90000
	case ErrorOnDataFind:
		val = 10002
	case ErrorOnDataSave:
		val = 10004
	case ErrorOnDataDelete:
		val = 10006
	case ErrorOnShallowCopy:
		val = 10008
	case ErrorGenerateToken:
		val = 10010
	case ErrorParserToken:
		val = 10012
	case LogOutError:
		val = 10013
	case PleaseRepeatLogin:
		val = 10014
	case DataConvertError:
		val = 10015
	case LookupError:
		val = 10016
	case UserExist:
		val = 10100
	case UserNotExist:
		val = 10101
	case OldPWDIncorrect:
		val = 10102
	case ChangePWDError:
		val = 10103
	case NoChangeOtherPWD:
		val = 10104
	case RoleExist:
		val = 10200
	case RoleNotExist:
		val = 10201
	case RoleDisable:
		val = 10202
	case RoleHasDel:
		val = 10203
	case UserRoleExist:
		val = 10300
	case UserRoleNotExist:
		val = 10302
	case MenuExist:
		val = 10400
	case MenuNotExist:
		val = 10401
	case ESSearchError:
		val = 10801
	case ESCreateIndexError:
		val = 10802
	case ESIndexExist:
		val = 10803
	case DocumentSaveFail:
		val = 10804
	case DocumentDelFail:
		val = 10805
	case DocumentNotFind:
		val = 10806
	case DocumentUpdateFail:
		val = 10807
	case DocumentDeleteFail:
		val = 10808
	case DateConvertError:
		val = 10900
	case DBTxBeginError:
		val = 501
	case DBTxCommitError:
		val = 502
	case DBTxRollbackError:
		val = 503
	case DBTxSavePointError:
		val = 504
	case DBTxRollbackToError:
		val = 505
	case OpLogCreateError:
		val = 11500
	default:

	}
	return
}
