package conversation_msg

import (
	"context"
	"github.com/friendlyhank/openim-sdk-core-annotated/internal/interaction"
)

/*
 * 会话信息
 */

type Conversation struct {
	*interaction.LongConnMgr // 长链接
}

func NewConversation(ctx context.Context, longConnMgr *interaction.LongConnMgr) *Conversation {
	n := &Conversation{
		LongConnMgr: longConnMgr,
	}
	return n
}
