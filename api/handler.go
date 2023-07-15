package api

import (
	"asia-yo-cur-svc/config"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Config *config.Config
	Router *gin.Engine
}

func NewHandler(config *config.Config) (*Handler, error) {
	srv := &Handler{
		Config: config,
	}

	srv.setupRouter()
	return srv, nil
}

func (handler *Handler) setupRouter() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/currency", handler.GetCurrency)
	}

	handler.Router = r
}
