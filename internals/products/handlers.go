package products

import "github.com/gin-gonic/gin"

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(c *gin.Context) {
//call the service layer to get the list of products and return the responcemto the client
//Return the list of products as json responce
products,err := h.service.ListProducts(c.Request.Context())
if err != nil {
	c.JSON(500,gin.H{
		"error" : "Failed to get products",
	})
	return
}
c.JSON(200,gin.H{
	"products" : products,
})
}
