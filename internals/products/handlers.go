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
products,err :=h.service.ListProducts(c.Request.Context())
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

func (h *handler) GetProduct(c * gin.Context){
	id := c.Param("id")
	product,err := h.service.GetProduct(c.Request.Context(),id)
	if err != nil {
		c.JSON(404,gin.H{
			"error" : "Failed to get product",
		})
		return
	}
	c.JSON(200,gin.H{
		"product":product,
	})
}

func (h *handler) CreateProduct(c *gin.Context){
	var product Product
	if err := c.ShouldBindJSON(&product);err != nil {
		c.JSON(400,gin.H{
			"error" : "Invalid request body",
			"details" : err.Error(),
		})
	}
	err := h.service.CreateProduct(c.Request.Context(),&product)
	if err!= nil {
		c.JSON(500,gin.H{
			"error" : "Failed to create product",
		})
		return
	}
	c.JSON(201,gin.H{
		"message" : "product created successfully",
	})
}

func(h *handler) UpdateProduct(c *gin.Context){
	id := c.Param("id")
	var product Product 
	if err := c.ShouldBindJSON(&product);err!=nil {
		c.JSON(400,gin.H{
			"error" : "Invalid request body",
		})
	}
	err := h.service.UpdateProduct(c.Request.Context(),id,&Product{})
	if err != nil {
		c.JSON(500,gin.H{
			"error" : "Failed to update product",
		})
		return
	}
	c.JSON(200,gin.H{
		"message" : "product updated successfully",
	})
}

func (h *handler) DeleteProduct(c *gin.Context){
	id := c.Param("id")
	err := h.service.DeleteProduct(c.Request.Context(),id)
	if err != nil {
		c.JSON(500,gin.H{
			"error" : "Failed to delete product",
		})
		return
	}
	c.JSON(200,gin.H{
		"message" : "product deleted successfully",
	})
}
