package sdk_struct

/*
 * sdk结构体
 */

// IMConfig - im配置信息
type IMConfig struct {
	PlatformID          int32  `json:"platformID"`          // 平台id
	ApiAddr             string `json:"apiAddr"`             // api地址
	WsAddr              string `json:"wsAddr"`              // 长链接地址
	DataDir             string `json:"dataDir"`             // 数据目录
	LogLevel            uint32 `json:"logLevel"`            // 日志等级
	IsLogStandardOutput bool   `json:"isLogStandardOutput"` // 是否输出到标准输出
	LogFilePath         string `json:"logFilePath"`         // 日志文件路径
}

// MsgStruct - 消息相关结构体
type MsgStruct struct {
	CreateTime       int64     `json:"createTime"`       // 创建时间
	SendTime         int64     `json:"sendTime"`         // 发送时间
	SendID           string    `json:"sendID,omitempty"` // 发送id
	RecvID           string    `json:"recvID,omitempty"` // 接收id
	ContentType      int32     `json:"contentType"`      // 内容类型
	SenderPlatformID int32     `json:"senderPlatformID"`
	SenderNickname   string    `json:"senderNickname,omitempty"` // 发送人昵称
	SenderFaceURL    string    `json:"senderFaceUrl,omitempty"`  // 发送人头像
	Content          string    `json:"content,omitempty"`        // 内容
	IsRead           bool      `json:"isRead"`                   // 是否已读
	Status           int32     `json:"status"`                   // 消息状态
	TextElem         *TextElem `json:"textElem,omitempty"`       // 文本消息
}

// TextElem - 文本消息返回
type TextElem struct {
	Content string `json:"content"`
}
