package routes

import (
	"encoding/json"
	"net/http"
	"playing-with-golang-on-k8s/api"
	"playing-with-golang-on-k8s/auth"
	"playing-with-golang-on-k8s/service"
	"playing-with-golang-on-k8s/store"
	"time"

	"github.com/asaskevich/govalidator"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//UserActions represents users controller actions
type UserActions struct {
	UserService *service.UserService
}

//NewUserActions create a new UserActions
func NewUserActions(userService *service.UserService) *UserActions {
	return &UserActions{
		UserService: userService,
	}
}

//CreateUser creates a new user
func (as UserActions) CreateUser(c *gin.Context) {
	req := new(api.UserCreationRequest)

	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := as.UserService.Create(c.Request.Context(), req, time.Now())

	if err == store.ErrEmailTaken {
		c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, api.UserFromStore(user))
}

// PatchUser creates a new user
func (as UserActions) PatchUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "PatchUser"})
}

// GetUser retrieves a single user
func (as UserActions) GetUser(c *gin.Context) {
	id := c.Param("id")
	if id == "me" {
		authed := auth.ExtractAuthenticated(c)
		if authed == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Identification info required for creating an org"})
			return
		}
		id = authed.ID
	}

	user, err := as.UserService.Get(c.Request.Context(), id)

	if err == store.ErrNoSuchEntity {
		c.JSON(http.StatusOK, gin.H{"message": "User not found"})
		return
	}

	bytes, _ := json.Marshal(&user)
	c.Writer.Write(bytes)
}

// DeleteUser deletes a user
func (as UserActions) DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "DeleteUser"})
}

// GetUsers retrieve users
func (as UserActions) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "GetUsers"})
}
