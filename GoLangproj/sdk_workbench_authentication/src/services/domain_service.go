package services

import (
	"fmt"
	"net/http"
	"sdk_workbench_authentication/src/config/logger"
	"sdk_workbench_authentication/src/errors"

	//"sdk_workbench_authentication/src/errors"
	"sdk_workbench_authentication/src/models"
	"sdk_workbench_authentication/src/repository"

	"github.com/google/uuid"
	//"strconv"
)

type IDomainService interface {
	CreateDomain(w http.ResponseWriter, req *http.Request, domainmodel models.Domain) (*models.Domain, error)
	DeleteDomain(w http.ResponseWriter, req *http.Request, key string, Domainmodel models.Domain) (models.Domain, error)
	UpdateDomain(key string, Domainmodel models.Domain) (models.Domain, error)
	//RemoteUpload(url models.Url) (string, error)
	//PrepareDeploymentModel(deploymentReq *http.Request) (models.ProcessDeploymentModel, error)
	GetDomains(w http.ResponseWriter, req *http.Request) (interface{}, int64, error)
	GetAllRoleData(w http.ResponseWriter, req *http.Request, ts string, role string, accessurl string) (models.UserWithRoleInfo, int64, error)
	CreateRoleData(w http.ResponseWriter, req *http.Request, entitlementmodel models.Entitlement) (*models.Entitlement, error)
}
type DomainserviceImpl struct {
	DomainRepository repository.IDomainRepository
}

// DeleteDomain implements IDomainService.
func (pds *DomainserviceImpl) DeleteDomain(w http.ResponseWriter, req *http.Request, key string, domainmodel models.Domain) (models.Domain, error) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err := validate.Struct(domainmodel)
	if err != nil {
		return domainmodel, err
	}
	logger.Info().Msg("Trying to persist process deployment data")
	response, err := pds.DomainRepository.DeleteDomain(key, domainmodel)
	if err != nil {
		logger.Error().Err(err)
		return response.(models.Domain), err
	}
	logger.Info().Msg("[✅] product creation persistence success!")

	return response.(models.Domain), nil
}

func (pds *DomainserviceImpl) CreateRoleData(w http.ResponseWriter, req *http.Request, entitlementmodel models.Entitlement) (*models.Entitlement, error) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	logger.Debug().Msg("Collection Open")
	//key, vrr := generateUniqueKey()
	//domainmodel.Uuid = key
	/*if vrr != nil {
		// handle error
		logger.Error().Msg("[🛑] failed to create collection")
		return &domainmodel, errors.New(vrr, "failed to create collection new UUID")
	}
	err := validate.Struct(domainmodel)
	if err != nil {
		return &domainmodel, err
	}*/
	logger.Info().Msg("Trying to persist process deployment data")
	response, err := pds.DomainRepository.CreateComponent(&entitlementmodel)
	if err != nil {
		logger.Error().Err(err)
		return response, err
	}
	logger.Info().Msg("[✅] product creation persistence success!")

	return response, nil
}

// CreateSampdata implements IProductService.
func (pds *DomainserviceImpl) CreateDomain(w http.ResponseWriter, req *http.Request, domainmodel models.Domain) (*models.Domain, error) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	logger.Debug().Msg("Collection Open")
	key, vrr := generateUniqueKey()
	domainmodel.Uuid = key
	if vrr != nil {
		// handle error
		logger.Error().Msg("[🛑] failed to create collection")
		return &domainmodel, errors.New(vrr, "failed to create collection new UUID")
	}
	err := validate.Struct(domainmodel)
	if err != nil {
		return &domainmodel, err
	}
	logger.Info().Msg("Trying to persist process deployment data")
	response, err := pds.DomainRepository.CreateDomain(&domainmodel)
	if err != nil {
		logger.Error().Err(err)
		return response, err
	}
	logger.Info().Msg("[✅] product creation persistence success!")

	return response, nil
}

func (pds *DomainserviceImpl) UpdateDomain(key string, domainmodel models.Domain) (models.Domain, error) {
	err := validate.Struct(domainmodel)
	if err != nil {
		return domainmodel, err
	}
	response := models.Domain{
		Uuid:   "err while update",
		Domain: ""}

	responseo, err := pds.DomainRepository.UpdateDomain(key, domainmodel)
	if err != nil {
		logger.Error().Err(err)
		return response, err
	}
	logger.Info().Msg("[✅] product creation persistence success!")

	return responseo, nil
	//response, err := pds.DomainRepository.FindDomain(key, &dproduct)

	//var name string = response.

	/*if dproduct.Uuid == "1" {
		logger.Info().Msg("Trying to persist update Db data")
		responseo, err := pds.DomainRepository.UpdateDomain(key, domainmodel)
		if err != nil {
			logger.Error().Err(err)
			return response.(models.Domain), err
		}
		logger.Info().Msg("[✅] product creation persistence success!")

		return responseo, nil
	} else {
		return response.(models.Domain), nil
	}*/
}

func (pds *DomainserviceImpl) GetDomains(w http.ResponseWriter, req *http.Request) (interface{}, int64, error) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	active, count, err := pds.DomainRepository.GetDomains()
	if err != nil {
		logger.Error().Err(err)
		return active, count, err
	}

	logger.Info().Msgf("There are %s active processes", fmt.Sprint(count))

	return active, count, nil
}

func (pds *DomainserviceImpl) GetAllRoleData(w http.ResponseWriter, req *http.Request, tokenstring string, role string, accessurl string) (models.UserWithRoleInfo, int64, error) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	entitlements, count, err := pds.DomainRepository.GetAllRoleData()
	us := models.UserWithRoleInfo{
		Token:       tokenstring,
		RoleName:    role,
		AccessUrl:   accessurl,
		Permissions: nil,
	}
	if err != nil {
		logger.Error().Err(err)
		return us, count, err
	}
	//associateRoleData := make(map[string][]models.Component)
	/*var tmpent []models.Entitlement
	d := []interface{}{entitlements}
	e := d[0].([]interface{})
	for i := range e {
		tmpent = append(tmpent, e[i].(models.Entitlement))
	}*/
	var componentUIArr []models.Component
	//var userWithRoleInfo UserWithRoleInfo
	//var c int=0
	for i := range entitlements {
		//strings.ToLower(dRoleName.String())
		if entitlements[i].Entitlement == "role1" { //have to make it dynamic whichver role is coming input we have to match and append those
			componentUIArr = append(componentUIArr, entitlements[i].Component)
		}
		//tmpent = append(tmpent, e[i].(models.Entitlement))
	}
	//associateRoleData["role1"] = componentUIArr
	userWithRoleInfo := models.UserWithRoleInfo{
		Token:       tokenstring,
		RoleName:    role,
		AccessUrl:   accessurl,
		Permissions: componentUIArr,
	}
	/*for key, el := range models.RoleMatrix {
		if key == "cmsuser" {
			arr := el
			for i := 0; i < len(arr); i++ {
				//fmt.Println(arr[i])
				//tmpar:=entitlements.
				associateRoleData[arr[i]] = entitlements

			}
		}
	}*/
	logger.Info().Msgf("There are %s active processes", fmt.Sprint(count))

	return userWithRoleInfo, count, nil
}

func NewDomaindepService() (IDomainService, error) {
	var DomainRepository, vrr = repository.NewJavaClassDepRepository()
	if vrr != nil {
		return nil, vrr
	}
	return &DomainserviceImpl{
		DomainRepository: DomainRepository,
	}, nil
}

func generateUniqueKey() (string, error) {
	return uuid.NewString(), nil
}
