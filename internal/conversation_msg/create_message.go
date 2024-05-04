package conversation_msg

import (
	"context"
	"github.com/friendlyhank/openim-sdk-core-annotated/pkg/constant"
	"github.com/friendlyhank/openim-sdk-core-annotated/sdk_struct"
)

// CreateTextMessage - 生成text的消息体
func (c *Conversation) CreateTextMessage(ctx context.Context, text string) (*sdk_struct.MsgStruct, error) {
	s := sdk_struct.MsgStruct{}
	c.initBasicInfo(ctx, &s, constant.Text)
	s.TextElem = &sdk_struct.TextElem{Content: text}
	return &s, nil
}
