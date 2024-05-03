package conversation_msg

import (
	"context"
	"github.com/OpenIMSDK/protocol/sdkws"
	"github.com/friendlyhank/openim-sdk-core-annotated/pkg/constant"
	"github.com/friendlyhank/openim-sdk-core-annotated/sdk_struct"
)

// SendMessage - 发送消息结构体
func (c *Conversation) SendMessage(ctx context.Context, s *sdk_struct.MsgStruct) (*sdk_struct.MsgStruct, error) {
	return c.sendMessageToServer(ctx, s)
}

// sendMessageToServer - 发送消息到服务器
func (c *Conversation) sendMessageToServer(ctx context.Context, s *sdk_struct.MsgStruct) (*sdk_struct.MsgStruct, error) {
	var wsMsgData sdkws.MsgData
	var sendMsgResp sdkws.UserSendMsgResp

	err := c.LongConnMgr.SendReqWaitResp(ctx, &wsMsgData, constant.SendMsg, &sendMsgResp)
	if err != nil {
	}
	return s, nil
}
