package conversation_msg

import (
	"context"
	"github.com/friendlyhank/openim-sdk-core-annotated/internal/interaction"
	"github.com/friendlyhank/openim-sdk-core-annotated/pkg/ccontext"
)

/*
 * 会话信息
 */

type Conversation struct {
	*interaction.LongConnMgr // 长链接
	platformID               int32
}

func NewConversation(ctx context.Context, longConnMgr *interaction.LongConnMgr) *Conversation {
	info := ccontext.Info(ctx)
	n := &Conversation{
		LongConnMgr: longConnMgr,
		platformID:  info.PlatformID(),
	}
	return n
}
