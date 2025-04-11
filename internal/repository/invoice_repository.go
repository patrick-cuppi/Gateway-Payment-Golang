package repository

import (
	"database/sql"

	"github.com/patrick-cuppi/Gateway-Payment-Golang/internal/domain"
)

type InvoiceRepository struct {
	db *sql.DB
}

func NewInvoiceRepository(db *sql.DB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

func (r *InvoiceRepository) Save(invoice *domain.Invoice) error {
	_, err := r.db.Exec(
		"INSERT INTO invoices (id, account_id, amount, status, description, payment_type, card_last_digits, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		invoice.ID, invoice.AccountID, invoice.Amount, invoice.Status, invoice.Description, invoice.PaymentType, invoice.CardLastDigits, invoice.CreatedAt, invoice.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
