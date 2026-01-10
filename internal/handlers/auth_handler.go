package handlers

import (
	"regexp"

	"github.com/gin-gonic/gin"
	appErrors "github.com/kareemhamed001/faq/internal/errors"
	"github.com/kareemhamed001/faq/internal/requests"
	"github.com/kareemhamed001/faq/internal/responses"
	"github.com/kareemhamed001/faq/internal/services"
	"github.com/kareemhamed001/faq/internal/types"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(service services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: &service,
	}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var request requests.RegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		responses.WriteError(ctx, appErrors.ErrInvalidInput.Status, appErrors.ErrInvalidInput.Code, err.Error())
		return
	}
	//validate email , password and role
	if !validateName(request.Name) {
		ctx.JSON(400, gin.H{"error": "Name is required"})
		return
	}
	if !validateEmail(request.Email) {
		ctx.JSON(400, gin.H{"error": "Invalid email format"})
		return
	}
	if !validatePassword(request.Password) {
		ctx.JSON(400, gin.H{"error": "Password must be at least 8 characters long and contain at least one number and one special character"})
		return
	}
	if !validateRole(types.UserRole(request.Role)) {
		ctx.JSON(400, gin.H{"error": "Invalid role"})
		return
	}

	user, err := h.authService.Register(request.Name, request.Email, request.Password, types.UserRole(request.Role))
	if err != nil {
		responses.WriteError(ctx, 400, "REGISTER_FAILED", err.Error())
		return
	}
	responses.WriteSuccess(ctx, 201, gin.H{"user": user}, nil)
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var request requests.LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		responses.WriteError(ctx, appErrors.ErrInvalidInput.Status, appErrors.ErrInvalidInput.Code, err.Error())
		return
	}
	if !validateEmail(request.Email) {
		ctx.JSON(400, gin.H{"error": "Invalid email format"})
		return
	}

	user, token, err := h.authService.Login(request.Email, request.Password)
	if err != nil {
		responses.WriteError(ctx, 401, "LOGIN_FAILED", err.Error())
		return
	}
	responses.WriteSuccess(ctx, 200, gin.H{"user": user, "token": token}, nil)
}

func validateName(name string) bool {
	return len(name) > 0
}

func validateEmail(email string) bool {
	if len(email) == 0 {
		return false
	}
	// Simple regex for email validation
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func validatePassword(password string) bool {
	// Password must be at least 8 characters, contain at least one number and one special character
	if len(password) < 8 {
		return false
	}

	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*]`).MatchString(password)

	return hasNumber && hasSpecial
}

func validateRole(role types.UserRole) bool {
	if len(role) == 0 {
		return false
	}
	if role != types.RoleMerchant && role != types.RoleCustomer {
		return false
	}
	// Add role validation logic here (e.g., check if role is one of the allowed roles)
	return true
}
