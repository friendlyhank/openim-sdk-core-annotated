package sdkerrs

import "github.com/OpenIMSDK/tools/errs"

var (
	ErrArgs = errs.NewCodeError(ArgsError, "ArgsError")
)
