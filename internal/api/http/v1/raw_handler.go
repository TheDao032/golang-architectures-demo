package v1

import (
	"net/http"

	createraw "github.com/TheDao032/golang-architectures-demo/internal/application/raw/commands/create_raw"
	"github.com/TheDao032/golang-architectures-demo/internal/service"

	// request "github.com/TheDao032/go-backend-utils-architecture/http/request"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
	v "github.com/TheDao032/go-backend-utils-architecture/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RAWHandler struct {
	service *service.Service
	logger  logger.Logger
}

func NewRAWHandler(service *service.Service, logger logger.Logger) *RAWHandler {
	return &RAWHandler{service, logger}
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
func (h *RAWHandler) CreateRAW(c *gin.Context) {
	var raws []createraw.CreateRAWCommand

	if err := c.ShouldBind(&raws); err != nil {
		var errors []string
		for _, fieldErr := range err.(validator.ValidationErrors) {
			errors = append(errors, v.GetErrorMessage(c, fieldErr, h.logger))
		}

		c.JSON(http.StatusBadRequest, errors)
		return
	}

	// gemResponse, err := h.service.GemService.GetGemSourceByUserHandler.Handle(c, gem)
	createRAWResponse, err := h.service.RAWService.CreateRAWHandler.Handle(c, raws)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, createRAWResponse)
}
