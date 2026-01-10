package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kareemhamed001/faq/internal/helpers"
	"github.com/kareemhamed001/faq/internal/services"
)

type StoreHandler struct {
	storeService *services.StoreService
}

func NewStoreHandler(storeService services.StoreService) *StoreHandler {
	return &StoreHandler{storeService: &storeService}
}

func (h *StoreHandler) ListStores(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))
	sortDir := ctx.DefaultQuery("sort", "desc")

	stores, err := h.storeService.ListStores(ctx.Request.Context(), page, pageSize, sortDir)
	if err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), 500)
		return
	}

	helpers.WriteAPIResponse(ctx, gin.H{"stores": stores}, "Stores retrieved successfully", 200)
}

func (h *StoreHandler) GetStore(ctx *gin.Context) {
	var uri struct {
		ID uint `uri:"id" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), 400)
		return
	}

	language := ctx.GetHeader("Accept-Language")
	if language == "" {
		language = "en"
	}

	storeWithFAQs, err := h.storeService.GetStoreWithFAQs(ctx.Request.Context(), uri.ID, language)
	if err != nil {
		status := 500
		if err == services.ErrStoreNotFound {
			status = 404
		}
		helpers.WriteAPIResponse(ctx, nil, err.Error(), status)
		return
	}

	helpers.WriteAPIResponse(ctx, gin.H{"store": storeWithFAQs}, "Store retrieved successfully", 200)
}
