package usms

/*
ReceiptPerPhone - Send back information for each destination cell phone number
*/
type ReceiptPerPhone struct {

	// Number of consumed SMS
	CostCount int

	// Mobile Number
	Phone string

	// Status Report Codes
	ReceiptCode string

	// Description of return results
	ReceiptDesc string

	// Results
	ReceiptResult string

	// Return receipt return time
	ReceiptTime int

	// Customized business identification ID
	UserId string
}

/*
ReceiptPerSession - The set of return results for each submission
*/
type ReceiptPerSession struct {

	// Collection of SMS return messages for each cell phone number
	ReceiptSet []ReceiptPerPhone

	// SessionNo return when send SMS
	SessionNo string
}

/*
OutTemplate - SMS Template
*/
type OutTemplate struct {

	// Creation time
	CreateTime int

	// Reasons for audit failure
	ErrDesc string

	// Template variable property description
	Instruction string

	// Template type, options: 1-Captcha class 2-Notification class 3-Membership promotion class
	Purpose int

	// Template Description
	Remark string

	// SMS template status; status description: 0-pending review, 1-under review, 2-cleared, 3-uncleared, 4-disabled
	Status int

	// Text message template content
	Template string

	// SMS Template ID
	TemplateId string

	// SMS Template Name
	TemplateName string
}

/*
FailPhoneDetail - The numbers that were not sent successfully in the batch task and their reasons
*/
type FailPhoneDetail struct {

	// Extension Number
	ExtendCode string

	// Send failure reason。
	FailureDetails string

	// Mobile
	Phone string

	// Template Parameters
	TemplateParams []string

	// User-defined ID
	UserId string
}

/*
BatchInfo - Details of messages that were not successfully sent in the batch sending
*/
type BatchInfo struct {

	// Unsuccessful delivery details
	FailureDetails string

	// Specific number information
	Target []FailPhoneDetail

	// Template ID
	TemplateId string
}
