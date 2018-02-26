package sms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	AccessKeyID     = "-"
	AccessKeySecret = "-"
	PhoneNumber     = "187xxxxxxx3"
	SignName        = "监控助手"
	TemplateCode    = "SMS_1xxxxxxx3"
	TemplateParam   = "{ \"alarmItem\":\"\", \"alarmSystem\":\"\", \"alarmTime\":\"\"}"
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
				AccessKeyID:     AccessKeyID,
				AccessKeySecret: AccessKeySecret,
				PhoneNumbers:    PhoneNumber,
				SignName:        SignName,
				TemplateCode:    TemplateCode,
				TemplateParam:   TemplateParam,
			},
			wantErr: nil,
		},
	}

	for _, v := range tests {
		err := SendMessage(v.body)
		assert.Equal(t, v.wantErr, err, "name: %v", v.name)
	}
}
