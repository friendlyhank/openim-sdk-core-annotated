package open_im_sdk

import (
	"context"
	"github.com/OpenIMSDK/tools/log"
	conv "github.com/friendlyhank/openim-sdk-core-annotated/internal/conversation_msg"
	"github.com/friendlyhank/openim-sdk-core-annotated/internal/interaction"
	"github.com/friendlyhank/openim-sdk-core-annotated/pkg/ccontext"
	"github.com/friendlyhank/openim-sdk-core-annotated/sdk_struct"
	"sync"
)

/*
 * im-sdk的核心初始化
 */

const (
	LogoutStatus = iota + 1 // 登出状态
	Logging                 // 正在登录状态
	Logged                  // 已登录
)

var (
	// UserForSDK Client-independent user class
	UserForSDK *LoginMgr
)

type LoginMgr struct {
	conversation *conv.Conversation // 会话相关

	w           sync.Mutex
	loginStatus int // 登录状态

	longConnMgr *interaction.LongConnMgr // 长链接信息(关键)

	ctx    context.Context // 上下文信息
	cancel context.CancelFunc
	info   *ccontext.GlobalConfig // 每个用户的配置信息
}

// ImConfig - 获取配置信息
func (u *LoginMgr) ImConfig() sdk_struct.IMConfig {
	return sdk_struct.IMConfig{
		PlatformID: u.info.PlatformID,
		ApiAddr:    u.info.ApiAddr,
		WsAddr:     u.info.WsAddr,
		DataDir:    u.info.DataDir,
		LogLevel:   u.info.LogLevel,
	}
}

// Conversation - 会话信息相关
func (u *LoginMgr) Conversation() *conv.Conversation {
	return u.conversation
}

// getLoginStatus - 获取登录状态
func (u *LoginMgr) getLoginStatus(_ context.Context) int {
	u.w.Lock()
	defer u.w.Unlock()
	return u.loginStatus
}

// setLoginStatus - 设置登录状态
func (u *LoginMgr) setLoginStatus(status int) {
	u.w.Lock()
	defer u.w.Unlock()
	u.loginStatus = status
}

// login - 登录
func (u *LoginMgr) login(ctx context.Context, userID, token string) error {
	// 设置用户上下文
	u.info.UserID = userID
	u.info.Token = token
	log.ZDebug(ctx, "login start... ", "userID", userID, "token", token)

	u.conversation = conv.NewConversation(ctx, u.longConnMgr)
	u.run(ctx) // 登录之后启动相关长链接
	return nil
}

// run - 登录后启动相关信息
func (u *LoginMgr) run(ctx context.Context) {
	u.longConnMgr.Run(ctx) // 长链接设置
}

// InitSDK - 初始化用户sdk
func (u *LoginMgr) InitSDK(config sdk_struct.IMConfig) bool {
	u.info = &ccontext.GlobalConfig{}
	u.info.IMConfig = config
	u.initResources()
	return true
}

func (u *LoginMgr) Context() context.Context {
	return u.ctx
}

// initResources - 初始化相关资源(这个也是核心)
func (u *LoginMgr) initResources() {
	ctx := ccontext.WithInfo(context.Background(), u.info)
	u.ctx, u.cancel = context.WithCancel(ctx)
	u.longConnMgr = interaction.NewLongConnMgr(u.ctx)
	u.setLoginStatus(LogoutStatus)
}
