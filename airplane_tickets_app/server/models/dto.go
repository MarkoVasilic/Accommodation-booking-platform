package models

type OnlyEmail struct {
	Email *string `json:"email"      validate:"email,required"`
}
