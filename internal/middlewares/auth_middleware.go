package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	appErrors "github.com/kareemhamed001/blog/internal/errors"
	"github.com/kareemhamed001/blog/internal/helpers"
	"github.com/kareemhamed001/blog/internal/responses"
)

// AuthMiddleware validates JWT token for any authenticated user and populates context claims.
func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
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

		if role, ok := claims["role"].(string); ok {
			ctx.Set("role", role)
		}
		if userID, ok := extractUserIDFromClaims(claims); ok {
			ctx.Set("user_id", userID)
		}

		ctx.Next()
	}
}
