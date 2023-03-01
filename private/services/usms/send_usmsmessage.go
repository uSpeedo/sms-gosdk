package usms

import (
	"github.com/uSpeedo/usms-sdk-go/um/request"
	"github.com/uSpeedo/usms-sdk-go/um/response"
)

// SendUSMSMessageRequest is request schema for SendUSMSMessage action
type SendUSMSMessageRequest struct {
	request.CommonBase

	// [PublicParam] Project ID. Default is 0
	AccountId *int `required:"true"`

	// Phone number array, the phone number format is (60)1xxxxxxxx
	PhoneNumbers []string `required:"true"`

	// Template variable parameters, filled in as arrays, for example，TemplateParams.0，TemplateParams.1，... If there are no variable parameters in the template, then this item can not be filled in; if there are variable parameters in the template, then this item is required
	TemplateParams []string `required:"true"`

	// Template ID
	TemplateId *string `required:"true"`
}

// SendUSMSMessageResponse is response schema for SendUSMSMessage action
type SendUSMSMessageResponse struct {
	response.CommonBase

	// Error description when error occurs
	Message string

	// The unique ID of the SMS submit for send this time, you can query the list of SMS sent this time according to this value
	SessionNo string
}

// NewSendUSMSMessageRequest will create request of SendUSMSMessage action.
func (c *USMSClient) NewSendUSMSMessageRequest() *SendUSMSMessageRequest {
	req := &SendUSMSMessageRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// SendUSMSMessage - Send SMS。
func (c *USMSClient) SendUSMSMessage(req *SendUSMSMessageRequest) (*SendUSMSMessageResponse, error) {
	var err error
	var res SendUSMSMessageResponse

	err = c.Client.InvokeAction("SendUSMSMessage", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
