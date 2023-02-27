// Code is generated by um-model, DO NOT EDIT IT.

package usms

import (
	"github.com/uSpeedo/usms-sdk-go/um"
	"github.com/uSpeedo/usms-sdk-go/um/auth"
)

// USMSClient is the client of USMS
type USMSClient struct {
	*um.Client
}

// NewClient will return a instance of USMSClient
func NewClient(config *um.Config, credential *auth.Credential) *USMSClient {
	meta := um.ClientMeta{Product: "USMS"}
	client := um.NewClientWithMeta(config, credential, meta)
	return &USMSClient{
		client,
	}
}
