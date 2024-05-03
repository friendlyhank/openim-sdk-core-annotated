package constant

const BigVersion = "v3"
const UpdateVersion = ".0.0"
const SdkVersion = "openim-sdk-core-"

func GetSdkVersion() string {
	return SdkVersion + BigVersion + UpdateVersion
}

const (
	SendMsg = 1003 // 发送消息(请求标识)
)
