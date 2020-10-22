package service

import (
	"context"
	"playing-with-golang-on-k8s/api"
	"playing-with-golang-on-k8s/shared"
	"playing-with-golang-on-k8s/store"
	"time"

	"github.com/pkg/errors"
)

//ProService service logic for users
type ProService struct {
	proStore store.ProductStore
}

//NewProService constructs a new ProService
func NewProService(proStore store.ProductStore) *ProService {
	return &ProService{
		proStore: proStore,
	}
}

//CreateOrg create a new org
func (ps *ProService) CreatePro(ctx context.Context, req *api.ProUpsertRequest) (*shared.Product, error) {
	if req.Pro == nil {
		return nil, errors.New("org cannot be nil in creation req")
	}
	created, err := ps.proStore.CreatePro(ctx, store.ProFromAPI(req.Pro))
	if err != nil {
		return nil, err
	}
	return api.ProductFromStore(created), nil
}

//Update update a job
func (ps *ProService) Update(ctx context.Context, req *api.ProUpsertRequest, userID string) (*shared.Product, error) {
	if req.UpdateMask == nil || len(req.UpdateMask) == 0 {
		return nil, errors.New("Update mask is required when doing update")
	}
	existing, err := ps.proStore.GetPro(ctx, req.Pro.ID)
	if err != nil {
		return nil, err
	}
	updated, err := updatePro(existing, req)
	if err != nil {
		return nil, err
	}
	result, err := ps.proStore.UpdatePro(ctx, updated)
	if err != nil {
		return nil, err
	}
	return api.ProductFromStore(result), nil
}

func updatePro(existing *store.Product, req *api.ProUpsertRequest) (*store.Product, error) {
	update := store.ProFromAPI(req.Pro)
	var updated bool
	for _, s := range req.UpdateMask {
		switch s {
		case api.ProNameMask:
			updated = existing.Name != update.Name || updated
			existing.Name = update.Name
		case api.ProDescriptionMask:
			updated = existing.Description != update.Description || updated
			existing.Description = update.Description
		}
	}
	if updated {
		now := time.Now()
		existing.UpdatedAt = now
	}

	return existing, nil
}
