package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/kareemhamed001/faq/internal/helpers"
	"github.com/kareemhamed001/faq/internal/services"
)

type FAQCategoryHandler struct {
	fAQCategoryService *services.FAQCategoryService
}

func NewFAQCategoryHandler(service services.FAQCategoryService) *FAQCategoryHandler {
	return &FAQCategoryHandler{
		fAQCategoryService: &service,
	}
}

func (h *FAQCategoryHandler) GetAllCategories(ctx *gin.Context) {
	//check if search param exists
	search := ctx.Query("search")
	if search != "" {
		categories, err := h.fAQCategoryService.SearchCategories(search)
		if err != nil {
			helpers.WriteAPIResponse(ctx, nil, err.Error(), 500)
			return
		}

		helpers.WriteAPIResponse(ctx, gin.H{"categories": categories}, "Categories retrieved successfully", 200)
		return
	}

	categories, err := h.fAQCategoryService.GetAllCategories()
	if err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), 500)
		return
	}
	helpers.WriteAPIResponse(ctx, gin.H{"categories": categories}, "Categories retrieved successfully", 200)

}

func (h *FAQCategoryHandler) GetCategoryByID(ctx *gin.Context) {
	var query struct {
		ID uint `uri:"id" binding:"required"`
	}
	err := ctx.ShouldBindUri(&query)
	if err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), 400)
		return
	}

	category, err := h.fAQCategoryService.GetCategoryByID(query.ID)
	if err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), 500)
		return
	}
	helpers.WriteAPIResponse(ctx, gin.H{"category": category}, "Category retrieved successfully", 200)
}

func (h *FAQCategoryHandler) CreateCategory(ctx *gin.Context) {
	var request struct {
		Name string `json:"name" binding:"required"`
	}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), 400)
		return
	}

	category, err := h.fAQCategoryService.CreateCategory(request.Name)
	if err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), 500)
		return
	}
	helpers.WriteAPIResponse(ctx, gin.H{"category": category}, "Category created successfully", 201)
}

func (h *FAQCategoryHandler) UpdateCategory(ctx *gin.Context) {
	var uri struct {
		ID uint `uri:"id" binding:"required"`
	}
	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), 400)
		return
	}

	var request struct {
		Name string `json:"name" binding:"required"`
	}
	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), 400)
		return
	}

	category, err := h.fAQCategoryService.UpdateCategory(uri.ID, request.Name)
	if err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), 500)
		return
	}
	helpers.WriteAPIResponse(ctx, gin.H{"category": category}, "Category updated successfully", 200)
}

func (h *FAQCategoryHandler) DeleteCategory(ctx *gin.Context) {
	var uri struct {
		ID uint `uri:"id" binding:"required"`
	}
	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), 400)
		return
	}

	err = h.fAQCategoryService.DeleteCategory(uri.ID)
	if err != nil {
		helpers.WriteAPIResponse(ctx, nil, err.Error(), 500)
		return
	}
	helpers.WriteAPIResponse(ctx, nil, "Category deleted successfully", 200)
}
