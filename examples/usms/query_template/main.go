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
	req := client.NewQueryUSMSTemplateRequest()
	req.AccountId = um.Int(1)
	req.Action = um.String("QueryUSMSTemplate")
	req.TemplateId = []string{"UTA230***IMRPE"}

	sign := credential.CreateSign(makeSignParamMap(req))

	// add header
	req.SetNonce(utils.RandStr(10))
	req.SetAccessKeyId(credential.AccessKeyId)
	req.SetSignature(sign)
	t, _ := time.ParseDuration("-2m")
	req.SetTimestamp(time.Now().Add(t).Unix())
	resp, err := client.QueryUSMSTemplate(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", resp)
}

func makeSignParamMap(req *usms.QueryUSMSTemplateRequest) map[string]interface{} {
	m := make(map[string]interface{}, 0)
	m["AccountId"] = req.AccountId
	m["TemplateId"] = req.TemplateId
	m["Action"] = req.Action
	return m
}
