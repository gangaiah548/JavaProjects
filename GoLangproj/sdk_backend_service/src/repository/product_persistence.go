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

type IProductRepository interface {
	CreateProduct(im *models.Product) (*models.Product, error)
	GetProduct() (interface{}, int64, error)
	DeleteProduct(key string, im interface{}) (interface{}, error)
	UpdateProduct(inpkey string, im *models.Product) (*models.Product, error)
	FindProduct(prm string, im *models.Product) (interface{}, error)
}

type ProductRepositoryImpl struct {
	dbClient                 arangodbdriver.Client
	databaseName             string
	deploymentCollectionName string
}

// DeleteProduct implements IProductRepository.
func (r *ProductRepositoryImpl) DeleteProduct(key string, im interface{}) (interface{}, error) {
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

	meta, err := collection.RemoveDocument(ctx, key)

	if err != nil {
		logger.Error().Msg("[🛑] failed to remove doc in processDeployment please eneter proper details")
		return im, errors.New(err, "failed to create processDeployment please eneter proper details")
	}
	logger.Debug().Msg("Persist Done meta [" + meta.ID.Key() + "]" + "[" + meta.ID.Collection() + "]" + "[" + meta.OldRev + "]" + "[" + meta.Rev + "]")

	response := im.(models.Product)
	response.Pid = 1
	response.Name = meta.ID.String() + meta.ID.Key()
	response.Description = meta.ID.Collection() + meta.Rev + "old in persis	" + meta.OldRev
	response.Price = 10000
	return response, nil
}

func NewProductDepRepository() (IProductRepository, error) {
	logger.Debug().Msg("Creating Arango Client")
	dbc, err := clients.ArangoDBConnect(env.GetProperties().ArangoAddr, env.GetProperties().ArangoUser, env.GetProperties().ArangoPwd)

	if err != nil {
		logger.Error().Err(errors.New(err, "[🛑] Error connecting DB!"))
		return nil, errors.New(err, "Error connecting DB!")
	}
	logger.Debug().Msg("Arango Client Created")
	return &ProductRepositoryImpl{
		dbClient:                 dbc,
		databaseName:             env.GetProperties().ArangoDbName,
		deploymentCollectionName: env.GetProperties().ArangoCollectionName,
	}, nil
}

func (r *ProductRepositoryImpl) UpdateProduct(inpkey string, im *models.Product) (*models.Product, error) {
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

	meta, err := collection.ReplaceDocument(ctx, inpkey, im)

	if err != nil {
		logger.Error().Msgf("[🛑] failed to open collection on %s database " + r.databaseName)
		return im, errors.New(err, fmt.Sprintf("failed to open %s collection on %s database", r.deploymentCollectionName,
			r.databaseName))
	}
	// Persistent Store
	//meta, err := collection.CreateDocument(ctx, im)
	if err != nil {
		logger.Error().Msg("[🛑] failed updateproduct")
		return im, errors.New(err, "failed update")
	}
	logger.Debug().Msg("Persist Done meta [" + meta.ID.Key() + "]" + "[" + meta.ID.Collection() + "]" + "[" + meta.OldRev + "]" + "[" + meta.Rev + "]")

	response := im
	response.Pid = im.Pid
	response.Name = im.Name
	response.Description = im.Description + meta.ID.String() + meta.ID.Key() + meta.ID.Collection() + meta.Rev + "old in persist test	" + meta.OldRev
	response.Price = im.Price
	return response, nil
}

func (r *ProductRepositoryImpl) CreateProduct(im *models.Product) (*models.Product, error) {
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
	logger.Debug().Msg("Collection Open")

	logger.Debug().Msg("Persist Open")
	// Persistent Store
	meta, err := collection.CreateDocument(ctx, im)
	if err != nil {
		logger.Error().Msg("[🛑] failed to create processDeployment")
		return im, errors.New(err, "failed to create processDeployment")
	}
	logger.Debug().Msg("Persist Done meta [" + meta.ID.Key() + "]" + "[" + meta.ID.Collection() + "]" + "[" + meta.OldRev + "]" + "[" + meta.Rev + "]")

	response := im
	response.Pid = im.Pid
	response.Name = im.Name
	response.Description = im.Description + meta.ID.String() + meta.ID.Key() + meta.ID.Collection() + meta.Rev + "old in persist test	" + meta.OldRev
	response.Price = im.Price
	return response, nil
}

func (r *ProductRepositoryImpl) GetProduct() (interface{}, int64, error) {
	queryBuilder := query.NewForQuery(r.deploymentCollectionName, "doc")
	qs := queryBuilder.Filter("pid", ">=", 0).Done().Return().String()

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

	result := []models.Product{}
	if cursor.Count() == 0 || !cursor.HasMore() {
		return result, cursor.Count(), nil
	}

	for {
		process := models.Product{}
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

func (r *ProductRepositoryImpl) FindProduct(prm string, im *models.Product) (interface{}, error) {
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

	im.Name = flag

	response := im
	response.Pid = 1
	response.Name = flag
	response.Description = ""
	response.Price = 0
	return response, nil
}
