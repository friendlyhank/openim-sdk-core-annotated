package constant

const BigVersion = "v3"
const UpdateVersion = ".0.0"
const SdkVersion = "openim-sdk-core-"

func GetSdkVersion() string {
	return SdkVersion + BigVersion + UpdateVersion
}

const (
	Text = 101 // 文本类型消息

	MsgStatusSending     = 1 // 消息发送中
	MsgStatusSendSuccess = 2 // 消息发送成功
	MsgStatusSendFailed  = 3 // 消息发送失败
	MsgStatusHasDeleted  = 4 // 消息已经删除
)

const (
	SendMsg = 1003 // 发送消息(请求标识)
)
