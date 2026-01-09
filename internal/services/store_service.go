package services

import (
	"context"
	"errors"

	"github.com/kareemhamed001/blog/internal/models"
	"gorm.io/gorm"
)

type StoreService struct {
	DB *gorm.DB
}

func NewStoreService(DB *gorm.DB) *StoreService {
	return &StoreService{DB: DB}
}

func (s *StoreService) ListStores(ctx context.Context, page, pageSize int, sortDir string) ([]models.Store, error) {
	if page < 1 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}

	order := "stores.id DESC"
	if sortDir == "asc" {
		order = "stores.id ASC"
	}

	var stores []models.Store
	err := s.DB.WithContext(ctx).
		Model(&models.Store{}).
		Order(order).
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&stores).Error
	if err != nil {
		return nil, err
	}

	return stores, nil
}
func (s *StoreService) GetStoreWithFAQs(ctx context.Context, storeID uint, language string) (*models.Store, error) {

	var store models.Store
	if err := s.DB.WithContext(ctx).First(&store, storeID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrStoreNotFound
		}
		return nil, err
	}

	query := s.DB.WithContext(ctx).
		Model(&models.FAQ{}).
		Where("store_id = ? OR is_global = ?", storeID, true).
		Preload("Category").
		Order("id DESC")

	translationArgs := []interface{}{}
	if language != "" {
		translationArgs = append(translationArgs, "language = ?", language)
	}
	query = query.Preload("Translations", translationArgs...)

	if err := query.Find(&store.FAQs).Error; err != nil {
		return nil, err
	}

	return &store, nil
}

func (s *StoreService) GetStoreByID(ctx context.Context, id uint) (*models.Store, error) {
	var store models.Store
	if err := s.DB.WithContext(ctx).First(&store, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrStoreNotFound
		}
		return nil, err
	}
	return &store, nil
}

func (s *StoreService) GetStoreByMerchantID(merchantID uint) (*models.Store, error) {
	var store models.Store
	err := s.DB.Where("merchant_id = ?", merchantID).First(&store).Error
	if err != nil {
		return nil, err
	}
	return &store, nil
}
