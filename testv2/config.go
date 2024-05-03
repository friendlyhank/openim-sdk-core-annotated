package testv2

import (
	"github.com/OpenIMSDK/protocol/constant"
	"github.com/friendlyhank/openim-sdk-core-annotated/sdk_struct"
)

const (
	APIADDR = "http://127.0.0.1:10002"
	WSADDR  = "ws://http://127.0.0.1:10001"
	UserID  = "11151515151"
)

func getConf(APIADDR, WSADDR string) sdk_struct.IMConfig {
	var cf sdk_struct.IMConfig
	cf.ApiAddr = APIADDR
	cf.WsAddr = WSADDR
	cf.DataDir = "../"
	cf.LogLevel = 6
	cf.PlatformID = constant.LinuxPlatformID
	return cf
}
