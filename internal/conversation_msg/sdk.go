package conversation_msg

import (
	"context"
	"github.com/OpenIMSDK/protocol/sdkws"
	"github.com/OpenIMSDK/tools/utils"
	"github.com/friendlyhank/openim-sdk-core-annotated/pkg/constant"
	"github.com/friendlyhank/openim-sdk-core-annotated/pkg/sdkerrs"
	"github.com/friendlyhank/openim-sdk-core-annotated/sdk_struct"
	"github.com/jinzhu/copier"
)

// SendMessage - 发送消息结构体
func (c *Conversation) SendMessage(ctx context.Context, s *sdk_struct.MsgStruct) (*sdk_struct.MsgStruct, error) {

	switch s.ContentType {
	case constant.Text:
		s.Content = utils.StructToJsonString(s.TextElem)
	default:
		return nil, sdkerrs.ErrMsgContentTypeNotSupport
	}

	return c.sendMessageToServer(ctx, s)
}

// sendMessageToServer - 发送消息到服务器
func (c *Conversation) sendMessageToServer(ctx context.Context, s *sdk_struct.MsgStruct) (*sdk_struct.MsgStruct, error) {
	var wsMsgData sdkws.MsgData
	copier.Copy(&wsMsgData, s)
	wsMsgData.Content = []byte(s.Content)
	wsMsgData.CreateTime = s.CreateTime
	var sendMsgResp sdkws.UserSendMsgResp

	err := c.LongConnMgr.SendReqWaitResp(ctx, &wsMsgData, constant.SendMsg, &sendMsgResp)
	if err != nil {
	}
	s.SendTime = sendMsgResp.SendTime
	s.Status = constant.MsgStatusSendSuccess
	return s, nil
}

// initBasicInfo - 初始化基础信息
func (c *Conversation) initBasicInfo(ctx context.Context, message *sdk_struct.MsgStruct, contentType int32) error {
	message.CreateTime = utils.GetCurrentTimestampByMill()
	message.SendTime = message.CreateTime
	message.IsRead = false
	message.Status = constant.MsgStatusSending
	message.ContentType = contentType
	message.SenderPlatformID = c.platformID
	return nil
}
