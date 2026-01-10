package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	appErrors "github.com/kareemhamed001/faq/internal/errors"
	"github.com/kareemhamed001/faq/internal/helpers"
	"github.com/kareemhamed001/faq/internal/responses"
	"github.com/kareemhamed001/faq/internal/types"
)

func HasRole(roles []types.UserRole, jwtSecret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			responses.WriteError(ctx, appErrors.ErrUnauthorized.Status, appErrors.ErrUnauthorized.Code, appErrors.ErrUnauthorized.Message)
			ctx.Abort()
			return
		}

		tokenParts := strings.SplitN(authHeader, " ", 2)
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			responses.WriteError(ctx, appErrors.ErrUnauthorized.Status, appErrors.ErrUnauthorized.Code, "invalid authorization format; expected Bearer token")
			ctx.Abort()
			return
		}

		claims, err := helpers.ValidateToken(tokenParts[1], jwtSecret)
		if err != nil {
			responses.WriteError(ctx, appErrors.ErrUnauthorized.Status, appErrors.ErrUnauthorized.Code, "invalid or expired token")
			ctx.Abort()
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			responses.WriteError(ctx, appErrors.ErrForbidden.Status, appErrors.ErrForbidden.Code, "role not found in token")
			ctx.Abort()
			return
		}

		if !containsRole(roles, types.UserRole(role)) {
			responses.WriteError(ctx, appErrors.ErrForbidden.Status, appErrors.ErrForbidden.Code, "access forbidden")
			ctx.Abort()
			return
		}

		if userID, ok := extractUserIDFromClaims(claims); ok {
			ctx.Set("user_id", userID)
		}
		ctx.Set("role", role)

		ctx.Next()
	}
}

func containsRole(roles []types.UserRole, role types.UserRole) bool {
	for _, r := range roles {
		if r == role {
			return true
		}
	}
	return false
}
