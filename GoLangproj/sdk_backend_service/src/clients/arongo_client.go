package clients

import (
	"context"
	"fmt"

	arangodbdriver "github.com/arangodb/go-driver"
	arangodbdriverhttp "github.com/arangodb/go-driver/http"

	"sdk_backend_service/src/config/logger"
	"sdk_backend_service/src/errors"
)

func ArangoDBConnect(address string, user string, password string) (arangodbdriver.Client, error) {
	// TODO Add TLS configuration handling
	conn, err := arangodbdriverhttp.NewConnection(arangodbdriverhttp.ConnectionConfig{
		Endpoints: []string{address},
	})
	if err != nil {
		logger.Error().Msg("[🛑] failed create connection with the database")
		return nil, errors.New(err, fmt.Sprintf("failed to connect %s database", address))
	}
	// TODO Add JWT configuration handling
	client, err := arangodbdriver.NewClient(arangodbdriver.ClientConfig{
		Connection:     conn,
		Authentication: arangodbdriver.BasicAuthentication(user, password),
	})
	if err != nil {
		logger.Error().Msg("[🛑] failed create Arango client")
		return nil, errors.New(err, fmt.Sprintf("failed to connect %s database", address))
	}

	return client, nil
}

func CreateOrUpdateDB(dbc arangodbdriver.Client, databaseName string, collectionName string) {

	ctx := context.Background()

	// check db
	var db arangodbdriver.Database
	//var collection arangodbdriver.Collection
	var err error
	_, err = dbc.Database(ctx, databaseName)
	if err != nil {
		logger.Warn().Msgf("[⚠️] failed to open %s database. Creating now... ", databaseName)
		db, err = dbc.CreateDatabase(ctx, databaseName, nil)
		if err != nil {
			logger.Error().Msg(fmt.Sprintf(err.Error()))
			logger.Fatal().Msg("[🛑] failed create Database")
		}
		// check collection
		_, err = db.Collection(ctx, collectionName)
		if err != nil {
			logger.Warn().Msgf("[⚠️] failed to open %s collection on %s database. Creating now... ", collectionName, databaseName)
			_, err = db.CreateCollection(ctx, collectionName, nil)
			if err != nil {
				logger.Fatal().Msgf("Failed to create collection: %s", collectionName)
			}
		}
	}
	logger.Info().Msgf("[✅] Collection %s on DB %s is available", collectionName, databaseName)
}
