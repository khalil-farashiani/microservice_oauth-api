package htpp

import (
	"github.com/gin-gonic/gin"
	"github.com/khalil-farashiani/microservice_oauth-api/src/domain/access_token"
	"net/http"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "implement me!")
}
