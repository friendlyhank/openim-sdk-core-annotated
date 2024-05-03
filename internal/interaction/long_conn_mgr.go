package interaction

import (
	"context"
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
	send chan Message // 消息结构体，为什么用channel
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
}

// readPump - 读消息循环
func (c *LongConnMgr) readPump(ctx context.Context) {

}

// writePump - 写消息循环
func (c *LongConnMgr) writePump(ctx context.Context) {

}

// heartbeat - 长链接心跳
func (c *LongConnMgr) heartbeat(ctx context.Context) {

}
