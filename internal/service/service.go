package service

import (
	acqservice "github.com/TheDao032/golang-architectures-demo/internal/application/acq"
	navservice "github.com/TheDao032/golang-architectures-demo/internal/application/nav"
	podservice "github.com/TheDao032/golang-architectures-demo/internal/application/pod"
	rawservice "github.com/TheDao032/golang-architectures-demo/internal/application/raw"
	staservice "github.com/TheDao032/golang-architectures-demo/internal/application/sta"
)

type Service struct {
	ACQService *acqservice.ACQService
	NAVService *navservice.NAVService
	PODService *podservice.PODService
	RAWService *rawservice.RAWService
	STAService *staservice.STAService
}

func NewService(
	acqService *acqservice.ACQService,
	navService *navservice.NAVService,
	podService *podservice.PODService,
	rawService *rawservice.RAWService,
	staService *staservice.STAService,
) *Service {
	return &Service{
		ACQService: acqService,
		NAVService: navService,
		PODService: podService,
		RAWService: rawService,
		STAService: staService,
	}
}
