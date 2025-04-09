package service

import (
	"github.com/patrick-cuppi/Gateway-Payment-Golang/internal/domain"
)

type AccountService struct {
	repository domain.AccountRepository
}

func NewAccountService(repository domain.AccountRepository) *AccountService {
	return &AccountService{repository: repository}
}
