package v1

import (
	"net/http"

	createacq "github.com/TheDao032/golang-architectures-demo/internal/application/acq/commands/create_acq"
	"github.com/TheDao032/golang-architectures-demo/internal/service"

	// request "github.com/TheDao032/go-backend-utils-architecture/http/request"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
	v "github.com/TheDao032/go-backend-utils-architecture/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ACQHandler struct {
	service *service.Service
	logger  logger.Logger
}

func NewACQHandler(service *service.Service, logger logger.Logger) *ACQHandler {
	return &ACQHandler{service, logger}
}

// @BasePath /v1

// CreateACQ godoc
// @Tags ACQ
// @Summary Create ACQ api
// @Schemes
// @Description Insert ACQ to DB
// @Accept json
// @Produce json
// @Param acq body createacq.CreateACQCommand true "ACQ data"
// @Success 200 {object} createacq.GetGemSourceByUserResponse
// @Router /gems/source [get]
func (h *ACQHandler) CreateACQ(c *gin.Context) {
	var acqs []createacq.CreateACQCommand

	if err := c.ShouldBind(&acqs); err != nil {
		var errors []string
		for _, fieldErr := range err.(validator.ValidationErrors) {
			errors = append(errors, v.GetErrorMessage(c, fieldErr, h.logger))
		}

		c.JSON(http.StatusBadRequest, errors)
		return
	}

	// gemResponse, err := h.service.GemService.GetGemSourceByUserHandler.Handle(c, gem)
	createACQResponse, err := h.service.ACQService.CreateACQHandler.Handle(c, acqs)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, createACQResponse)
}
