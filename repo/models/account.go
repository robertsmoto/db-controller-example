package models

type Account struct {
	ID        string `json:"ID" validate:"required,uuid4"`
	Name      string `json:"name" validate:"omitempty"`
	Auth      string `json:"auth" validate:"omitempty,uuid4"`
	Prefix    string `json:"prefix" validate:"omitempty"`
	IsCurrent bool   `json:"isCurrent" validate:"bool"`
	CreatedAt string `json:"createdAt"`
}
