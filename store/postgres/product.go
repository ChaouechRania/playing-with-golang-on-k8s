package postgres

import (
	"context"
	"playing-with-golang-on-k8s/store"
	"strings"
	"time"

	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
	gonanoid "github.com/matoous/go-nanoid"
	"github.com/pkg/errors"
)

//ListProds lists prods for the current page
func (c *Client) ListProds(ctx context.Context, req *store.GetProdsRequest, offset int, limit int) ([]*store.Product, error) {
	os := []*store.Product{}
	err := c.db.Set("gorm:auto_preload", true).Model(&store.Product{}).
		Where("created_by_id=?", req.UserID).
		Preload("CreatedBy").
		Offset(offset).
		Limit(limit).
		Find(&os).Error
	if err != nil {
		return nil, err
	}

	return os, err
}

//DeletePro deletes the product with given ID
func (c *Client) DeletePro(ctx context.Context, id string) error {
	org, err := c.GetPro(ctx, id)
	if err != nil {
		return err
	}
	if org == nil {
		return store.ErrNoSuchEntity
	}

	now := time.Now()
	org.DeletedAt = &now
	_, err = c.UpdatePro(ctx, org)
	return err
}

//UpdateOrg updates an oganization
func (c *Client) UpdatePro(ctx context.Context, org *store.Product) (*store.Product, error) {
	existing, err := c.GetPro(ctx, org.ID)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	org.CreatedAt = existing.CreatedAt
	org.UpdatedAt = now
	org.Slug = slug.Make(org.Name)
	err = c.db.Where(store.Product{ID: org.ID}).Save(org).Error
	if err != nil {
		return nil, err
	}
	return org, nil
}

//GetPro an product
func (c *Client) GetPro(ctx context.Context, id string) (*store.Product, error) {
	org := new(store.Product)
	err := c.db.Where(store.Product{ID: id}).Take(org).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, store.ErrNoSuchEntity
		}
		return nil, err
	}

	return org, nil
}

//CreatePro creates a new product
func (c *Client) CreatePro(ctx context.Context, organization *store.Product) (*store.Product, error) {
	exists, err := c.ProAlreadyExists(ctx, organization)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, store.ErrOrgAlreadyExists
	}
	now := time.Now()
	organization.CreatedAt = now
	organization.UpdatedAt = now
	organization.Slug = slug.Make(organization.Name)

	id, err := gonanoid.Nanoid()
	if err != nil {
		return nil, errors.Wrap(err, "generating job id")
	}
	organization.ID = id
	err = c.db.Save(organization).Error
	if err != nil {
		return nil, err
	}
	return organization, nil
}

//ProAlreadyExists tests if the provided product name already exists
func (c *Client) ProAlreadyExists(ctx context.Context, org *store.Product) (bool, error) {
	slug := strings.ToLower(slug.Make(org.Name))
	count := 0
	err := c.db.Model(&store.Product{}).Where(&store.Product{Slug: slug}).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, err
}
