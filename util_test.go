package aliyun_go_sms

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_packageParams(t *testing.T) {

	tests := []struct {
		name string
		args MessageBody
		want string
	}{
		{
			name: "正常流程",
			args: MessageBody{
				AccessKeyID:     "id",
				AccessKeySecret: "secret",
				PhoneNumbers:    "187xxxxxxx3",
				SignName:        "签名",
				TemplateCode:    "模版编号",
				TemplateParam:   "{\"模版参数\":\"模版参数\"}",
			},
			//去掉了 SignatureNonce 与 Timestamp 与时间相关不好测试
			want: "AccessKeyId=id&Action=SendSms&Format=JSON&PhoneNumbers=187xxxxxxx3&RegionId=cn-hangzhou&SignName=%E7%AD%BE%E5%90%8D&SignatureMethod=HMAC-SHA1&" +
				"SignatureVersion=1.0&TemplateCode=%E6%A8%A1%E7%89%88%E7%BC%96%E5%8F%B7&TemplateParam=%7B%22%E6%A8%A1%E7%89%88%E5%8F%82%E6%95%B0%22%3A%22%E6%A8%A1%E7%89%88%E5%8F%82%E6%95%B0%22%7D&" +
				"Version=2017-05-25",
		},
	}
	for _, v := range tests {
		result := packageParams(v.args)
		index := strings.Index(result, "SignatureNonce=")
		result = result[:index] + result[index+26:]
		index = strings.Index(result, "Timestamp=")
		result = result[:index] + result[index+35:]
		assert.Equal(t, v.want, result, "name: %v", v.name)
	}
}

func Test_genSignature(t *testing.T) {
	tests := []struct {
		name            string
		accessKeySecret string
		requestStr      string
		want            string
	}{
		{
			name:            "正常流程",
			accessKeySecret: "testSecret",
			requestStr:      "AccessKeyId%3DtestId%26Action%3DSendSms%26Format%3DXML%26OutId%3D123%26PhoneNumbers%3D15300000001%26RegionId%3Dcn-hangzhou%26SignName%3D%25E9%2598%25BF%25E9%2587%258C%25E4%25BA%2591%25E7%259F%25AD%25E4%25BF%25A1%25E6%25B5%258B%25E8%25AF%2595%25E4%25B8%2593%25E7%2594%25A8%26SignatureMethod%3DHMAC-SHA1%26SignatureNonce%3D45e25e9b-0a6f-4070-8c85-2956eda1b466%26SignatureVersion%3D1.0%26TemplateCode%3DSMS_71390007%26TemplateParam%3D%257B%2522customer%2522%253A%2522test%2522%257D%26Timestamp%3D2017-07-12T02%253A42%253A19Z%26Version%3D2017-05-25",
			want:            "7KzUM9Y6gbM8+BUjbu1MJHDV2ZA=",
		},
	}
	for _, v := range tests {
		result := genSignature(v.accessKeySecret, v.requestStr)
		assert.Equal(t, v.want, result, "name: %v", v.name)
	}
}

func Test_sendRequest(t *testing.T) {
	tests := []struct {
		name       string
		requestStr string
		sign       string
		wantErr    error
	}{
		{
			name:       "仅测试能否跑通",
			requestStr: "AccessKeyId=id&Action=SendSms&Format=JSON&PhoneNumbers=187xxxxxxx3&RegionId=cn-hangzhou&SignName=%E7%AD%BE%E5%90%8D&SignatureMethod=HMAC-SHA1&SignatureNonce=1519636246&SignatureVersion=1.0&TemplateCode=%E6%A8%A1%E7%89%88%E7%BC%96%E5%8F%B7&TemplateParam=%7B%22%E6%A8%A1%E7%89%88%E5%8F%82%E6%95%B0%22%3A%22%E6%A8%A1%E7%89%88%E5%8F%82%E6%95%B0%22%7D&Timestamp=2018-02-26T9%3A29%3A46Z&Version=2017-05-25",
			sign:       "sign",
			wantErr:    errors.New("InvalidTimeStamp.Expired"),
		},
	}
	for _, v := range tests {
		err := sendRequest(v.requestStr, v.sign)
		assert.Equal(t, v.wantErr, err, "name: %v", v.name)
	}
}
