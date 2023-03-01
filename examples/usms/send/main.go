package main

import (
	"fmt"
	"github.com/uSpeedo/usms-sdk-go/private/utils"

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
	req := client.NewSendUSMSMessageRequest()
	req.AccountId = um.Int(1)
	req.SigContent = um.String("...")
	req.TemplateId = um.String("UTA2***50501BD")
	req.PhoneNumbers = []string{
		"...",
		"...",
	}
	req.TemplateParams = []string{
		"424242",
	}
	// add header
	req.SetNonce(utils.RandStr(10))
	req.SetAccessKeyId(credential.AccessKeyId)
	req.SetSignature(credential.CreateSign(makeSendParamMap(req)))
	req.SetTimestamp(1669370992)
	resp, err := client.SendUSMSMessage(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", resp)
}

func makeSendParamMap(req *usms.SendUSMSMessageRequest) map[string]interface{} {
	m := make(map[string]interface{}, 0)
	m["AccountId"] = req.AccountId
	m["SigContent"] = req.SigContent
	m["TemplateId"] = req.TemplateId
	m["PhoneNumbers"] = req.PhoneNumbers
	m["TemplateParams"] = req.TemplateParams
	m["Action"] = req.Action
	return m
}
