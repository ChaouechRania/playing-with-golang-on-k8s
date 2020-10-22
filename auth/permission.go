package auth

import (
	"context"
	"playing-with-golang-on-k8s/store"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Permission model for permissions
type Permission struct {
	ID string
}

//PermModifyJob permission to modify job
var PermModifyJob = &Permission{"modify_job"}

//PermCreateJob permission to modify job
var PermCreateJob = &Permission{"create_job"}

//PermModifyOrg permission to modify job
var PermModifyOrg = &Permission{"modify_org"}

//PermissionService the service for permissions
type PermissionService struct {
	UserStore store.UserStore

	Context context.Context
}

//NewPermissionService constructs PermissionService
func NewPermissionService(ctx context.Context, userStore store.UserStore) *PermissionService {
	return &PermissionService{

		UserStore: userStore,
		Context:   ctx,
	}
}

//HasPermissionToJob checks permissions for ops on jobs
func (perms *PermissionService) HasPermissionToJob(userID, jobID string, perm *Permission) bool {
	return true
}

//HasPermissionToOrganization checks permissions for ops on orgs
func (perms *PermissionService) HasPermissionToOrganization(userID, orgID string, perm *Permission) bool {
	return true
}

func sampleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if id == 0 {
			c.JSON(400, gin.H{"message": "invalid id"})
			c.Abort()
		}
	}
}
