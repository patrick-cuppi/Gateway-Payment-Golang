package domain

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
)

type Invoice struct {
	ID             string
	AccountID      string
	Amount         float64
	Status         Status
	Description    string
	PaymentType    string
	CardLastDigits string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CreditCard struct {
	Number         string
	CVV            string
	ExpiryMonth    int
	ExpiryYear     int
	CardholderName string
}

func NewInvoice(accountID string, amount float64, description string, paymentType string, card CreditCard) (*Invoice, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	lastDigits := card.Number[len(card.Number)-4:]

	return &Invoice{
		ID:             uuid.New().String(),
		AccountID:      accountID,
		Amount:         amount,
		Status:         StatusPending,
		Description:    description,
		PaymentType:    paymentType,
		CardLastDigits: lastDigits,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}
