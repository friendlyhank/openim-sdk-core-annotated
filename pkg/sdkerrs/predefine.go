package sdkerrs

import "github.com/OpenIMSDK/tools/errs"

var (
	ErrArgs        = errs.NewCodeError(ArgsError, "ArgsError")                               // 参数错误
	ErrCtxDeadline = errs.NewCodeError(CtxDeadlineExceededError, "CtxDeadlineExceededError") // 上下文超时

	//消息相关
	ErrMsgContentTypeNotSupport = errs.NewCodeError(MsgContentTypeNotSupportError, "contentType not support currently") // 不支持的消息类型
)
