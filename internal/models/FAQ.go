package models

type FAQ struct {
	ID           uint          `gorm:"primaryKey" json:"id"`
	CategoryID   uint          `json:"category_id"`
	Category     Category      `json:"category"`
	StoreID      *uint         `json:"store_id"` // Nullable if its global
	IsGlobal     bool          `json:"is_global"`
	Translations []Translation `json:"translations"`
	Store        *Store        `json:"store,omitempty"`
}
