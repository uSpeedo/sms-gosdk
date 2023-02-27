package main

import (
	"fmt"

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
	req := client.NewSendUSMSMessageRequest()
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
	req.SetNonce("hz3xevqz")
	req.SetAccessKeyId("314d47318c25a38f5c24df03f6a2a255")
	req.SetSignature("314d47318c25a38f5c24df03f6a2a255")
	req.SetTimestamp(1669370992)
	resp, err := client.SendUSMSMessage(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", resp)
}
