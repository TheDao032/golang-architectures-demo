package http

import (
	"github.com/TheDao032/golang-architectures-demo/docs"
	"github.com/TheDao032/golang-architectures-demo/pkg/healthcheck"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func MapRoutes(
	router *gin.Engine,
	healthcheckHandler *HealthcheckHandler,
	// gemHandler *v1.GemHandler
) {
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "Gem API"
	docs.SwaggerInfo.Description = "Gem API"
	docs.SwaggerInfo.BasePath = "v1"

	health := healthcheck.RunHealthCheck(
		healthcheckHandler.logger,
		healthcheckHandler.cfg,
		healthcheckHandler.readDb,
		healthcheckHandler.writeDb,
	)

	// v1 := router.Group("v1/gems")
	// {
	// 	// External
	// 	v1.GET("/source", gemHandler.GetGemSourceByUserId)
	// 	v1.GET("/dashboard", gemHandler.GetGemDashboard)
	// }

	router.Use(gin.WrapH(health))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1)))
}
