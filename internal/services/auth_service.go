package services

import (
	"errors"

	"github.com/kareemhamed001/blog/internal/helpers"
	"github.com/kareemhamed001/blog/internal/models"
	"github.com/kareemhamed001/blog/internal/types"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	DB            *gorm.DB
	JWTPrivateKey string
}

func NewAuthService(DB *gorm.DB, JWTPrivateKey string) *AuthService {
	return &AuthService{DB: DB, JWTPrivateKey: JWTPrivateKey}
}

func (s *AuthService) Register(name, email, password string, role types.UserRole) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("Error while hashing password")
	}
	user := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     role,
	}

	err = s.DB.Transaction(func(db *gorm.DB) error {

		err := db.Create(&user).Error
		if err != nil {
			return err
		}

		if role == types.RoleMerchant {
			store := models.Store{
				Name:       user.Name + "'s Store",
				MerchantID: user.ID,
			}

			err := db.Create(&store).Error
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("email already exists")
		}
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errors.New("duplicated email")
		}
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)" {
			return nil, errors.New("email already exists")
		}
		//print error type

		return nil, err
	}

	return &user, nil
}

func (s *AuthService) Login(email, password string) (*models.User, string, error) {
	var user models.User
	err := s.DB.Preload("Store").Where("email=?", email).First(&user).Error
	if err != nil {
		return nil, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", err
	}

	signedToken, err := helpers.GenerateToken(map[string]interface{}{
		"sub":     user.ID,
		"user_id": user.ID,
		"role":    user.Role,
	}, s.JWTPrivateKey)

	if err != nil {
		return nil, "", err
	}
	return &user, signedToken, nil
}
