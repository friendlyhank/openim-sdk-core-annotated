package interaction

// GeneralWsResp - ws响应结构体
type GeneralWsResp struct {
	ReqIdentifier int    `json:"reqIdentifier"`
	ErrCode       int    `json:"errCode"`
	ErrMsg        string `json:"errMsg"`
	OperationID   string `json:"operationID"`
	Data          []byte `json:"data"`
}

// GeneralWsReq - ws请求结构体
type GeneralWsReq struct {
	ReqIdentifier int    `json:"reqIdentifier"` // 请求认证
	Token         string `json:"token"`
	SendID        string `json:"sendID"`
	OperationID   string `json:"operationID"`
	Data          []byte `json:"data"`
}
