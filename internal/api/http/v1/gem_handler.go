package v1

// type GemHandler struct {
// 	service *service.Service
// 	logger  logger.Logger
// }

// func NewGemHandler(service *service.Service, logger logger.Logger) *GemHandler {
// 	return &GemHandler{service, logger}
// }

// // @BasePath /v1

// // GetGem godoc
// // @Tags Gems
// // @Summary Get gem source
// // @Schemes
// // @Description Get gem source by userId
// // @Accept json
// // @Produce json
// // @Param gem body getgemsourcebyuser.GetGemSourceByUserQuery true "Gem Source data"
// // @Success 200 {object} getgemsourcebyuser.GetGemSourceByUserResponse
// // @Router /gems/source [get]
// func (h *GemHandler) GetGemSourceByUserId(c *gin.Context) {
// 	userContext := request.GetUserContext(c)

// 	filter, _ := c.GetQuery("filter")
// 	gem := &getgemsourcebyuser.GetGemSourceByUserQuery{
// 		UserId: userContext.UserId,
// 		Filter: filter,
// 	}

// 	// if err := c.ShouldBind(&gem); err != nil {
// 	// 	var errors []string
// 	// 	for _, fieldErr := range err.(validator.ValidationErrors) {
// 	// 		errors = append(errors, v.GetErrorMessage(c, fieldErr, h.logger))
// 	// 	}

// 	// 	c.JSON(http.StatusBadRequest, errors)
// 	// 	return
// 	// }

// 	gemResponse, err := h.service.GemService.GetGemSourceByUserHandler.Handle(c, gem)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(200, gemResponse)
// }

// // @BasePath /v1

// // GetGem godoc
// // @Tags Gems
// // @Summary Get gem dashboard
// // @Schemes
// // @Description Get gem dashboard by userId
// // @Accept json
// // @Produce json
// // @Param gem body getgemdashboard.GetGemDashboardQuery true "Gem dashboard data"
// // @Success 200 {object} getgemdashboard.GetGemDashboardResponse
// // @Router /gems/dashboard [get]
// func (h *GemHandler) GetGemDashboard(c *gin.Context) {
// 	userContext := request.GetUserContext(c)
// 	gem := &getgemdashboard.GetGemDashboardQuery{
// 		UserId: userContext.UserId,
// 	}

// 	if err := c.ShouldBind(&gem); err != nil {
// 		var errors []string
// 		for _, fieldErr := range err.(validator.ValidationErrors) {
// 			errors = append(errors, v.GetErrorMessage(c, fieldErr, h.logger))
// 		}

// 		c.JSON(http.StatusBadRequest, errors)
// 		return
// 	}

// 	gemResponse, err := h.service.GemService.GetGemDashboardHandler.Handle(c, gem)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(200, gemResponse)
// }
