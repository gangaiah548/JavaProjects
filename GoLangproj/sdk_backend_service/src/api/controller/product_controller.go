package controller

import (
	"net/http"
	"sdk_backend_service/src/models"
	"sdk_backend_service/src/services"

	"github.com/gin-gonic/gin"
)

type IProductController interface {
	Create(c *gin.Context)
	UpdateProduct(c *gin.Context)
	GetProducts(c *gin.Context)
	DeleteProduct(c *gin.Context)
}
type ProductControllerImpl struct {
	productService services.IProductService
}

// / GetProducts implements IProductController.
func (pdc *ProductControllerImpl) GetProducts(c *gin.Context) {
	active, count, err := pdc.productService.GetProduct()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			map[string]interface{}{"err": err,
				"objectValue": map[string]interface{}{"arrayValue": active}},
		)
		return
	}

	//logger.Info().Msgf("There are %s active processes", fmt.Sprint(count))

	c.JSON(
		http.StatusOK,
		map[string]interface{}{"count": count,
			"objectValue": map[string]interface{}{"arrayValue": active}},
	)

}
func (pdc *ProductControllerImpl) DeleteProduct(c *gin.Context) {

	// Create book
	key := c.Query("key")
	dproduct := models.Product{
		Name:        "",
		Description: "",
		Price:       0}

	productResponse, err := pdc.productService.DeleteProduct(key, dproduct)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			map[string]interface{}{"data": err.Error()},
		)
		return
	}

	// returning empty string since process definition as return is not required as it will unnecessarily increase the return payload size
	//deploymentResponse.Definition = ""

	c.JSON(
		http.StatusOK,
		map[string]interface{}{"data": "success", "Keydetails": productResponse.Name,
			"description": productResponse.Description})
}

func (pdc *ProductControllerImpl) UpdateProduct(c *gin.Context) {
	key := c.Query("key")
	if key == "" || len(key) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"key is not empty or proper": key})
		return
	}
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dproduct := models.Product{
		Pid:         input.Pid,
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price}

	productResponse, err := pdc.productService.UpdateProduct(key, dproduct)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			map[string]interface{}{"data": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		map[string]interface{}{"data": "success", "pid": productResponse.Pid,
			"Keydetails":  productResponse.Name,
			"description": productResponse.Description, "price": productResponse.Price})
	return
}

// Deployment implements IProductController.
func (pdc *ProductControllerImpl) Create(c *gin.Context) {

	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dproduct := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price}

	productResponse, err := pdc.productService.CreateProduct(dproduct)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			map[string]interface{}{"data": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		map[string]interface{}{"data": "success", "pid": productResponse.Pid,
			"Keydetails":  productResponse.Name,
			"description": productResponse.Description, "price": productResponse.Price})
	return
}

func Newcreation() (IProductController, error) {
	var productDBconnectionService, err = services.NewProductdepService()

	if err != nil {
		return nil, err
	}
	return &ProductControllerImpl{
		productService: productDBconnectionService,
	}, nil

}
