//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/TheDao032/golang-architectures-demo/config"
	database "github.com/TheDao032/golang-architectures-demo/database"
	"github.com/TheDao032/golang-architectures-demo/internal/api"
	"github.com/TheDao032/golang-architectures-demo/internal/api/http"

	v1 "github.com/TheDao032/golang-architectures-demo/internal/api/http/v1"
	acqservice "github.com/TheDao032/golang-architectures-demo/internal/application/acq"
	navservice "github.com/TheDao032/golang-architectures-demo/internal/application/nav"
	podservice "github.com/TheDao032/golang-architectures-demo/internal/application/pod"
	rawservice "github.com/TheDao032/golang-architectures-demo/internal/application/raw"
	staservice "github.com/TheDao032/golang-architectures-demo/internal/application/sta"

	createacq "github.com/TheDao032/golang-architectures-demo/internal/application/acq/commands/create_acq"
	createnav "github.com/TheDao032/golang-architectures-demo/internal/application/nav/commands/create_nav"
	createpod "github.com/TheDao032/golang-architectures-demo/internal/application/pod/commands/create_pod"
	createraw "github.com/TheDao032/golang-architectures-demo/internal/application/raw/commands/create_raw"
	createsta "github.com/TheDao032/golang-architectures-demo/internal/application/sta/commands/create_sta"

	acqrepo "github.com/TheDao032/golang-architectures-demo/internal/infrastructure/persistent/acq"
	navrepo "github.com/TheDao032/golang-architectures-demo/internal/infrastructure/persistent/nav"
	podrepo "github.com/TheDao032/golang-architectures-demo/internal/infrastructure/persistent/pod"
	rawrepo "github.com/TheDao032/golang-architectures-demo/internal/infrastructure/persistent/raw"
	starepo "github.com/TheDao032/golang-architectures-demo/internal/infrastructure/persistent/sta"
	service "github.com/TheDao032/golang-architectures-demo/internal/service"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
	"github.com/google/wire"
)

var container = wire.NewSet(
	api.NewApiContainer,
)

var apiSet = wire.NewSet(
	http.NewServer,
)

var serviceSet = wire.NewSet(
	service.NewService,
	http.NewHealthcheckHandler,
	v1.NewACQHandler,
	v1.NewNAVHandler,
	v1.NewPODHandler,
	v1.NewRAWHandler,
	v1.NewSTAHandler,
)

var specificServiceSet = wire.NewSet(
	acqservice.NewACQService,
	navservice.NewNAVService,
	podservice.NewPODService,
	rawservice.NewRAWService,
	staservice.NewSTAService,
)

var handlerSet = wire.NewSet(
	createacq.NewCreateACQHandler,
	createnav.NewCreateNAVHandler,
	createpod.NewCreatePODHandler,
	createraw.NewCreateRAWHandler,
	createsta.NewCreateSTAHandler,
)

var repoSet = wire.NewSet(
	acqrepo.NewACQCommandRepository,
	navrepo.NewNAVCommandRepository,
	podrepo.NewPODCommandRepository,
	rawrepo.NewRAWCommandRepository,
	starepo.NewSTACommandRepository,
)

func InitializeContainer(
	appCfg *config.AppConfig,
	logger logger.Logger,
	readDB *database.ReadDB,
	writeDB *database.WriteDB,
) *api.ApiContainer {
	wire.Build(handlerSet, repoSet, specificServiceSet, serviceSet, apiSet, container)
	return &api.ApiContainer{}
}
