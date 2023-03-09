package main

import (
	"encoding/json"
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
	req.AccountId = um.Int(0)
	req.Action = um.String("SendBatchUSMSMessage")
	req.TaskContent = []usms.SendBatchInfo{
		{
			TemplateId: "...",
			SenderId:   "",
			Target: []usms.SendBatchTarget{
				{
					Phone: "130xxxx1321",
				},
				{
					Phone: "130xxxx1321",
				},
			},
		},
	}
	//add header
	req.SetNonce(utils.RandStr(10))
	req.SetAccessKeyId(credential.AccessKeyId)
	req.SetSignature(credential.CreateSign(JSONMethod(req)))
	t, _ := time.ParseDuration("-2m")
	req.SetTimestamp(time.Now().Add(t).Unix())
	resp, err := client.SendBatchUSMSMessage(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", resp)
}

func JSONMethod(content interface{}) map[string]interface{} {
	data, _ := json.Marshal(&content)
	m := make(map[string]interface{})
	_ = json.Unmarshal(data, &m)
	return m
}
