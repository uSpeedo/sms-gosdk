package usms

// ReceiptPerSession - The set of return results for each submission

type ReceiptPerSession struct {

	// SessionNo returned when sending SMS
	SessionNo string

	// Collection of SMS return messages for each cell phone number
	ReceiptSet []ReceiptPerPhone
}
