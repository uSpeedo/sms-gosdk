package main

import (
	"fmt"
	"time"

	"github.com/uSpeedo/usms-sdk-go/services/usms"
	"github.com/uSpeedo/usms-sdk-go/um"
	"github.com/uSpeedo/usms-sdk-go/um/auth"
	"github.com/uSpeedo/usms-sdk-go/um/config"
	"github.com/uSpeedo/usms-sdk-go/um/log"
)

func main() {
	cfg := config.NewConfig()
	cfg.LogLevel = log.DebugLevel

	credential := auth.NewCredential()
	credential.AccessKeySecret = "..."

	client := usms.NewClient(&cfg, &credential)

	// send request
	req := client.NewCreateUSMSTemplateRequest()

	req.AccountId = um.Int(1)
	req.Action = um.String("CreateUSMSTemplate")
	req.Purpose = um.Int(1)
	req.TemplateName = um.String("Test")
	req.Template = um.String("测试模板")
	req.International = um.Bool(true)
	// add header
	req.SetNonce("hz3xevqz")
	req.SetAccessKeyId("b9c3be2916bab10f219ec24138e5c27d")
	req.SetSignature("23756362a77cd0a1f798e574b644664c23df6dfa")
	t, _ := time.ParseDuration("-2m")
	req.SetTimestamp(time.Now().Add(t).Unix())
	resp, err := client.CreateUSMSTemplate(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", resp)
}
