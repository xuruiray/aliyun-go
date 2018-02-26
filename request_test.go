package aliyun_go_sms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	realAccessKeyID     = "-"
	realAccessKeySecret = "-"
	realPhoneNumber     = "187xxxxxxx3"
	realSignName        = "监控助手"
	realTemplateCode    = "SMS_1xxxxxxx3"
	realTemplateParam   = "{ \"alarmItem\":\"\", \"alarmSystem\":\"\", \"alarmTime\":\"\"}"
)

func TestSendMessage(t *testing.T) {

	tests := []struct {
		name    string
		body    MessageBody
		wantErr error
	}{
		{
			name: "正常流程",
			body: MessageBody{
				AccessKeyID:     realAccessKeyID,
				AccessKeySecret: realAccessKeySecret,
				PhoneNumbers:    realPhoneNumber,
				SignName:        realSignName,
				TemplateCode:    realTemplateCode,
				TemplateParam:   realTemplateParam,
			},
			wantErr: nil,
		},
	}

	for _, v := range tests {
		err := SendMessage(v.body)
		assert.Equal(t, v.wantErr, err, "name: %v", v.name)
	}
}
