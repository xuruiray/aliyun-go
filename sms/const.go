package sms

const (
	// Host api接口地址
	Host = "dysmsapi.aliyuncs.com"
	// TimestampFormat 时间戳格式
	TimestampFormat = "2006-01-02T15:04:05Z"
	// SignatureMethod 签名加密方式
	SignatureMethod = "HMAC-SHA1"
	// SignatureVersion 签名版本号 由服务方提供
	SignatureVersion = "1.0"
	// Action 需要使用的 api 功能 ，发送短信
	Action = "SendSms"
	// Format 返回值格式
	Format = "JSON"
	// Version api 版本
	Version = "2017-05-25"
	// RegionId api 支持的服务器区域ID
	RegionId = "cn-hangzhou"
)
