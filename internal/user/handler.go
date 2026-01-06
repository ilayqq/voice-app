package user

import (
	"net/http"
	"voice-app/dto"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler { return &Handler{service: service} }

// GetUsers godoc
//
//	@Summary		Get users
//	@Description	Get all users or filter by phone number
//	@Tags			users
//	@Param			phone_number	query		string	false	"Phone number"
//	@Success		200				{array}		dto.UserResponse
//	@Failure		500				{object}	map[string]string
//	@Router			/users [get]
//	@Security		BearerAuth
func (h *Handler) GetUsers(c *gin.Context) {
	phoneNumber := c.Query("phone_number")

	if phoneNumber != "" {
		user, err := h.service.GetByPhoneNumber(phoneNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
			return
		}

		c.JSON(http.StatusOK, user)
		return
	}

	users, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	userDTOs := make([]dto.UserResponse, len(users))
	for i, u := range users {
		roleName := ""
		if len(u.Roles) > 0 {
			roleName = u.Roles[0].Name
		}
		userDTOs[i] = dto.UserResponse{
			ID:          u.ID,
			FullName:    u.FullName,
			PhoneNumber: u.PhoneNumber,
			RoleName:    roleName,
		}
	}

	c.JSON(http.StatusOK, userDTOs)
}
