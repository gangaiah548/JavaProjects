package services

import (
	"fmt"
	"sdk_backend_service/src/config/logger"
	"sdk_backend_service/src/errors"

	//"sdk_backend_service/src/errors"
	"sdk_backend_service/src/models"
	"sdk_backend_service/src/repository"

	"github.com/google/uuid"
	//"strconv"
)

type ISmpdataService interface {
	CreateSmpdata(smpdatamodel models.Smpdata) (*models.Smpdata, error)
	DeleteSmpdata(key string, smpdatamodel models.Smpdata) (models.Smpdata, error)
	UpdateSmpData(key string, smpdatamodel models.Smpdata) (*models.Smpdata, error)
	//RemoteUpload(url models.Url) (string, error)
	//PrepareDeploymentModel(deploymentReq *http.Request) (models.ProcessDeploymentModel, error)
	GetSmpdata() ([]models.Smpdata, int64, error)
}
type SmpdataserviceImpl struct {
	smpdataRepository repository.ISmpdataRepository
}

// DeleteSmpdata implements ISmpdataService.
func (pds *SmpdataserviceImpl) DeleteSmpdata(key string, smpdatamodel models.Smpdata) (models.Smpdata, error) {
	err := validate.Struct(smpdatamodel)
	if err != nil {
		return smpdatamodel, err
	}
	logger.Info().Msg("Trying to persist process deployment data")
	response, err := pds.smpdataRepository.DeleteSmpdata(key, smpdatamodel)
	if err != nil {
		logger.Error().Err(err)
		return response.(models.Smpdata), err
	}
	logger.Info().Msg("[✅] product creation persistence success!")

	return response.(models.Smpdata), nil
}

// CreateSampdata implements IProductService.
func (pds *SmpdataserviceImpl) CreateSmpdata(smpdatamodel models.Smpdata) (*models.Smpdata, error) {
	logger.Debug().Msg("Collection Open")
	key, vrr := generateUniqueKey()
	smpdatamodel.Uuid = key
	if vrr != nil {
		// handle error
		logger.Error().Msg("[🛑] failed to create collection")
		return &smpdatamodel, errors.New(vrr, "failed to create collection new UUID")
	}
	err := validate.Struct(smpdatamodel)
	if err != nil {
		return &smpdatamodel, err
	}
	logger.Info().Msg("Trying to persist process deployment data")
	response, err := pds.smpdataRepository.CreateSmpdata(&smpdatamodel)
	if err != nil {
		logger.Error().Err(err)
		return response, err
	}
	logger.Info().Msg("[✅] product creation persistence success!")

	return response, nil
}
func (pds *SmpdataserviceImpl) UpdateSmpData(key string, productModel models.Smpdata) (*models.Smpdata, error) {
	err := validate.Struct(productModel)
	if err != nil {
		return &productModel, err
	}
	dproduct := models.Smpdata{
		Uuid: "0",
		Node: ""}

	response, err := pds.smpdataRepository.FindSmpdata(key, &dproduct)

	//var name string = response.

	if dproduct.Node == "1" {
		logger.Info().Msg("Trying to persist update Db data")
		responseo, err := pds.smpdataRepository.UpdateSmpdata(key, &productModel)
		if err != nil {
			logger.Error().Err(err)
			return response.(*models.Smpdata), err
		}
		logger.Info().Msg("[✅] product creation persistence success!")

		return responseo, nil
	} else {
		return response.(*models.Smpdata), nil
	}
}

func (pds *SmpdataserviceImpl) GetSmpdata() ([]models.Smpdata, int64, error) {
	active, count, err := pds.smpdataRepository.GetSmpdata()
	if err != nil {
		logger.Error().Err(err)
		return active.([]models.Smpdata), count, err
	}

	logger.Info().Msgf("There are %s active processes", fmt.Sprint(count))

	return active.([]models.Smpdata), count, nil
}

func generateUniqueKey() (string, error) {
	return uuid.New().String(), nil
}

func NewSmpdatadepService() (ISmpdataService, error) {
	var smpdataRepository, vrr = repository.NewSmpdataDepRepository()
	if vrr != nil {
		return nil, vrr
	}
	return &SmpdataserviceImpl{
		smpdataRepository: smpdataRepository,
	}, nil
}
