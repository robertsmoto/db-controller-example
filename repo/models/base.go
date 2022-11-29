package models

type BaseData struct {
	ID            string `json:"ID" validate:"required,uuid4"`
	SchemaVersion SchemaVersion
	CreatedAt     string
	UpdatedAt     string
}
