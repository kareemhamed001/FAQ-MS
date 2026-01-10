package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	dtos "github.com/kareemhamed001/faq/internal/DTOs"
	"github.com/kareemhamed001/faq/internal/helpers"
	"github.com/kareemhamed001/faq/internal/services"
)

type FAQHandler struct {
	fAQService   *services.FAQService
	storeService *services.StoreService
}

func NewFAQHandler(faqService services.FAQService, storeService services.StoreService) *FAQHandler {
	return &FAQHandler{
		fAQService:   &faqService,
		storeService: &storeService,
	}
}

func (h *FAQHandler) GetAllFAQs(ctx *gin.Context) {
	search := ctx.Query("search")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))
	sortDir := ctx.DefaultQuery("sort", "desc")

	userId, Role, err := helpers.GetUserIDAndRoleFromContext(ctx)
	if err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), 401)
		return
	}

	faqs, err := h.fAQService.GetAllFAQs(ctx.Request.Context(), search, Role, uint(userId), page, pageSize, sortDir)
	if err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), h.statusForError(err))
		return
	}

	helpers.WriteAPIResponse(ctx, gin.H{"faqs": faqs}, "FAQs retrieved successfully", 200)
}

func (h *FAQHandler) GetFAQByID(ctx *gin.Context) {
	var query struct {
		ID uint `uri:"id" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&query); err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), 400)
		return
	}

	userId, Role, err := helpers.GetUserIDAndRoleFromContext(ctx)
	if err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), 401)
		return
	}

	faq, err := h.fAQService.GetFAQByID(ctx.Request.Context(), query.ID, Role, uint(userId))
	if err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), h.statusForError(err))
		return
	}

	helpers.WriteAPIResponse(ctx, gin.H{"faq": faq}, "FAQ retrieved successfully", 200)
}

func (h *FAQHandler) CreateFAQ(ctx *gin.Context) {
	var request struct {
		CategoryID   uint                  `json:"category_id" binding:"required"`
		Translations []dtos.TranslationDTO `json:"translations" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), 400)
		return
	}

	userID, Role, err := helpers.GetUserIDAndRoleFromContext(ctx)
	if err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), 401)
		return
	}

	faq, err := h.fAQService.CreateFAQ(ctx.Request.Context(), uint(userID), request.CategoryID, request.Translations, Role)
	if err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), h.statusForError(err))
		return
	}

	helpers.WriteAPIResponse(ctx, gin.H{"faq": faq}, "FAQ created successfully", 201)
}

func (h *FAQHandler) UpdateFAQ(ctx *gin.Context) {
	var uri struct {
		ID uint `uri:"id" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), 400)
		return
	}

	var request struct {
		CategoryID   *uint                 `json:"category_id"`
		Translations []dtos.TranslationDTO `json:"translations" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), 400)
		return
	}

	userID, Role, err := helpers.GetUserIDAndRoleFromContext(ctx)
	if err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), 401)
		return
	}

	updatedFaq, err := h.fAQService.UpdateFAQ(ctx.Request.Context(), uri.ID, uint(userID), request.CategoryID, request.Translations, Role)
	if err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), h.statusForError(err))
		return
	}

	helpers.WriteAPIResponse(ctx, gin.H{"faq": updatedFaq}, "FAQ updated successfully", 200)
}

func (h *FAQHandler) DeleteFAQ(ctx *gin.Context) {
	var uri struct {
		ID uint `uri:"id" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), 400)
		return
	}

	userID, Role, err := helpers.GetUserIDAndRoleFromContext(ctx)
	if err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), 401)
		return
	}

	if err := h.fAQService.DeleteFAQ(ctx.Request.Context(), uri.ID, Role, uint(userID)); err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), h.statusForError(err))
		return
	}

	helpers.WriteAPIResponse(ctx, nil, "FAQ deleted successfully", 204)
}

func (h *FAQHandler) statusForError(err error) int {
	switch {
	case errors.Is(err, services.ErrFAQNotFound):
		return 404
	case errors.Is(err, services.ErrUnauthorizedFAQ):
		return 403
	case errors.Is(err, services.ErrCategoryNotFound), errors.Is(err, services.ErrStoreNotFound):
		return 400
	case errors.Is(err, services.ErrUnsupportedRole):
		return 403
	default:
		return 500
	}
}
