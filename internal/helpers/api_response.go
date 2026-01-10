package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/kareemhamed001/faq/internal/responses"
)

// WriteAPIResponse is kept for backward compatibility; prefer responses.WriteSuccess / WriteError.
func WriteAPIResponse(ctx *gin.Context, data interface{}, message string, statusCode int) map[string]interface{} {
	if statusCode >= 400 {
		responses.WriteError(ctx, statusCode, "ERROR", message)
		return nil
	}
	responses.WriteSuccess(ctx, statusCode, data, nil)
	return nil
}
