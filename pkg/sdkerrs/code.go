package sdkerrs

// 通用错误码
const (
	ArgsError                = 10002 //输入参数错误
	CtxDeadlineExceededError = 10003 //上下文超时

	//消息相关
	MsgContentTypeNotSupportError = 10205 //消息类型不支持
)
