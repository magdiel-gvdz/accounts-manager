package handlers

import (
	"net/http"
	"strconv"

	"github.com/Magdiel-GVdz/accounts-manager/internal/application"
	"github.com/Magdiel-GVdz/accounts-manager/internal/domain"
	"github.com/Magdiel-GVdz/accounts-manager/middlewares"
	"github.com/Magdiel-GVdz/accounts-manager/utils"
	"github.com/gin-gonic/gin"
)

// UserHandler exposes HTTP endpoints for user operations.
type UserHandler struct {
	service *application.UserService
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(s *application.UserService) *UserHandler { return &UserHandler{service: s} }

// RegisterRoutes registers user routes on the given router.
func (h *UserHandler) RegisterRoutes(router *gin.Engine) {
	router.POST("/users", h.CreateUser)
	router.POST("/login", h.LoginUser)

	auth := router.Group("/users")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.GET("/:id", h.GetUser)
		auth.PUT("/:id", h.UpdateUser)
		auth.DELETE("/:id", h.DeleteUser)
		auth.GET("/", h.ListUsers)
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.service.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	user, err := h.service.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	user.ID = uint(id)
	if err := h.service.Update(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.service.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var req domain.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.service.Authenticate(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	accessToken, _ := utils.GenerateJWT(user.Username)
	refreshToken, _ := utils.GenerateRefreshToken(user.Username)

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken, "refresh_token": refreshToken})
}
