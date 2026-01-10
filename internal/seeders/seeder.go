package seeders

import (
	"errors"
	"fmt"
	"log"

	"github.com/kareemhamed001/faq/internal/models"
	"github.com/kareemhamed001/faq/internal/types"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Seeder struct {
	DB *gorm.DB
}

func NewSeeder(db *gorm.DB) *Seeder {
	return &Seeder{DB: db}
}

// SeedAdminUser seeds an admin user to the database
func (s *Seeder) SeedAdminUser(name, email, password string) (*models.User, error) {
	// Check if admin user already exists
	var existingUser models.User
	if err := s.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return nil, fmt.Errorf("admin user with email %s already exists", email)
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("error checking for existing user: %w", err)
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	// Create admin user
	adminUser := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     types.RoleAdmin,
	}

	if err := s.DB.Create(&adminUser).Error; err != nil {
		return nil, fmt.Errorf("error creating admin user: %w", err)
	}

	log.Printf("✓ Admin user created successfully with email: %s\n", email)
	return &adminUser, nil
}

// SeedDefaultAdminUser seeds a default admin user (admin@example.com)
func (s *Seeder) SeedDefaultAdminUser() (*models.User, error) {
	return s.SeedAdminUser("Admin", "admin@example.com", "admin123")
}

// SeedAllDefaults seeds all default data
func (s *Seeder) SeedAllDefaults() error {
	_, err := s.SeedDefaultAdminUser()
	if err != nil {
		log.Printf("Warning: %v\n", err)
		// Don't return error, continue seeding
	}
	return nil
}
