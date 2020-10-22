package routes

import (
	"fmt"
	"net/http"
	"playing-with-golang-on-k8s/api"
	"playing-with-golang-on-k8s/auth"
	"playing-with-golang-on-k8s/service"
	"playing-with-golang-on-k8s/store"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//ProductActions handles requests related to organizations
type ProductActions struct {
	Prods        store.ProductStore
	ProService   *service.ProService
	PermsService *auth.PermissionService
}

//NewProductActions create a new actions object for orgs.
func NewProductActions(store store.ProductStore, proService *service.ProService, permService *auth.PermissionService) *ProductActions {
	return &ProductActions{
		Prods:        store,
		ProService:   proService,
		PermsService: permService,
	}
}

//Create handles org creation logic
func (as ProductActions) Create(c *gin.Context) {
	authed := auth.ExtractAuthenticated(c)
	if authed == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Identification info required for creating an org"})
		return
	}

	req := new(api.ProUpsertRequest)
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

	/*req := new(api.Organization)
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}*/

	req.WithCreatedBy(authed.ID)
	product, err := as.ProService.CreatePro(c.Request.Context(), req)

	if err != nil {
		if err == store.ErrProAlreadyExists {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, product)
}

//Get retrieves an organization
func (as ProductActions) Get(c *gin.Context) {
	authed := auth.ExtractAuthenticated(c)
	if authed == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Identification info required for getting an org"})
		return
	}

	id := c.Param("id")
	pro, err := as.Prods.GetPro(c.Request.Context(), id)
	if err != nil {
		if err == store.ErrNoSuchEntity {
			c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("No org with ID : %s", id)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, api.ProductFromStore(pro))
}

//List retrieves an organization listing
func (as ProductActions) List(c *gin.Context) {
	authed := auth.ExtractAuthenticated(c)
	if authed == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Identification info required for getting orgs"})
		return
	}

	req := store.GetProdsRequest{UserID: authed.ID}
	prods, err := as.Prods.ListProds(c.Request.Context(), &req, 0, 100)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, api.ToProductsResult(prods))
}

//Delete retrieves an organization
func (as ProductActions) Delete(c *gin.Context) {
	id := c.Param("id")
	err := as.Prods.DeletePro(c.Request.Context(), id)
	if err != nil {
		if err == store.ErrNoSuchEntity {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Org with %s deleted with success", id)})
}

//Update retrieves an organization
func (as ProductActions) Update(c *gin.Context) {
	authed := auth.ExtractAuthenticated(c)
	if authed == nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Invalid or no token"})
	}

	req := new(api.ProUpsertRequest)
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

	id := c.Param("id")
	req.Pro.ID = id
	pro, err := as.ProService.Update(c.Request.Context(), req, authed.ID)
	if err != nil {
		if err == store.ErrNoSuchEntity {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pro)
}
