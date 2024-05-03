package open_im_sdk

import (
	"context"
	conv "github.com/friendlyhank/openim-sdk-core-annotated/internal/conversation_msg"
)

/*
 * im-sdk的核心初始化
 */

var (
	// UserForSDK Client-independent user class
	UserForSDK *LoginMgr
)

type LoginMgr struct {
	conversation *conv.Conversation // 会话相关
	ctx          context.Context    // 上下文信息
}

// Conversation - 会话信息相关
func (u *LoginMgr) Conversation() *conv.Conversation {
	return u.conversation
}

func (u *LoginMgr) Context() context.Context {
	return u.ctx
}

// login - 登录
func (u *LoginMgr) login(ctx context.Context, userID, token string) error {
	u.conversation = conv.NewConversation(ctx)
	return nil
}
