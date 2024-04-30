// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * CDR Energy API
 *
 * Consumer Data Right end points and payloads for the Energy sector
 *
 * API version: 1.30.0
 */

package energyserver




type EnergyInvoice struct {

	// The ID of the account for which the invoice was issued
	AccountId string `json:"accountId"`

	// The number assigned to this invoice by the energy Retailer
	InvoiceNumber string `json:"invoiceNumber"`

	// The date that the invoice was actually issued (as opposed to generated or calculated)
	IssueDate string `json:"issueDate"`

	// The date that the invoice is due to be paid
	DueDate string `json:"dueDate,omitempty"`

	Period EnergyInvoicePeriod `json:"period,omitempty"`

	// The net amount due for this invoice regardless of previous balance
	InvoiceAmount string `json:"invoiceAmount,omitempty"`

	// The total GST amount for this invoice.  If absent then zero is assumed
	GstAmount string `json:"gstAmount,omitempty"`

	PayOnTimeDiscount EnergyInvoicePayOnTimeDiscount `json:"payOnTimeDiscount,omitempty"`

	// The account balance at the time the invoice was issued
	BalanceAtIssue string `json:"balanceAtIssue"`

	// Array of service point IDs to which this invoice applies. May be empty if the invoice contains no electricity usage related charges
	ServicePoints []string `json:"servicePoints"`

	Gas EnergyInvoiceGasUsageCharges `json:"gas,omitempty"`

	Electricity EnergyInvoiceElectricityUsageCharges `json:"electricity,omitempty"`

	AccountCharges EnergyInvoiceAccountCharges `json:"accountCharges,omitempty"`

	// Indicator of the payment status for the invoice
	PaymentStatus string `json:"paymentStatus"`
}

// AssertEnergyInvoiceRequired checks if the required fields are not zero-ed
func AssertEnergyInvoiceRequired(obj EnergyInvoice) error {
	elements := map[string]interface{}{
		"accountId": obj.AccountId,
		"invoiceNumber": obj.InvoiceNumber,
		"issueDate": obj.IssueDate,
		"balanceAtIssue": obj.BalanceAtIssue,
		"servicePoints": obj.ServicePoints,
		"paymentStatus": obj.PaymentStatus,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertEnergyInvoicePeriodRequired(obj.Period); err != nil {
		return err
	}
	if err := AssertEnergyInvoicePayOnTimeDiscountRequired(obj.PayOnTimeDiscount); err != nil {
		return err
	}
	if err := AssertEnergyInvoiceGasUsageChargesRequired(obj.Gas); err != nil {
		return err
	}
	if err := AssertEnergyInvoiceElectricityUsageChargesRequired(obj.Electricity); err != nil {
		return err
	}
	if err := AssertEnergyInvoiceAccountChargesRequired(obj.AccountCharges); err != nil {
		return err
	}
	return nil
}

// AssertEnergyInvoiceConstraints checks if the values respects the defined constraints
func AssertEnergyInvoiceConstraints(obj EnergyInvoice) error {
	return nil
}
