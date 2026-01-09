package dtos

type TranslationDTO struct {
	Language string `json:"language"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
