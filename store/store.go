package store

import (
	"context"
	"errors"
)

//ErrEmailTaken is returned when trying to create a user the a taken email address
var ErrEmailTaken = errors.New("Email address already taken")

//ErrNoSuchEntity error for not found
var ErrNoSuchEntity = errors.New("No such entity exists")

//ErrCannotPublishDeletedJob error for not found
var ErrCannotPublishDeletedJob = errors.New("Cannot publish a deleted posting")

//ErrJobAlreadyPublished posting has already been published status
var ErrJobAlreadyPublished = errors.New("Job already posted")

//ErrJobAlreadyArchived posting has already been archived
var ErrJobAlreadyArchived = errors.New("Job already archived")

//ErrInvalidJobStatusTransition error when trying to proceed to an invalid job status change
var ErrInvalidJobStatusTransition = errors.New("Invalid job status transition")

//ErrHasAlreadyProfile is returned when trying to create a user the a taken email address
var ErrHasAlreadyProfile = errors.New("User has already a profile")

//ErrOrgAlreadyExists is returned when trying to create an ogr that exists
var ErrOrgAlreadyExists = errors.New("The org already exists")
var ErrProAlreadyExists = errors.New("The product already exists")

//JobStore represents the interface to manage jobs storage
/*type JobStore interface {
	Create(context.Context, *Job) (*Job, error)
	Update(context.Context, *Job) (*Job, error)
	Delete(context.Context, string) error
	Get(context.Context, string) (*Job, error)
	List(ctx context.Context, offset int, limit int) ([]Job, error)
	Count(context.Context) (int, error)
	Publish(context.Context, string) error
	Archive(context.Context, string) error
}*/

//UserStore represents the interface to manage users storage
type UserStore interface {
	CreateUser(context.Context, *User) (*User, error)
	GetUser(context.Context, string) (*User, error)
	IsEmailTaken(context.Context, string) (bool, error)
	Authenticate(context.Context, string, string) (*User, error)
}

//ProductStore the store interface for products
type ProductStore interface {
	CreatePro(context.Context, *Product) (*Product, error)
	GetPro(context.Context, string) (*Product, error)
	UpdatePro(context.Context, *Product) (*Product, error)
	DeletePro(context.Context, string) error
	ListProds(ctx context.Context, req *GetProdsRequest, offset, limit int) ([]*Product, error)
	ProAlreadyExists(ctx context.Context, org *Product) (bool, error)
}

//ProfileStore represents the interface to manage recruiters storage
/*type ProfileStore interface {
	CreateProfile(context.Context, *Profile) (*Profile, error)
	Authenticate(context.Context, string, string) (*User, error)
	HasAlreadyProfile(context.Context, *User) (bool, error)
	GetForUser(context.Context, string) (*Profile, error)
}*/

//OrganizationStore the store interface for organizations
/*type OrganizationStore interface {
	CreateOrg(context.Context, *Organization) (*Organization, error)
	GetOrg(context.Context, string) (*Organization, error)
	UpdateOrg(context.Context, *Organization) (*Organization, error)
	DeleteOrg(context.Context, string) error
	ListOrgs(ctx context.Context, req *GetOrgsRequest, offset, limit int) ([]*Organization, error)
	OrgAlreadyExists(ctx context.Context, org *Organization) (bool, error)
}*/
