package open_im_sdk

import "context"

func InitSDK(config string) bool {
	if UserForSDK != nil {
		return true
	}
	UserForSDK = new(LoginMgr)
	return true
}

// Login - 登录
func (u *LoginMgr) Login(ctx context.Context, userID, token string) error {

}
