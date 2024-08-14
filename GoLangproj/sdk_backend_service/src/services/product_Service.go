package services

import (
	"fmt"
	"sdk_backend_service/src/config/logger"

	//"sdk_backend_service/src/errors"
	"sdk_backend_service/src/models"
	"sdk_backend_service/src/repository"
	//"strconv"
)

type IProductService interface {
	CreateProduct(productModel models.Product) (*models.Product, error)
	DeleteProduct(key string, productModel models.Product) (models.Product, error)
	UpdateProduct(key string, productModel models.Product) (*models.Product, error)
	//RemoteUpload(url models.Url) (string, error)
	//PrepareDeploymentModel(deploymentReq *http.Request) (models.ProcessDeploymentModel, error)
	GetProduct() ([]models.Product, int64, error)
}

type ProductserviceImpl struct {
	productRepository repository.IProductRepository
}

// UpdateProduct implements IProductService.
func (pds *ProductserviceImpl) UpdateProduct(key string, productModel models.Product) (*models.Product, error) {
	err := validate.Struct(productModel)
	if err != nil {
		return &productModel, err
	}
	dproduct := models.Product{
		Pid:         0,
		Name:        "",
		Description: "",
		Price:       00}

	response, err := pds.productRepository.FindProduct(key, &dproduct)

	//var name string = response.

	if dproduct.Name == "1" {
		logger.Info().Msg("Trying to persist update Db data")
		responseo, err := pds.productRepository.UpdateProduct(key, &productModel)
		if err != nil {
			logger.Error().Err(err)
			return response.(*models.Product), err
		}
		logger.Info().Msg("[✅] product creation persistence success!")

		return responseo, nil
	} else {
		return response.(*models.Product), nil
	}
}

// DeleteProduct implements IProductService.
func (pds *ProductserviceImpl) DeleteProduct(key string, productModel models.Product) (models.Product, error) {
	err := validate.Struct(productModel)
	if err != nil {
		return productModel, err
	}
	logger.Info().Msg("Trying to persist process deployment data")
	response, err := pds.productRepository.DeleteProduct(key, productModel)
	if err != nil {
		logger.Error().Err(err)
		return response.(models.Product), err
	}
	logger.Info().Msg("[✅] product creation persistence success!")

	return response.(models.Product), nil
}

// CreateProduct implements IProductService.
func (pds *ProductserviceImpl) CreateProduct(productModel models.Product) (*models.Product, error) {
	//validate
	err := validate.Struct(productModel)
	if err != nil {
		return &productModel, err
	}
	logger.Info().Msg("Trying to persist process deployment data")
	response, err := pds.productRepository.CreateProduct(&productModel)
	if err != nil {
		logger.Error().Err(err)
		return response, err
	}
	logger.Info().Msg("[✅] product creation persistence success!")

	// add deployed process to cache
	/*var deploymentResponse models.ProcessDeploymentModel = response.(models.ProcessDeploymentModel)
	if deploymentResponse.Status == string(constants.PROCESS_STATUS_ACTIVE) {
		bpmnEngineManager.AddProcessDefinitionToLocalCache(deploymentResponse)
		if deploymentModel.InstancePool > 0 {
			go bpmnEngineManager.CreateProcessInstancePool(deploymentResponse.Key, deploymentModel.InstancePool)
		}
		logger.Info().Msg("[✅] process added to cache and instance pool created!")
	}*/

	return response, nil
}

// GetProduct implements IProductService.
func (pds *ProductserviceImpl) GetProduct() ([]models.Product, int64, error) {
	active, count, err := pds.productRepository.GetProduct()
	if err != nil {
		logger.Error().Err(err)
		return active.([]models.Product), count, err
	}

	logger.Info().Msgf("There are %s active processes", fmt.Sprint(count))

	return active.([]models.Product), count, nil
}

func NewProductdepService() (IProductService, error) {
	var productRepository, err = repository.NewProductDepRepository()

	if err != nil {
		return nil, err
	}
	return &ProductserviceImpl{
		productRepository: productRepository,
	}, nil

}
