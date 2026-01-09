package services

import (
	"context"
	"errors"

	dtos "github.com/kareemhamed001/blog/internal/DTOs"
	"github.com/kareemhamed001/blog/internal/models"
	"github.com/kareemhamed001/blog/internal/types"
	"gorm.io/gorm"
)

var (
	ErrFAQNotFound      = errors.New("faq not found")
	ErrCategoryNotFound = errors.New("category not found")
	ErrStoreNotFound    = errors.New("store not found for merchant")
	ErrUnauthorizedFAQ  = errors.New("unauthorized to access faq")
	ErrUnsupportedRole  = errors.New("role not permitted for this action")
)

type FAQService struct {
	DB *gorm.DB
}

func NewFAQService(DB *gorm.DB) *FAQService {
	return &FAQService{DB: DB}
}

func (s *FAQService) GetAllFAQs(ctx context.Context, search string, role types.UserRole, userId uint, page, pageSize int, sortDir string) ([]models.FAQ, error) {
	faqQuery := s.DB.WithContext(ctx).
		Model(&models.FAQ{}).
		Preload("Translations").
		Preload("Category")

	if search != "" {
		faqQuery = faqQuery.Joins("JOIN translations ON translations.faq_id = faqs.id").
			Where("translations.question ILIKE ? OR translations.answer ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	switch role {
	case types.RoleAdmin:
		// Admin sees everything
	case types.RoleMerchant:
		faqQuery = faqQuery.Where("faqs.is_global = ? OR faqs.store_id IN (SELECT id FROM stores WHERE merchant_id = ?)", true, userId)
	case types.RoleCustomer:
		faqQuery = faqQuery.Where("faqs.is_global = ?", true)
	default:
		return nil, ErrUnsupportedRole
	}

	if page < 1 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}

	order := "faqs.id DESC"
	if sortDir == "asc" {
		order = "faqs.id ASC"
	}

	var faqs []models.FAQ
	err := faqQuery.Distinct("faqs.*").
		Order(order).
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&faqs).Error
	if err != nil {
		return nil, err
	}
	return faqs, nil
}

func (s *FAQService) GetFAQByID(ctx context.Context, id uint, role types.UserRole, userId uint) (*models.FAQ, error) {
	faq, err := s.loadFAQ(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := s.ensureCanViewFAQ(ctx, role, userId, faq); err != nil {
		return nil, err
	}

	return faq, nil
}

func (s *FAQService) CreateFAQ(ctx context.Context, userId, categoryId uint, translations []dtos.TranslationDTO, role types.UserRole) (*models.FAQ, error) {
	var createdFAQID uint

	err := s.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.assertCategoryExists(tx, categoryId); err != nil {
			return err
		}

		faq := models.FAQ{
			CategoryID: categoryId,
			IsGlobal:   false,
		}

		switch role {
		case types.RoleAdmin:
			faq.IsGlobal = true
		case types.RoleMerchant:
			storeID, err := s.getMerchantStoreID(tx, userId)
			if err != nil {
				return err
			}
			faq.StoreID = &storeID
		default:
			return ErrUnsupportedRole
		}

		for _, t := range translations {
			faq.Translations = append(faq.Translations, models.Translation{
				Language: t.Language,
				Question: t.Question,
				Answer:   t.Answer,
			})
		}

		if err := tx.Create(&faq).Error; err != nil {
			return err
		}

		createdFAQID = faq.ID
		return nil
	})

	if err != nil {
		return nil, err
	}

	return s.loadFAQ(ctx, createdFAQID)
}

func (s *FAQService) UpdateFAQ(ctx context.Context, id uint, userId uint, categoryId *uint, translations []dtos.TranslationDTO, role types.UserRole) (*models.FAQ, error) {
	err := s.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		faq := models.FAQ{}
		if err := tx.Preload("Translations").First(&faq, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrFAQNotFound
			}
			return err
		}

		if err := s.ensureCanManageFAQ(tx, role, userId, &faq); err != nil {
			return err
		}

		if categoryId != nil {
			if err := s.assertCategoryExists(tx, *categoryId); err != nil {
				return err
			}
			faq.CategoryID = *categoryId
			if err := tx.Model(&faq).Update("category_id", faq.CategoryID).Error; err != nil {
				return err
			}
		}

		existing := make(map[string]models.Translation)
		for _, tr := range faq.Translations {
			existing[tr.Language] = tr
		}

		seen := make(map[string]bool)
		for _, t := range translations {
			seen[t.Language] = true
			if current, ok := existing[t.Language]; ok {
				if err := tx.Model(&models.Translation{}).
					Where("id = ?", current.ID).
					Updates(map[string]interface{}{
						"question": t.Question,
						"answer":   t.Answer,
					}).Error; err != nil {
					return err
				}
			} else {
				newTranslation := models.Translation{
					FAQID:    faq.ID,
					Language: t.Language,
					Question: t.Question,
					Answer:   t.Answer,
				}
				if err := tx.Create(&newTranslation).Error; err != nil {
					return err
				}
			}
		}

		for lang, current := range existing {
			if !seen[lang] {
				if err := tx.Delete(&models.Translation{}, current.ID).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return s.loadFAQ(ctx, id)
}

func (s *FAQService) DeleteFAQ(ctx context.Context, id uint, role types.UserRole, userId uint) error {
	return s.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		faq := models.FAQ{}
		if err := tx.First(&faq, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrFAQNotFound
			}
			return err
		}

		if err := s.ensureCanManageFAQ(tx, role, userId, &faq); err != nil {
			return err
		}

		if err := tx.Where("faq_id = ?", id).Delete(&models.Translation{}).Error; err != nil {
			return err
		}

		return tx.Delete(&faq).Error
	})
}

func (s *FAQService) AddTranslation(faqId uint, language, question, answer string) (*models.Translation, error) {
	translation := models.Translation{
		FAQID:    faqId,
		Language: language,
		Question: question,
		Answer:   answer,
	}

	err := s.DB.Create(&translation).Error
	if err != nil {
		return nil, err
	}
	return &translation, nil
}

func (s *FAQService) UpdateTranslation(id uint, language, question, answer string) (*models.Translation, error) {
	var translation models.Translation
	err := s.DB.First(&translation, id).Error
	if err != nil {
		return nil, err
	}

	translation.Language = language
	translation.Question = question
	translation.Answer = answer

	err = s.DB.Save(&translation).Error
	if err != nil {
		return nil, err
	}

	return &translation, nil
}

func (s *FAQService) DeleteTranslation(id uint) error {
	err := s.DB.Delete(&models.Translation{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *FAQService) loadFAQ(ctx context.Context, id uint) (*models.FAQ, error) {
	faq := models.FAQ{}
	err := s.DB.WithContext(ctx).
		Preload("Translations").
		Preload("Category").
		First(&faq, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrFAQNotFound
	}

	if err != nil {
		return nil, err
	}

	return &faq, nil
}

func (s *FAQService) ensureCanViewFAQ(ctx context.Context, role types.UserRole, userId uint, faq *models.FAQ) error {
	switch role {
	case types.RoleAdmin:
		return nil
	case types.RoleMerchant:
		if faq.IsGlobal {
			return nil
		}
		if faq.StoreID == nil {
			return ErrUnauthorizedFAQ
		}
		storeID, err := s.getMerchantStoreID(s.DB.WithContext(ctx), userId)
		if err != nil {
			return err
		}
		if *faq.StoreID != storeID {
			return ErrUnauthorizedFAQ
		}
		return nil
	case types.RoleCustomer:
		if faq.IsGlobal {
			return nil
		}
		return ErrUnauthorizedFAQ
	default:
		return ErrUnsupportedRole
	}
}

func (s *FAQService) ensureCanManageFAQ(db *gorm.DB, role types.UserRole, userId uint, faq *models.FAQ) error {
	switch role {
	case types.RoleAdmin:
		return nil
	case types.RoleMerchant:
		if faq.IsGlobal || faq.StoreID == nil {
			return ErrUnauthorizedFAQ
		}
		storeID, err := s.getMerchantStoreID(db, userId)
		if err != nil {
			return err
		}
		if *faq.StoreID != storeID {
			return ErrUnauthorizedFAQ
		}
		return nil
	default:
		return ErrUnauthorizedFAQ
	}
}

func (s *FAQService) assertCategoryExists(db *gorm.DB, categoryId uint) error {
	var category models.Category
	err := db.Select("id").First(&category, categoryId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrCategoryNotFound
	}
	if err != nil {
		return err
	}
	return nil
}

func (s *FAQService) getMerchantStoreID(db *gorm.DB, merchantID uint) (uint, error) {
	var store models.Store
	err := db.Select("id").Where("merchant_id = ?", merchantID).First(&store).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, ErrStoreNotFound
	}
	if err != nil {
		return 0, err
	}
	return store.ID, nil
}
