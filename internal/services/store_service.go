package services

import (
	"context"
	"errors"

	"github.com/kareemhamed001/faq/internal/models"
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

	if language == "" {
		language = "en"
	}

	query := s.DB.WithContext(ctx).
		Model(&models.FAQ{}).
		Where("store_id = ? OR is_global = ?", storeID, true).
		Preload("Category").
		Preload("Translations").
		Order("id DESC")

	if err := query.Find(&store.FAQs).Error; err != nil {
		return nil, err
	}

	// Apply translation fallback per FAQ: prefer requested language, then English, then first available
	for i := range store.FAQs {
		store.FAQs[i].Translations = filterTranslationsWithFallback(store.FAQs[i].Translations, language)
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

// filterTranslationsWithFallback returns a single translation slice honoring the requested language,
// then English, then any available translation.
func filterTranslationsWithFallback(translations []models.Translation, language string) []models.Translation {
	if len(translations) == 0 {
		return translations
	}

	for _, t := range translations {
		if t.Language == language {
			return []models.Translation{t}
		}
	}

	for _, t := range translations {
		if t.Language == "en" {
			return []models.Translation{t}
		}
	}

	return []models.Translation{translations[0]}
}
