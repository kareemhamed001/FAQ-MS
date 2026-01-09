package models

type Translation struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	FAQID    uint   `json:"faq_id"`
	Language string `json:"language"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
