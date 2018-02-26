# aliyun-go-sms
阿里云短信 api 工具包，使用 go 实现

```
go get github.com/xuruiray/aliyun-go-sms
```

```
package main
import github.com/xuruiray/aliyun-go-sms

func main(){

    // 填写相关信息
    messageInfo := MessageBody{
        AccessKeyID:     "AccessKeyID",
        AccessKeySecret: "AccessKeySecret",
        PhoneNumbers:    "PhoneNumber",
        SignName:        "SignName",
        TemplateCode:    "TemplateCode",
        TemplateParam:   "TemplateParam",
    }
    
    // 调用短信接口发送短信
    err := SendMessage(messageInfo)
    if err != nil{
        //错误处理
    }

}

```

初步可用，待进一步完善

[阿里云 api 接口文档](https://help.aliyun.com/document_detail/56189.html?spm=a2c4g.11186623.6.580.LaGR0O)