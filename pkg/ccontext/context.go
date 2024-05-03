package ccontext

import (
	"context"
	"github.com/OpenIMSDK/tools/mcontext"
	"github.com/friendlyhank/openim-sdk-core-annotated/sdk_struct"
)

/*
 * 上下文信息设置
 */

// GlobalConfig - 每个用户的全局配置，放在上下文
type GlobalConfig struct {
	UserID string
	Token  string

	sdk_struct.IMConfig
}

type ContextInfo interface {
	UserID() string
	Token() string
	PlatformID() int32
	ApiAddr() string
	WsAddr() string
	DataDir() string
	LogLevel() uint32
	OperationID() string
}

func Info(ctx context.Context) ContextInfo {
	conf := ctx.Value(GlobalConfigKey{}).(*GlobalConfig)
	return &info{
		conf: conf,
		ctx:  ctx,
	}
}

// WithInfo - 将用户配置信息放到上下文
func WithInfo(ctx context.Context, conf *GlobalConfig) context.Context {
	return context.WithValue(ctx, GlobalConfigKey{}, conf)
}

// WithOperationID - operationID
func WithOperationID(ctx context.Context, operationID string) context.Context {
	return mcontext.SetOperationID(ctx, operationID)
}

type GlobalConfigKey struct{}

type info struct {
	conf *GlobalConfig
	ctx  context.Context
}

func (i info) UserID() string {
	return i.conf.UserID
}

func (i info) Token() string {
	return i.conf.Token
}

func (i info) PlatformID() int32 {
	return i.conf.PlatformID
}

func (i info) ApiAddr() string {
	return i.conf.ApiAddr
}

func (i info) WsAddr() string {
	return i.conf.WsAddr
}

func (i info) DataDir() string {
	return i.conf.DataDir
}

func (i info) LogLevel() uint32 {
	return i.conf.LogLevel
}

func (i info) OperationID() string {
	return mcontext.GetOperationID(i.ctx)
}
