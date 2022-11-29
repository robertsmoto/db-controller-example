package models

type Collection struct {
	BaseData
	Type CollectionType `json:"type" validate:"required"`
}

// CollectionDataLayer interface ..
type CollectionDataLayer interface {
	Get(Key string) (*Collection, error)
	Modify(Key string, args ...interface{}) (*Collection, error)
	Save(coll *Collection) error
	Delete(coll *Collection) error
}
