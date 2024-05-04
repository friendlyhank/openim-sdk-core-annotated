package interaction

import (
	"context"
	"errors"
	"github.com/OpenIMSDK/tools/errs"
	"github.com/OpenIMSDK/tools/log"
	"github.com/friendlyhank/openim-sdk-core-annotated/pkg/ccontext"
	"github.com/friendlyhank/openim-sdk-core-annotated/pkg/sdkerrs"
	"github.com/golang/protobuf/proto"
)

/*
 * 长链接信息的核心
 */

type LongConnMgr struct {
	// The long connection,can be set tcp or websocket.
	conn LongConn // 用接口设计可以多种实体实现
	// Buffered channel of outbound messages.
	send chan Message // 消息结构体，长链接的消息异步获取消息的方式解耦
	ctx  context.Context
}

// Message - ws消息结构体
type Message struct {
	Message GeneralWsReq
	Resp    chan *GeneralWsResp
}

// NewLongConnMgr - 初始化长链接管理
func NewLongConnMgr(ctx context.Context) *LongConnMgr {
	l := &LongConnMgr{}
	l.send = make(chan Message, 10)
	l.ctx = ctx
	return l
}

// Run - 长链接启动
func (c *LongConnMgr) Run(ctx context.Context) {
	//fmt.Println(mcontext.GetOperationID(ctx), "login run", string(debug.Stack()))
	go c.readPump(ctx)
	go c.writePump(ctx)
	go c.heartbeat(ctx)
}

// SendReqWaitResp - 发送消息并等待结果(传入具体消息结构体)
func (c *LongConnMgr) SendReqWaitResp(ctx context.Context, m proto.Message, reqIdentifier int, resp proto.Message) error {
	data, err := proto.Marshal(m)
	if err != nil {
		return sdkerrs.ErrArgs
	}
	msg := Message{
		Message: GeneralWsReq{
			ReqIdentifier: reqIdentifier,
			SendID:        ccontext.Info(ctx).UserID(),
			OperationID:   ccontext.Info(ctx).OperationID(),
			Data:          data,
		},
		Resp: make(chan *GeneralWsResp, 1),
	}
	c.send <- msg
	log.ZDebug(ctx, "send message to send channel success", "msg", m, "reqIdentifier", reqIdentifier)
	select {
	case <-ctx.Done():
		return sdkerrs.ErrCtxDeadline
	case v, ok := <-msg.Resp:
		if !ok {
			return errors.New("response channel closed")
		}
		if v.ErrCode != 0 {
			return errs.NewCodeError(v.ErrCode, v.ErrMsg)
		}
		if err := proto.Unmarshal(v.Data, resp); err != nil {
			return sdkerrs.ErrArgs
		}
		return nil
	}
}

// readPump - 读消息循环
func (c *LongConnMgr) readPump(ctx context.Context) {

}

// writePump - 写消息循环
func (c *LongConnMgr) writePump(ctx context.Context) {

	defer func() {
		close(c.send)
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case message, ok := <-c.send: // 发送消息
			if !ok {
				// The hub closed the channel. channel关闭了
			}
			log.ZDebug(c.ctx, "writePump recv message", "reqIdentifier", message.Message.ReqIdentifier,
				"operationID", message.Message.OperationID, "sendID", message.Message.SendID)
		}
	}

}

// heartbeat - 长链接心跳
func (c *LongConnMgr) heartbeat(ctx context.Context) {

}
