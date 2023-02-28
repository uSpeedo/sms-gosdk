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

	client := usms.NewClient(&cfg, &credential)

	// send request
	req := client.NewQueryUSMSTemplateRequest()
	req.AccountId = um.Int(1)
	req.Action = um.String("QueryUSMSTemplate")
	req.TemplateId = um.String("UTA2302277HBSY1")

	sign := getSignature(req)

	// add header
	req.SetNonce("hz3xevqz")
	req.SetAccessKeyId("b9c3be2916bab10f219ec24138e5c27d")
	req.SetSignature(sign)
	t, _ := time.ParseDuration("-2m")
	req.SetTimestamp(time.Now().Add(t).Unix())
	resp, err := client.QueryUSMSTemplate(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", resp)
}

func getSignature(req *usms.QueryUSMSTemplateRequest) string {
	cred := &auth.Credential{
		AccessKeySecret: "YmZmYWJiZTItZmFlNC00MWMwLTk4MzUtOWM5NjZhZjhhODJm",
	}

	m := make(map[string]interface{}, 0)
	m["AccountId"] = req.AccountId
	m["TemplateId"] = req.TemplateId
	m["Action"] = req.Action
	return cred.CreateSign(m)
}
