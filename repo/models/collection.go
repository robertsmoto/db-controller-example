package models

type Collection struct {
	BaseData
	Type CollectionType `json:"type" validate:"required"`
}

// CollectionDataLayer interface ..
type CollectionDataLayer interface {
	// put function signature here
}
