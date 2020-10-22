package store

import (
	"playing-with-golang-on-k8s/shared"
	"time"

	"github.com/gosimple/slug"
)

//Product represents a product in our plateform
type Product struct {
	ID          string `gorm:"primary_key"`
	Name        string
	Slug        string
	Description string
	CreatedByID string
	CreatedBy   User `gorm:"foreignkey:CreatedByID;association_foreignkey:ID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

//EnsureSlugExists sets the slug if it is empty
func (p *Product) EnsureSlugExists() {
	if p.Slug == "" {
		p.Slug = slug.Make(p.Name)
	}
}

//ProFromAPI converts api object to the store one
func ProFromAPI(pro *shared.Product) *Product {
	return &Product{
		ID:          pro.ID,
		Name:        pro.Name,
		Description: pro.Description,
		CreatedByID: pro.CreatedByID,
		CreatedAt:   pro.CreatedAt,
		DeletedAt:   pro.DeletedAt,
	}
}
