package http

import (
	"github.com/TheDao032/golang-architectures-demo/docs"
	v1 "github.com/TheDao032/golang-architectures-demo/internal/api/http/v1"
	"github.com/TheDao032/golang-architectures-demo/pkg/healthcheck"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func MapRoutes(
	router *gin.Engine,
	healthcheckHandler *HealthcheckHandler,
  acqHandler *v1.ACQHandler,
  navHandler *v1.NAVHandler,
  podHandler *v1.PODHandler,
  rawHandler *v1.RAWHandler,
  staHandler *v1.STAHandler,
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

	v1 := router.Group("v1")
	{
		// External
		v1.POST("/acq", acqHandler.CreateACQ)
		v1.POST("/nav", navHandler.CreateNAV)
		v1.POST("/pod", podHandler.CreatePOD)
		v1.POST("/raw", rawHandler.CreateRAW)
		v1.POST("/sta", staHandler.CreateSTA)
	}

	router.Use(gin.WrapH(health))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1)))
}
