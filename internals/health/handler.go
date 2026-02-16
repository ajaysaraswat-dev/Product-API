package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CheckHealth(c *gin.Context) {
	result := h.service.CheckHealth(c.Request.Context())
	if result["status"]!="up"{
		c.JSON(http.StatusServiceUnavailable,result)
		return
	}
	c.JSON(http.StatusOK,result)
}