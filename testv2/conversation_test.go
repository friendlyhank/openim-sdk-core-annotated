package testv2

import (
	"github.com/friendlyhank/openim-sdk-core-annotated/open_im_sdk"
	"testing"
)

/*
 * 会话相关测试用例
 */

// Test_SendMessage - 测试发送消息
func Test_SendMessage(t *testing.T) {
	open_im_sdk.UserForSDK.Conversation().CreateTextMessage(ctx, "textMsg")
	_, err := open_im_sdk.UserForSDK.Conversation().SendMessage()
}
