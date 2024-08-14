package repository

import (
	"context"
	"fmt"

	arangodbdriver "github.com/arangodb/go-driver"

	"sdk_backend_service/src/clients"
	"sdk_backend_service/src/config/env"
	"sdk_backend_service/src/config/logger"
	"sdk_backend_service/src/errors"
	"sdk_backend_service/src/models"
	"sdk_backend_service/src/repository/arango/query"
)

type ISmpdataRepository interface {
	CreateSmpdata(im *models.Smpdata) (*models.Smpdata, error)
	GetSmpdata() (interface{}, int64, error)
	DeleteSmpdata(key string, im interface{}) (interface{}, error)
	UpdateSmpdata(inpkey string, im *models.Smpdata) (*models.Smpdata, error)
	FindSmpdata(prm string, im *models.Smpdata) (interface{}, error)
}

type SmpdataRepositoryImpl struct {
	dbClient                 arangodbdriver.Client
	databaseName             string
	deploymentCollectionName string
}

// DeleteSmpdata implements ISmpdataRepository.
func (r *SmpdataRepositoryImpl) DeleteSmpdata(_key string, im interface{}) (interface{}, error) {
	//var abc string  = key

	//fmt.Printf("queryBuilder %s", qs)
	//panic("unimplemented")
	ctx := context.Background()
	logger.Debug().Msg("Open DB")
	// Open database
	db, err := r.dbClient.Database(ctx, r.databaseName)

	//db, err := r.dbClient.Database(pCtx, r.databaseName)

	logger.Info().Msgf("There are %s after Db %s connect", db.Name(), r.databaseName)

	collection, err := db.Collection(ctx, r.deploymentCollectionName)

	meta, err := collection.RemoveDocument(ctx, _key)

	if err != nil {
		logger.Error().Msg("[🛑] failed to remove doc in collection please eneter proper details _key " + _key + " not proper")
		return im, errors.New(err, "failed to create processDeployment please eneter proper details _key "+_key+" not proper")
	}
	logger.Debug().Msg("Persist Done meta [" + meta.ID.Key() + "]" + "[" + meta.ID.Collection() + "]" + "[" + meta.OldRev + "]" + "[" + meta.Rev + "]")

	response := im.(models.Smpdata)
	response.Uuid = _key
	response.Node = meta.ID.String() + meta.ID.Key()
	return response, nil
}

func NewSmpdataDepRepository() (ISmpdataRepository, error) {
	logger.Debug().Msg("Creating Arango Client")
	dbc, err := clients.ArangoDBConnect(env.GetProperties().ArangoAddr, env.GetProperties().ArangoUser, env.GetProperties().ArangoPwd)

	if err != nil {
		logger.Error().Err(errors.New(err, "[🛑] Error connecting DB!"))
		return nil, errors.New(err, "Error connecting DB!")
	}
	logger.Debug().Msg("Arango Client Created")
	return &SmpdataRepositoryImpl{
		dbClient:                 dbc,
		databaseName:             env.GetProperties().ArangoDbName,
		deploymentCollectionName: env.GetProperties().ArangoCollectionName,
	}, nil
}

func (r *SmpdataRepositoryImpl) UpdateSmpdata(_key string, im *models.Smpdata) (*models.Smpdata, error) {
	ctx := context.Background()
	logger.Debug().Msg("Open DB")
	// Open database
	//ctx := context.Background()
	logger.Debug().Msg("Open DB")
	// Open database
	db, err := r.dbClient.Database(ctx, r.databaseName)

	//db, err := r.dbClient.Database(pCtx, r.databaseName)

	logger.Info().Msgf("There are %s after Db %s connect", db.Name(), r.databaseName)

	collection, err := db.Collection(ctx, r.deploymentCollectionName)

	meta, err := collection.ReplaceDocument(ctx, _key, im)

	if err != nil {
		logger.Error().Msgf("[🛑] failed to open collection on %s database " + r.databaseName)
		return im, errors.New(err, fmt.Sprintf("failed to open %s collection on %s database", r.deploymentCollectionName,
			r.databaseName))
	}
	// Persistent Store
	//meta, err := collection.CreateDocument(ctx, im)
	if err != nil {
		logger.Error().Msg("[🛑] failed updateSmpdata with _key" + _key)
		return im, errors.New(err, "failed update with _key"+_key)
	}
	logger.Debug().Msg("Persist Done meta [" + meta.ID.Key() + "]" + "[" + meta.ID.Collection() + "]" + "[" + meta.OldRev + "]" + "[" + meta.Rev + "]")

	response := im
	response.Uuid = "im.Pid"
	response.Node = "from update methods"

	return response, nil
}

func (r *SmpdataRepositoryImpl) CreateSmpdata(im *models.Smpdata) (*models.Smpdata, error) {
	ctx := context.Background()
	logger.Debug().Msg("Open DB")
	// Open database
	db, err := r.dbClient.Database(ctx, r.databaseName)

	if err != nil {
		logger.Error().Msgf("[🛑] failed to open %s database ", r.databaseName)
		return im, errors.New(err, fmt.Sprintf("failed to open %s database", r.databaseName))
	}
	logger.Debug().Msg("DB Open")
	logger.Debug().Msg("Open Collection")
	// Open collection
	collection, err := db.Collection(ctx, r.deploymentCollectionName)

	if err != nil {
		logger.Error().Msgf("[🛑] failed to open collection on %s database " + r.databaseName)
		return im, errors.New(err, fmt.Sprintf("failed to open %s collection on %s database", r.deploymentCollectionName,
			r.databaseName))
	}

	//im.Uuid = key

	logger.Debug().Msg("Persist Open")
	// Persistent Store
	meta, err := collection.CreateDocument(ctx, im)
	if err != nil {
		logger.Error().Msg("[🛑] failed to create collection")
		return im, errors.New(err, "failed to create collection")
	}
	logger.Debug().Msg("Persist Done meta [" + meta.ID.Key() + "]" + "[" + meta.ID.Collection() + "]" + "[" + meta.OldRev + "]" + "[" + meta.Rev + "]")

	response := im
	response.Uuid = im.Uuid
	response.Node = meta.ID.Key() + "]" + "[" + meta.ID.Collection() + "]" + "[" + meta.OldRev + "]" + "[" + meta.Rev + "]"
	response.Houseaddress = models.Address{Hno: "testing",
		Pin: "345645646"}
	return response, nil
}

func (r *SmpdataRepositoryImpl) GetSmpdata() (interface{}, int64, error) {
	queryBuilder := query.NewForQuery(r.deploymentCollectionName, "doc")
	qs := queryBuilder.Filter("document.procedures.node", "NOT LIKE", "'%rrrr%'").Done().Return().String()
	//FOR document IN processDeploymentCollection
	// FILTER NOT LIKE(document.procedures.node, '%hemogram%') RETURN document
	fmt.Printf("queryBuilder %s", qs)

	pCtx := context.Background()
	ctx := arangodbdriver.WithQueryCount(pCtx)

	// Open database
	db, err := r.dbClient.Database(pCtx, r.databaseName)

	logger.Info().Msgf("There are %s after Db %s connect", db.Name(), r.databaseName)

	if err != nil {
		logger.Warn().Msgf("[⚠️] failed to open %s database.", r.databaseName)
		return nil, 0, errors.New(err, fmt.Sprintf("failed to open [%s] database", r.databaseName))
	}

	cursor, err := db.Query(ctx, qs, nil)
	if err != nil {
		logger.Warn().Msgf("[⚠️] failed to query [%s] on %s database", qs, r.databaseName)
		return nil, 0, errors.New(err, fmt.Sprintf("failed to query [%s] on %s database", qs, r.databaseName))
	}

	//defer cursor.Close()

	result := []models.Smpdata{}
	if cursor.Count() == 0 || !cursor.HasMore() {
		return result, cursor.Count(), nil
	}

	for {
		process := models.Smpdata{}
		_, err = cursor.ReadDocument(pCtx, &process)
		if arangodbdriver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, 0, errors.New(err, fmt.Sprintf("failed to read response from cursor for query [%s]", qs))
		}

		result = append(result, process)
	}

	return result, cursor.Count(), nil
}

func (r *SmpdataRepositoryImpl) FindSmpdata(prm string, im *models.Smpdata) (interface{}, error) {
	ctx := context.Background()
	db, err := r.dbClient.Database(ctx, r.databaseName)

	//db, err := r.dbClient.Database(pCtx, r.databaseName)

	logger.Info().Msgf("There are %s after Db %s connect", db.Name(), r.databaseName)

	collection, err := db.Collection(ctx, r.deploymentCollectionName)

	meta, err := collection.DocumentExists(ctx, prm) //readdocument also we can use

	if err != nil {
		logger.Error().Msgf("[🛑] failed to open collection on %s database " + r.databaseName)
		return im, errors.New(err, fmt.Sprintf("failed to open %s collection on %s database", r.deploymentCollectionName,
			r.databaseName))
	}
	var flag string = ""
	if meta == true {
		flag = "1"
	} else {
		flag = "0"
	}
	//meta, err := collection.CreateDocument(ctx, im)
	if err != nil {
		logger.Error().Msg("[🛑] failed to create processDeployment")
		return im, errors.New(err, "failed to create processDeployment")
	}
	logger.Debug().Msgf("search is Done meta %t", meta)

	im.Node = flag

	response := im
	response.Uuid = "1"
	response.Node = flag
	return response, nil
}
