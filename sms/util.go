package sms

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type ApiResult struct {
	Recommend string `json:"Recommend"`
	Message   string `json:"Message"`
	RequestId string `json:"RequestId"`
	HostId    string `json:"HostId"`
	Code      string `json:"Code"`
}

// packageParams 将参数进行url编码
func packageParams(body MessageBody) (requestStr string) {

	u := url.Values{}
	u.Set("AccessKeyId", body.AccessKeyID)
	u.Set("Action", Action)
	u.Set("Format", Format)
	u.Set("PhoneNumbers", body.PhoneNumbers)
	u.Set("RegionId", RegionId)
	u.Set("SignName", body.SignName)
	u.Set("SignatureMethod", SignatureMethod)
	u.Set("SignatureNonce", strconv.Itoa(int(time.Now().Unix())))
	u.Set("SignatureVersion", SignatureVersion)
	u.Set("TemplateCode", body.TemplateCode)
	u.Set("TemplateParam", body.TemplateParam)
	u.Set("Timestamp", time.Now().UTC().Format(TimestampFormat))
	u.Set("Version", Version)

	// 将 value 进行 url编码，并按照 key 字母顺序，以 & 连接
	requestStr = u.Encode()

	//加号（+）替换成 %20、星号（*）替换成 %2A、%7E 替换回波浪号（~）
	requestStr = strings.Replace(requestStr, "%7E", "~", -1)
	requestStr = strings.Replace(requestStr, "+", "%20", -1)
	requestStr = strings.Replace(requestStr, "*", "%2A", -1)

	return
}

// genSignature 根据 accessKeySecret requestStr 获取包含签名的请求字符串
// 签名采用 HmacSHA1 + Base64
func genSignature(accessKeySecret string, requestStr string) string {

	accessKeySecret += "&"
	stringToSign := urlEncode(requestStr)
	stringToSign = "GET&%2F&" + stringToSign

	//hmac ,use sha1
	key := []byte(accessKeySecret)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(stringToSign))
	encodeString := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return encodeString
}

// sendRequest 发送 http 请求
func sendRequest(requestStr string, sign string) error {
	u := url.Values{}
	u.Set("Signature", sign)
	resp, err := http.Get("http://" + Host + "/?" + u.Encode() + "&" + requestStr)
	if err != nil {
		return err
	}

	// 读取响应
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// 解析 body 查看 API 是否返回错误
	var apiResult ApiResult
	json.Unmarshal(body, &apiResult)
	if apiResult.Code != "OK" {
		return errors.New(apiResult.Code)
	}

	return nil
}

//加号（+）替换成 %20、星号（*）替换成 %2A、%7E 替换回波浪号（~）
func urlEncode(requestStr string) string {
	u := url.Values{}
	u.Set("", requestStr)
	requestStr = u.Encode()[1:]
	requestStr = strings.Replace(requestStr, "%7E", "~", -1)
	requestStr = strings.Replace(requestStr, "+", "%20", -1)
	requestStr = strings.Replace(requestStr, "*", "%2A", -1)
	return requestStr
}
