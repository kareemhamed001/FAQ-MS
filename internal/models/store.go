package models

import "time"

type Store struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `json:"name"`
	MerchantID uint      `json:"merchant_id"` // FK to Users table
	CreatedAt  time.Time `json:"created_at"`
	FAQs       []FAQ     `gorm:"foreignKey:StoreID" json:"faqs,omitempty"`
}
