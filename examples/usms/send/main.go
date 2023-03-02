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
	req := client.NewSendBatchUSMSMessageRequest()
	req.AccountId = um.Int(1)
	req.Action = um.String("SendBatchUSMSMessage")
	req.Target = &usms.SendBatchInfo{
		TemplateId: "...",
		Targets: []usms.SendBatchTarget{
			{TemplateParams: []string{"1311"}, Phone: "138xxxx1123"},
		},
	}
	//add header
	req.SetNonce(utils.RandStr(10))
	req.SetAccessKeyId(credential.AccessKeyId)
	req.SetSignature(credential.CreateSign(makeSendParamMap(req)))
	t, _ := time.ParseDuration("-2m")
	req.SetTimestamp(time.Now().Add(t).Unix())
	resp, err := client.SendBatchUSMSMessage(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", resp)
}

func makeSendParamMap(req *usms.SendBatchUSMSMessageRequest) map[string]interface{} {
	m := make(map[string]interface{}, 0)
	m["AccountId"] = req.AccountId
	m["Target"] = req.Target
	m["Action"] = req.Action
	return m
}
