package usms

import (
	"github.com/uSpeedo/usms-sdk-go/um/request"
	"github.com/uSpeedo/usms-sdk-go/um/response"
)

// GetUSMSSendReceiptRequest is request schema for GetUSMSSendReceipt action
type GetUSMSSendReceiptRequest struct {
	request.CommonBase

	// [PublicParam] Project ID. Default is 0
	AccountId *int `required:"true"`

	// The set of SessionNo returned when sending SMS，SessionNoSet.0,SessionNoSet.1....
	SessionNoSet []string `required:"true"`
}

// GetUSMSSendReceiptResponse is response schema for GetUSMSSendReceipt action
type GetUSMSSendReceiptResponse struct {
	response.CommonBase

	// Error description when error occurs
	Message string

	// Collection of return information
	Data []ReceiptPerSession
}

// NewGetUSMSSendReceiptRequest will create request of GetUSMSSendReceipt action.
func (c *USMSClient) NewGetUSMSSendReceiptRequest() *GetUSMSSendReceiptRequest {
	req := &GetUSMSSendReceiptRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// GetUSMSSendReceipt - Get SMS send back information。
// There will be a delay in the return of the receipt information from the downstream service provider, so it is recommended to call the interface after 5-10 minutes to pull the receipt information after sending the SMS.
// If it takes more than 12 hours to return, please contact technical support to confirm the reason.
func (c *USMSClient) GetUSMSSendReceipt(req *GetUSMSSendReceiptRequest) (*GetUSMSSendReceiptResponse, error) {
	var err error
	var res GetUSMSSendReceiptResponse

	err = c.Client.InvokeAction("GetUSMSSendReceipt", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
