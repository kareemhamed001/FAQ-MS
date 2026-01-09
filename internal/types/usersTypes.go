package types

type UserRole string

const (
	RoleAdmin    UserRole = "admin"
	RoleMerchant UserRole = "merchant"
	RoleCustomer UserRole = "customer"
)

type TranslationInput struct {
	Language string `json:"language" binding:"required"`
	Question string `json:"question" binding:"required"`
	Answer   string `json:"answer" binding:"required"`
}
