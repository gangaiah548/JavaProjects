package controller

import (
	"net/http"
	"sdk_backend_service/src/models"
	"sdk_backend_service/src/services"

	"github.com/gin-gonic/gin"
)

type ISmpdataController interface {
	Createsmpdata(c *gin.Context)
	GetSmpdata(c *gin.Context)
	DeleteSmpdata(c *gin.Context)
	UpdateSmpData(c *gin.Context)
}
type SmpdataControllerImpl struct {
	sampdataService services.ISmpdataService
}

func (pdc *SmpdataControllerImpl) GetSmpdata(c *gin.Context) {
	active, count, err := pdc.sampdataService.GetSmpdata()

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
func (pdc *SmpdataControllerImpl) DeleteSmpdata(c *gin.Context) {

	// Create book
	key := c.Query("key")
	dproduct := models.Smpdata{
		Node: ""}

	productResponse, err := pdc.sampdataService.DeleteSmpdata(key, dproduct)
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
		map[string]interface{}{"data": "success", "node": productResponse.Node,
			"uuid": productResponse.Uuid})
}

func (pdc *SmpdataControllerImpl) UpdateSmpData(c *gin.Context) {
	key := c.Query("key")
	if key == "" || len(key) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"key is not empty or proper": key})
		return
	}
	var input models.Smpdata
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dproduct := models.Smpdata{
		Uuid: input.Uuid,
		Node: input.Node,
	}
	productResponse, err := pdc.sampdataService.UpdateSmpData(key, dproduct)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			map[string]interface{}{"data": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		map[string]interface{}{"data": "success", "uuid": productResponse.Uuid,
			"node": productResponse.Node,
		})
	return
}

// Createsmpdata implements IProductController.
func (pdc *SmpdataControllerImpl) Createsmpdata(c *gin.Context) {
	var input models.Smpdata
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dhouseAddress := models.Address{
		Hno: input.Houseaddress.Hno,
		Pin: input.Houseaddress.Pin,
	}
	dproduct := models.Smpdata{
		Node:         input.Node,
		Houseaddress: dhouseAddress}

	productResponse, err := pdc.sampdataService.CreateSmpdata(dproduct)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			map[string]interface{}{"data": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		map[string]interface{}{"data": "success",
			"uuid": productResponse.Uuid,
			"node": productResponse.Node,
			"Hno":  productResponse.Houseaddress.Hno,
			"Pin":  productResponse.Houseaddress.Pin})
	return
}

func Newsmpcreation() (ISmpdataController, error) {
	var smpdataDBconnectionService, vrr = services.NewSmpdatadepService()

	if vrr != nil {
		return nil, vrr
	}
	return &SmpdataControllerImpl{
		sampdataService: smpdataDBconnectionService,
	}, nil

}
