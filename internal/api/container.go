package api

import (
	"github.com/TheDao032/golang-architectures-demo/internal/api/http"
)

type ApiContainer struct {
	HttpServer *http.Server
}

func NewApiContainer(http *http.Server) *ApiContainer {
	return &ApiContainer{HttpServer: http}
}
