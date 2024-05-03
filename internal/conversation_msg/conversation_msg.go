package conversation_msg

import "context"

/*
 * 会话信息
 */

type Conversation struct {
}

func NewConversation(ctx context.Context) *Conversation {
	n := &Conversation{}
	return n
}
