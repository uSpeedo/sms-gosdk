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
	client := um.NewClient(config, credential)
	return &USMSClient{
		client,
	}
}
