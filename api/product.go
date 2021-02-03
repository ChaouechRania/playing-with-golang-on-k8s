package api

import (
	"playing-with-golang-on-k8s/shared"
	"time"
)

//Org update masks
const (
	ProNameMask        = "pro.name"
	ProDescriptionMask = "pro.description"
	ProPriceMask       = "pro.price"
	ProQuantityMask    = "pro.quantity"
)

//Product is used for product creation
type Product struct {
	ID          string     `json:"id"`
	Name        string     `json:"name" valid:"required~The name is required"`
	Description string     `json:"description"`
	Price       float32    `json:"price"`
	Quantity    int        `json:"quantity"`
	CreatedAt   time.Time  `json:"createdAt,omitempty"`
	UpdatedAt   time.Time  `json:"updatedAt,omitempty"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty"`
}

//OrgUpsertRequest is used for job creation
type ProUpsertRequest struct {
	Pro        *shared.Product `json:"product,omitempty" valid:"required"`
	UpdateMask []string        `json:"updateMask"`
}

//WithID set job ID
func (req *ProUpsertRequest) WithID(id string) {
	req.Pro.ID = id
}

//WithCreatedBy set created by ID
func (req *ProUpsertRequest) WithCreatedBy(id string) {
	req.Pro.CreatedByID = id
}

//JobsResult job result
type ProductsResults struct {
	Products []*Product `json:"products"`
}
