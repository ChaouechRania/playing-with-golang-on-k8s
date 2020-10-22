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

//ProfileRequest user profile
/*type ProfileRequest struct {
	JobTitle         string `json:"jobTitle" valid:"runelength(2|50)~Password must have at least 2 characters"`
	OrganizationName string `json:"organizationName" valid:"runelength(2|50)~Password must have at least 2 characters"`
}*/

//RecruiterCreationRequest represents a view object to create a recruiter
type RecruiterCreationRequest struct {
	UserCreationRequest
	FirstName        string `json:"firstname" valid:"runelength(2|50)~Password must have at least 2 characters"`
	LastName         string `json:"lastname" valid:"runelength(2|50)~Password must have at least 2 characters"`
	JobTitle         string `json:"jobTitle" valid:"runelength(2|50)~Password must have at least 2 characters"`
	OrganizationName string `json:"organizationName" valid:"required,runelength(2|50)~organizationName must have at least 2 characters"`
}

//ToUserCreationRequest to user creation request
func (request *RecruiterCreationRequest) ToUserCreationRequest() *UserCreationRequest {
	return &UserCreationRequest{
		Email:      request.Email,
		Password:   request.Password,
		RememberMe: true,
	}
}

//ToStoreUser converts a api.User to sotre.User object
func (request *RecruiterCreationRequest) ToStoreUser(createdAt time.Time) *store.User {
	return &store.User{
		Email:      request.Email,
		Password:   request.Password,
		RememberMe: true,
		CreatedAt:  &createdAt,
	}
}

//ToStoreProfile converts a api.User to sotre.User object
/*func (request *ProfileRequest) ToStoreProfile(userID string, createdAt time.Time) *store.Profile {
	return &store.Profile{
		JobTitle:         request.JobTitle,
		OrganizationName: request.OrganizationName,
		UserID:           userID,
		CreatedAt:        &createdAt,
	}
}*/

//Profile represents a user profile
/*type Profile struct {
	ID               string     `json:"id,omitempty" json:"_id,omitempty"`
	JobTitle         string     `json:"jobTitle,omitempty"`
	OrganizationName string     `json:"organizationName,omitempty"`
	UserID           string     `json:"userID,omitempty"`
	CreatedAt        *time.Time `json:"createdAt"`
	UpdatedAt        time.Time  `json:"updatedAt"`
}*/

//ProfileFromStore converts store.User to User
/*func ProfileFromStore(profile *store.Profile) *Profile {
	return &Profile{
		ID:               profile.ID,
		JobTitle:         profile.JobTitle,
		OrganizationName: profile.OrganizationName,
		UserID:           profile.UserID,
		CreatedAt:        profile.CreatedAt,
		UpdatedAt:        profile.UpdatedAt,
	}
}*/

//ToProfileRequest converts a api.RecruiterCreationRequest to api.ProfileRequest object
/*func (request *RecruiterCreationRequest) ToProfileRequest() *ProfileRequest {
	return &ProfileRequest{
		JobTitle:         request.JobTitle,
		OrganizationName: request.OrganizationName,
	}
}*/

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

//AddressCreate represents a postal address as defined in schema.org
type AddressCreate struct {
	Locality    string `json:"locality"`
	Street      string `json:"street"`
	Region      string `json:"region"`
	PostalCode  string `json:"postalCode"`
	CountryCode string `json:"countryCode"`
}

//OrganizationFromStore converts store.Organization to Organization
// func OrganizationFromStore(org *store.Organization) *shared.Organization {
// 	return &shared.Organization{
// 		ID:           org.ID,
// 		Slug:         org.Slug,
// 		Name:         org.Name,
// 		Logo:         org.Logo,
// 		Website:      org.Website,
// 		Description:  org.Description,
// 		FoundingDate: org.FoundingDate,
// 		CreatedByID:  org.CreatedByID,
// 		CreatedAt:    org.CreatedAt,
// 		UpdatedAt:    &org.UpdatedAt,
// 		DeletedAt:    org.DeletedAt,
// 	}
// }

//OrganizationFromStore converts store.Organization to Organization
func ProductFromStore(org *store.Product) *shared.Product {
	return &shared.Product{
		ID:          org.ID,
		Slug:        org.Slug,
		Name:        org.Name,
		Description: org.Description,
		CreatedByID: org.CreatedByID,
		CreatedAt:   org.CreatedAt,
		UpdatedAt:   &org.UpdatedAt,
		DeletedAt:   org.DeletedAt,
	}
}

//ToStore converts api object to the store one
// func ToStore(org *shared.Organization) *store.Organization {
// 	return &store.Organization{
// 		ID:           org.ID,
// 		Name:         org.Name,
// 		Logo:         org.Logo,
// 		Website:      org.Website,
// 		Description:  org.Description,
// 		FoundingDate: org.FoundingDate,
// 		CreatedAt:    org.CreatedAt,
// 		DeletedAt:    org.DeletedAt,
// 	}
// }

//ToStore converts api object to the store one
func ToStoreProduct(org *shared.Product) *store.Product {
	return &store.Product{
		ID:          org.ID,
		Name:        org.Name,
		Description: org.Description,
		CreatedAt:   org.CreatedAt,
		DeletedAt:   org.DeletedAt,
	}
}

// //OrganizationsResult represents the result of orgs listing
// type OrganizationsResult struct {
// 	Organizations []*shared.Organization `json:"organizations"`
// }

//OrganizationsResult represents the result of orgs listing
type ProductsResult struct {
	Products []*shared.Product `json:"products"`
}

// //ToOrganizationsResult converts orgs to orgs result object
// func ToOrganizationsResult(orgs []*store.Organization) OrganizationsResult {
// 	reads := make([]*shared.Organization, len(orgs))

// 	for i, org := range orgs {
// 		reads[i] = OrganizationFromStore(org)
// 	}

// 	return OrganizationsResult{Organizations: reads}
// }

//ToOrganizationsResult converts orgs to orgs result object
func ToProductsResult(pros []*store.Product) ProductsResult {
	reads := make([]*shared.Product, len(pros))

	for i, pro := range pros {
		reads[i] = ProductFromStore(pro)
	}

	return ProductsResult{Products: reads}
}
