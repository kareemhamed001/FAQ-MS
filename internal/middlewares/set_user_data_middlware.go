package middlewares

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kareemhamed001/faq/internal/helpers"
)

func SetUserData(jwtSecret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.Next()
			return
		}

		tokenParts := strings.SplitN(authHeader, " ", 2)
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			ctx.Next()
			return
		}

		claims, err := helpers.ValidateToken(tokenParts[1], jwtSecret)
		if err != nil {
			ctx.Next()
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			ctx.Next()
			return
		}

		if userID, ok := extractUserIDFromClaims(claims); ok {
			ctx.Set("user_id", userID)
		}
		ctx.Set("role", role)

		ctx.Next()
	}
}

func extractUserIDFromClaims(claims map[string]interface{}) (uint64, bool) {
	if raw, ok := claims["user_id"]; ok {
		if id, ok := convertToUint64(raw); ok {
			return id, true
		}
	}

	if userMap, ok := claims["user"].(map[string]interface{}); ok {
		if raw, ok := userMap["id"]; ok {
			if id, ok := convertToUint64(raw); ok {
				return id, true
			}
		}
	}

	return 0, false
}

func convertToUint64(v interface{}) (uint64, bool) {
	switch val := v.(type) {
	case uint64:
		return val, true
	case uint:
		return uint64(val), true
	case int:
		return uint64(val), true
	case int64:
		return uint64(val), true
	case float64:
		return uint64(val), true
	case string:
		parsed, err := strconv.ParseUint(val, 10, 64)
		if err == nil {
			return parsed, true
		}
	}
	return 0, false
}
