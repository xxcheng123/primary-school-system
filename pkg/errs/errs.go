package errs

import "google.golang.org/grpc/codes"

/*
 * @Author: xxcheng
 * @Email: developer@xxcheng.cn
 * @Date: 2024-03-08 14:57:20
 * @LastEditTime: 2024-03-08 16:45:53
 */

type Code uint32

const (
	NotFound         Code = 404
	ParamParseFailed Code = 9998

	Finally Code = 9999
)

const (
	MySQLInternalError Code = 1001 + iota
	MySQLRegisterError
)

const (
	UserNotExist Code = 2000 + iota
	PasswordIncorrect
	UserExpired
	UserBan

	CategoryNotExist

	NotSupportedOperate
	OperateFailure
)

var code2Message = map[Code]string{
	NotFound:         "未找到",
	Finally:          "内部错误",
	ParamParseFailed: "参数解析失败",

	MySQLInternalError: "数据库错误",
	MySQLRegisterError: "注册失败",

	UserNotExist:      "用户不存在",
	PasswordIncorrect: "密码错误",
	UserExpired:       "用户过期",
	UserBan:           "用户被封禁",

	CategoryNotExist: "分类不存在",

	NotSupportedOperate: "不支持的操作",
	OperateFailure:      "操作失败",
}

func (c Code) Error() string {
	if m, ok := code2Message[c]; ok {
		return m
	} else {
		return code2Message[Finally]
	}
}
func (c Code) Code() codes.Code {
	return codes.Code(c)
}
