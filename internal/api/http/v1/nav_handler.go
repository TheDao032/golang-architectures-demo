package v1

import (
	"net/http"

	createnav "github.com/TheDao032/golang-architectures-demo/internal/application/nav/commands/create_nav"
	"github.com/TheDao032/golang-architectures-demo/internal/service"

	// request "github.com/TheDao032/go-backend-utils-architecture/http/request"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
	v "github.com/TheDao032/go-backend-utils-architecture/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type NAVHandler struct {
	service *service.Service
	logger  logger.Logger
}

func NewNAVHandler(service *service.Service, logger logger.Logger) *NAVHandler {
	return &NAVHandler{service, logger}
}

// @BasePath /v1

// CreateNAV godoc
// @Tags NAV
// @Summary Create NAV api
// @Schemes
// @Description Insert NAV to DB
// @Accept json
// @Produce json
// @Param nav body createnav.CreateNAVCommand true "NAV data"
// @Success 200 {object} createnav.GetGemSourceByUserResponse
// @Router /gems/source [get]
func (h *NAVHandler) CreateNAV(c *gin.Context) {
	var navs []createnav.CreateNAVCommand

	if err := c.ShouldBind(&navs); err != nil {
		var errors []string
		for _, fieldErr := range err.(validator.ValidationErrors) {
			errors = append(errors, v.GetErrorMessage(c, fieldErr, h.logger))
		}

		c.JSON(http.StatusBadRequest, errors)
		return
	}

	createNAVResponse, err := h.service.NAVService.CreateNAVHandler.Handle(c, navs)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, createNAVResponse)
}
