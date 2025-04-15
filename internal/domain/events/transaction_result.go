package events

import "github.com/patrick-cuppi/Gateway-Payment-Golang/internal/domain"

type TransactionResult struct {
	InvoiceID string `json:"invoice_id"`
	Status    string `json:"status"`
}

func NewTransactionResult(invoiceID string, status string) *TransactionResult {
	return &TransactionResult{
		InvoiceID: invoiceID,
		Status:    status,
	}
}

func (t *TransactionResult) ToDomainStatus() domain.Status {
	return domain.Status(t.Status)
}
