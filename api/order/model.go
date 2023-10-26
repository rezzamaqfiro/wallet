package order

type DisbursementRequest struct {
	InvoiceID string `json:"invoice_id"`
	Status    string `json:"status"`
}
