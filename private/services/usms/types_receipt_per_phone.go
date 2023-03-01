package usms

//ReceiptPerPhone - Send back information for each destination cell phone number

type ReceiptPerPhone struct {

	//  Phone Number
	Phone string

	// Number of consumed SMS
	CostCount int

	// Reply Results
	ReceiptResult string

	// Description of return results
	ReceiptDesc string

	// Return time
	ReceiptTime int
}
