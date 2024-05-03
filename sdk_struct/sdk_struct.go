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
	TextElem *TextElem `json:"textElem,omitempty"` // 文本消息
}

// TextElem - 文本消息返回
type TextElem struct {
	Content string `json:"content"`
}
