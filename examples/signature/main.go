package main

import (
	"fmt"

	"github.com/usms-sdk-go/um/auth"
)

func main() {
	cred := &auth.Credential{
		AccessKeySecret: "46f09bb9fab4f12dfc160dae12273d5332b5debe",
	}
	d := "Action=DescribeUHostInstance&Limit=10&Region=cn-bj2"
	fmt.Println(cred.CreateSign(d))
}
