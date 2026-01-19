package product

import (
	"net/http"
	"voice-app/domain"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

// GetAll godoc
//
//	@Summary		Get products
//	@Description	Get all products
//	@Tags			products
//	@Success		200	{array}		domain.Product
//	@Failure		401	{object}	map[string]string
//	@Failure		403	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/products [get]
//	@Security		BearerAuth
func (h *Handler) GetAll(c *gin.Context) {
	barcode := c.Query("barcode")

	if barcode != "" {
		product, err := h.service.GetByBarcode(barcode)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "product not found"})
			return
		}

		c.JSON(http.StatusOK, product)
		return
	}

	products, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// AddProduct godoc
//
//	@Summary		Add product
//	@Description	Add new product
//	@Tags			products
//	@Param			data	body		dto.ProductRequest	true	"Product data"
//	@Success		200		{array}		domain.Product
//	@Failure		401		{object}	domain.ErrorResponse
//	@Failure		403		{object}	domain.ErrorResponse
//	@Failure		500		{object}	domain.ErrorResponse
//	@Router			/products [post]
//	@Security		BearerAuth
func (h *Handler) AddProduct(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	createdProduct, err := h.service.Create(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}
