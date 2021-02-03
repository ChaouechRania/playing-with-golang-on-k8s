package api

import (
	"playing-with-golang-on-k8s/shared"
	"playing-with-golang-on-k8s/store"
	"time"
)

//UserCreationRequest represents a view object to create a user
type UserCreationRequest struct {
	Email      string `json:"email" valid:"required,email~Invalid email"`
	Password   string `json:"password" valid:"required,runelength(3|50)~Password must have at least 3 characters"`
	RememberMe bool   `json:"RememberMe,omitempty"`
}

//ToStore converts a api.User to sotre.User object
func (request *UserCreationRequest) ToStore(createdAt time.Time) *store.User {
	return &store.User{
		Email:      request.Email,
		Password:   request.Password,
		RememberMe: request.RememberMe,
		CreatedAt:  &createdAt,
	}
}

//UserFromStore converts store.User to User
func UserFromStore(user *store.User) *User {
	return &User{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
	}
}

// User represents a view object used to read a user
type User struct {
	ID         string     `json:"id,omitempty"`
	Email      string     `json:"email,omitempty"`
	FirstName  string     `json:"firstname,omitempty"`
	LastName   string     `json:"lastname,omitempty"`
	RememberMe bool       `json:"RememberMe,omitempty"`
	CreatedAt  *time.Time `json:"createdAt,omitempty"`
	UpdatedAt  *time.Time `json:"updatedAt,omitempty"`
}

//Login represents auth data
type Login struct {
	Email    string `json:"email" valid:"email~Invalid email"`
	Password string `json:"password" valid:"required~The password is required"`
}

//AuthenticatedUser represents an authed user
type AuthenticatedUser struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}

//AuthUserFromStore converts user to authenticated user
func AuthUserFromStore(user *store.User) *AuthenticatedUser {
	return &AuthenticatedUser{
		ID:    user.ID,
		Email: user.Email,
	}
}

const identityKey = "id"

//OrganizationFromStore converts store.Organization to Organization
func ProductFromStore(org *store.Product) *shared.Product {
	return &shared.Product{
		ID:          org.ID,
		Slug:        org.Slug,
		Name:        org.Name,
		Description: org.Description,
		Price:       org.Price,
		Quantity:    org.Quantity,
		CreatedByID: org.CreatedByID,
		CreatedAt:   org.CreatedAt,
		UpdatedAt:   &org.UpdatedAt,
		DeletedAt:   org.DeletedAt,
	}
}
func ProductStore(org *store.Product) *Product {
	return &Product{
		ID:          org.ID,
		Name:        org.Name,
		Description: org.Description,
		CreatedAt:   org.CreatedAt,
		DeletedAt:   org.DeletedAt,
	}
}

//ToStore converts api object to the store one
func ToStoreProduct(org *shared.Product) *store.Product {
	return &store.Product{
		ID:          org.ID,
		Name:        org.Name,
		Description: org.Description,
		Price:       org.Price,
		Quantity:    org.Quantity,
		CreatedAt:   org.CreatedAt,
		DeletedAt:   org.DeletedAt,
	}
}

//ProductsResult represents the result of orgs listing
type ProductsResult struct {
	Products []*shared.Product `json:"products"`
}

//ToProductsResult converts orgs to orgs result object
func ToProductsResult(pros []*store.Product) ProductsResult {
	reads := make([]*shared.Product, len(pros))

	for i, pro := range pros {
		reads[i] = ProductFromStore(pro)
	}

	return ProductsResult{Products: reads}
}
