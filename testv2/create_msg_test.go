package testv2

import (
	"github.com/friendlyhank/openim-sdk-core-annotated/open_im_sdk"
	"testing"
)

func Test_CreateTextMessage(t *testing.T) {
	message, err := open_im_sdk.UserForSDK.Conversation().CreateTextMessage(ctx, "textMsg")
	if err != nil {
		t.Error(err)
	}
	t.Log(message)
}
