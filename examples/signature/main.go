package main

import (
	"fmt"

	"github.com/uSpeedo/usms-sdk-go/um/auth"
)

func main() {
	cred := &auth.Credential{
		AccessKeySecret: "YmZmYWJiZTItZmFlNC00MWMwLTk4MzUtOWM5NjZhZjhhODJm",
	}

	parMap := make(map[string]interface{}, 0)
	parMap["AccountId"] = 1
	parMap["Action"] = "CreateUSMSTemplate"
	parMap["Purpose"] = 1
	parMap["TemplateName"] = "Test"
	parMap["Template"] = "测试模板"
	parMap["International"] = true

	fmt.Println(cred.CreateSign(parMap))
}
