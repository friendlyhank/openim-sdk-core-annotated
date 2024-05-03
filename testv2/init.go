package testv2

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/OpenIMSDK/protocol/constant"
	"github.com/friendlyhank/openim-sdk-core-annotated/open_im_sdk"
	"github.com/friendlyhank/openim-sdk-core-annotated/pkg/ccontext"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var (
	ctx context.Context
)

func init() {
	fmt.Println("------------------------>>>>>>>>>>>>>>>>>>> test v2 init funcation <<<<<<<<<<<<<<<<<<<------------------------")
	rand.Seed(time.Now().UnixNano())
	config := getConf(APIADDR, WSADDR)
	config.DataDir = "./"
	configData, err := json.Marshal(config)
	if err != nil {
		panic(err)
	}
	isInit := open_im_sdk.InitSDK()
	if !isInit {
		panic("init sdk failed")
	}
	ctx = open_im_sdk.UserForSDK.Context()
	ctx = ccontext.WithOperationID(ctx, "initOperationID_"+strconv.Itoa(int(time.Now().UnixMilli())))
	token, err := GetUserToken(ctx, UserID)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}

// GetUserToken - 获取用户token
func GetUserToken(ctx context.Context, userID string) (string, error) {
	jsonReqData, err := json.Marshal(map[string]any{
		"userID":     userID,
		"platformID": constant.LinuxPlatformID,
		"secret":     "openIM123",
	})
	if err != nil {
		return "", err
	}
	path := APIADDR + "/auth/user_token"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path, bytes.NewReader(jsonReqData))
	if err != nil {
		return "", err
	}
	req.Header.Set("operationID", ctx.Value("operationID").(string))
	client := http.Client{Timeout: time.Second * 3}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	type Result struct {
		ErrCode int    `json:"errCode"`
		ErrMsg  string `json:"errMsg"`
		ErrDlt  string `json:"errDlt"`
		Data    struct {
			Token             string `json:"token"`
			ExpireTimeSeconds int    `json:"expireTimeSeconds"`
		} `json:"data"`
	}
	var result Result
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}
	if result.ErrCode != 0 {
		return "", fmt.Errorf("errCode:%d, errMsg:%s, errDlt:%s", result.ErrCode, result.ErrMsg, result.ErrDlt)
	}
	return result.Data.Token, nil
}
