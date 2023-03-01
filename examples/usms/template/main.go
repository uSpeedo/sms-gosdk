package main

import (
	"fmt"
	"github.com/uSpeedo/usms-sdk-go/private/utils"
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
	credential.AccessKeyId = "..."
	credential.AccessKeySecret = "..."

	client := usms.NewClient(&cfg, &credential)

	// send request
	req := client.NewCreateUSMSTemplateRequest()

	req.AccountId = um.Int(1)
	req.Action = um.String("CreateUSMSTemplate")
	req.Purpose = um.Int(1)
	req.TemplateName = um.String("sdk test")
	req.Template = um.String("sdk template example")
	// add header
	req.SetNonce(utils.RandStr(10))
	req.SetAccessKeyId(credential.AccessKeyId)
	req.SetSignature(credential.CreateSign(makeCreateTemplateParamMap(req)))
	t, _ := time.ParseDuration("-2m")
	req.SetTimestamp(time.Now().Add(t).Unix())
	resp, err := client.CreateUSMSTemplate(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", resp)
}

func makeCreateTemplateParamMap(req *usms.CreateUSMSTemplateRequest) map[string]interface{} {
	m := make(map[string]interface{}, 0)
	m["AccountId"] = req.AccountId
	m["Action"] = req.Action
	m["Purpose"] = req.Purpose
	m["TemplateName"] = req.TemplateName
	m["Template"] = req.Template
	return m
}
