package api

import "playing-with-golang-on-k8s/shared"

//Org update masks
const (
	ProNameMask        = "pro.name"
	ProDescriptionMask = "pro.description"
)

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
