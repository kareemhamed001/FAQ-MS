package helpers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kareemhamed001/faq/internal/types"
)

func GetUserIDAndRoleFromContext(ctx *gin.Context) (uint64, types.UserRole, error) {
	userID, exists := ctx.Get("user_id")
	if !exists || userID == nil {
		return 0, "", errors.New("user_id not found in context")
	}

	role, exists := ctx.Get("role")
	if !exists || role == nil {
		return 0, "", errors.New("role not found in context")
	}

	// Handle different types for user_id (could be float64 from JWT claims)
	var userIDUint uint64
	switch v := userID.(type) {
	case string:
		userIDUint, _ = strconv.ParseUint(v, 10, 64)
	case float64:
		userIDUint = uint64(v)
	case int:
		userIDUint = uint64(v)
	case int64:
		userIDUint = uint64(v)
	case uint:
		userIDUint = uint64(v)
	case uint64:
		userIDUint = v
	default:
		return 0, "", errors.New("invalid user_id type in context")
	}

	roleStr, ok := role.(string)
	if !ok {
		return userIDUint, "", errors.New("invalid role type in context")
	}

	return userIDUint, types.UserRole(roleStr), nil
}
