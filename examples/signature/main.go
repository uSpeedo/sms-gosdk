package main

import (
	"fmt"

	"github.com/uSpeedo/usms-sdk-go/um/auth"
)

func main() {
	cred := &auth.Credential{
		AccessKeySecret: "...",
	}

	parMap := make(map[string]interface{}, 0)
	parMap["AccountId"] = 1
	parMap["Action"] = "CreateUSMSTemplate"
	parMap["Purpose"] = 1
	parMap["TemplateName"] = "test template"
	parMap["Template"] = "test notice"

	fmt.Println(cred.CreateSign(parMap))
}
