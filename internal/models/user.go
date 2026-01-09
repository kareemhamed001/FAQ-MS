package models

import (
	"time"

	"github.com/kareemhamed001/blog/internal/types"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	Email     string         `gorm:"uniqueIndex" json:"email"`
	Password  string         `json:"-"`
	Role      types.UserRole `gorm:"type:varchar(20);not null" json:"role"`
	Store     *Store         `gorm:"foreignKey:MerchantID" json:"store,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt time.Time      `gorm:"index" json:"-"`
}
