package aliyun_go_sms

// MessageBody 短信相关信息
type MessageBody struct {
	// Access Key ID 和 Access Key Secret 是访问阿里云API的密钥
	AccessKeyID     string
	AccessKeySecret string
	// PhoneNumbers 手机号
	PhoneNumbers string
	// SignName 短信签名
	SignName string
	// TemplaceCode 短信模板编号
	TemplateCode string
	// TemplateParam 短信模版中参数 格式为 Json
	TemplateParam string
}

// SendMessage 短信发送函数
func SendMessage(body MessageBody) error {
	requestStr := packageParams(body)
	sign := genSignature(body.AccessKeySecret, requestStr)
	err := sendRequest(requestStr, sign)
	return err
}
