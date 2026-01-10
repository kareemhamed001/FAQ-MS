package services

import (
	"github.com/kareemhamed001/faq/internal/models"
	"gorm.io/gorm"
)

type FAQCategoryService struct {
	DB *gorm.DB
}

func NewFAQCategoryService(DB *gorm.DB) *FAQCategoryService {
	return &FAQCategoryService{DB: DB}
}

func (s *FAQCategoryService) GetAllCategories() ([]models.Category, error) {
	var faqCategories []models.Category
	err := s.DB.Find(&faqCategories).Error
	if err != nil {
		return nil, err
	}
	return faqCategories, nil
}

func (s *FAQCategoryService) GetCategoryByID(id uint) (*models.Category, error) {
	var category models.Category
	err := s.DB.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}
func (s *FAQCategoryService) CreateCategory(name string) (*models.Category, error) {
	category := models.Category{
		Name: name,
	}

	err := s.DB.Create(&category).Error
	if err != nil {
		return nil, err

	}
	return &category, nil
}

func (s *FAQCategoryService) UpdateCategory(id uint, name string) (*models.Category, error) {
	var category models.Category
	err := s.DB.First(&category, id).Error
	if err != nil {
		return nil, err
	}

	category.Name = name

	err = s.DB.Save(&category).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (s *FAQCategoryService) DeleteCategory(id uint) error {
	err := s.DB.Delete(&models.Category{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *FAQCategoryService) SearchCategories(search string) ([]models.Category, error) {
	var categories []models.Category
	err := s.DB.Where("name ILIKE ?", "%"+search+"%").Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}
