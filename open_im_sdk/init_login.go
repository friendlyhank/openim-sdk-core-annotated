package open_im_sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/OpenIMSDK/tools/log"
	"github.com/OpenIMSDK/tools/mcontext"
	"github.com/friendlyhank/openim-sdk-core-annotated/pkg/constant"
	"github.com/friendlyhank/openim-sdk-core-annotated/sdk_struct"
	"strings"
)

/*
 * 初始化登录信息
 */

const (
	// 日志相关配置
	rotateCount  uint = 0
	rotationTime uint = 24
)

// GetSdkVersion - 获取sdk版本
func GetSdkVersion() string {
	return constant.GetSdkVersion()
}

func InitSDK(operationID string, config string) bool {
	if UserForSDK != nil {
		fmt.Println(operationID, "Initialize multiple times, use the existing ", UserForSDK, " Previous configuration ", UserForSDK.ImConfig(), " now configuration: ", config)
		return true
	}
	var configArgs sdk_struct.IMConfig
	if err := json.Unmarshal([]byte(config), &configArgs); err != nil {
		fmt.Println(operationID, "Unmarshal failed ", err.Error(), config)
		return false
	}
	if configArgs.PlatformID == 0 {
		return false
	}
	if err := log.InitFromConfig("open-im-sdk-core", "", int(configArgs.LogLevel), configArgs.IsLogStandardOutput, false, configArgs.LogFilePath, rotateCount, rotationTime); err != nil {
		fmt.Println(operationID, "log init failed ", err.Error())
	}
	fmt.Println("init log success")
	// localLog.NewPrivateLog("", configArgs.LogLevel)
	ctx := mcontext.NewCtx(operationID)
	if !strings.Contains(configArgs.ApiAddr, "http") {
		log.ZError(ctx, "api is http protocol, api format is invalid", nil)
		return false
	}
	if !strings.Contains(configArgs.WsAddr, "ws") {
		log.ZError(ctx, "ws is ws protocol, ws format is invalid", nil)
		return false
	}

	log.ZInfo(ctx, "InitSDK info", "config", configArgs, "sdkVersion", GetSdkVersion())
	UserForSDK = new(LoginMgr)
	return UserForSDK.InitSDK(configArgs)
}

// Login - 登录
func (u *LoginMgr) Login(ctx context.Context, userID, token string) error {
	return u.login(ctx, userID, token)
}
