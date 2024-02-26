package service

import (
	acqservice "github.com/TheDao032/golang-architectures-demo/internal/application/acq"
)

type Service struct {
	ACQService *acqservice.ACQService
}

func NewService(acqService *acqservice.ACQService) *Service {
	return &Service{ACQService: acqService}
}
