package v1

import (
	"net/http"

	createpod "github.com/TheDao032/golang-architectures-demo/internal/application/pod/commands/create_pod"
	"github.com/TheDao032/golang-architectures-demo/internal/service"

	// request "github.com/TheDao032/go-backend-utils-architecture/http/request"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
	v "github.com/TheDao032/go-backend-utils-architecture/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PODHandler struct {
	service *service.Service
	logger  logger.Logger
}

func NewPODHandler(service *service.Service, logger logger.Logger) *PODHandler {
	return &PODHandler{service, logger}
}

// @BasePath /v1

// CreatePOD godoc
// @Tags POD
// @Summary Create NAV api
// @Schemes
// @Description Insert NAV to DB
// @Accept json
// @Produce json
// @Param nav body createnav.CreateNAVCommand true "NAV data"
// @Success 200 {object} createnav.GetGemSourceByUserResponse
// @Router /gems/source [get]
func (h *PODHandler) CreatePOD(c *gin.Context) {
	var pods []createpod.CreatePODCommand

	if err := c.ShouldBind(&pods); err != nil {
		var errors []string
		for _, fieldErr := range err.(validator.ValidationErrors) {
			errors = append(errors, v.GetErrorMessage(c, fieldErr, h.logger))
		}

		c.JSON(http.StatusBadRequest, errors)
		return
	}

	createPODResponse, err := h.service.PODService.CreatePODHandler.Handle(c, pods)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, createPODResponse)
}
