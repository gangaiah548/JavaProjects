package repository

import (
	"context"
	"fmt"
	"reflect"

	"sdk_workbench_authentication/src/clients"
	"sdk_workbench_authentication/src/config/env"
	"sdk_workbench_authentication/src/config/logger"
	"sdk_workbench_authentication/src/errors"
	"sdk_workbench_authentication/src/models"
	"sdk_workbench_authentication/src/repository/arango/query"

	arangodbdriver "github.com/arangodb/go-driver"
)

type IDomainRepository interface {
	CreateDomainAndSubDomain(im *[]models.SubDomain, dm *models.Domain) (*models.Domain, error)
	CreateDomain(dm *models.Domain) (*models.Domain, error)
	CreateComponent(dm *models.Entitlement) (*models.Entitlement, error)
	GetJavaClass() (interface{}, int64, error)
	GetDomains() (interface{}, int64, error)
	GetAllRoleData() ([]models.Entitlement, int64, error)
	DeleteJavaClass(key string, im interface{}) (interface{}, error)
	DeleteDomain(key string, im interface{}) (interface{}, error)
	UpdateDomain(inpkey string, im models.Domain) (models.Domain, error)
	FindDomain(prm string, im *models.Domain) (interface{}, error)
}

type DomainRepositoryImpl struct {
	dbClient                 arangodbdriver.Client
	databaseName             string
	deploymentCollectionName string
}

// DeleteJavaClass implements IJavaClassRepository.
func (r *DomainRepositoryImpl) DeleteJavaClass(_key string, im interface{}) (interface{}, error) {
	//var abc string  = key

	//fmt.Printf("queryBuilder %s", qs)
	//panic("unimplemented")
	/*FOR doc IN processDeploymentCollection
	  FILTER doc.uuid == '608906a5-c528-4663-83bf-cf1f5f374e9e'
	  remove doc IN processDeploymentCollection
	  RETURN OLD*/
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

	response := im.(models.Domain)
	//response.PackageName = _key
	//response.Type = meta.ID.String() + meta.ID.Key()
	return response, nil
}
func (r *DomainRepositoryImpl) DeleteDomain(uuid string, im interface{}) (interface{}, error) {
	ctx := context.Background()
	logger.Debug().Msg("Open DB")
	// Open database
	db, err := r.dbClient.Database(ctx, r.databaseName)
	if err != nil {
		logger.Error().Msg("[🛑] failed to connect DB %s " + r.databaseName + " not proper details")
		return im, errors.New(err, "failed to remove doc in collection please eneter proper details _key "+uuid+" not proper and collectionaname %s"+r.deploymentCollectionName)
	}

	logger.Info().Msgf("There are %s after Db %s connect", db.Name(), r.databaseName)

	collection, err := db.Collection(ctx, r.deploymentCollectionName)

	// Define the AQL query
	query := `
		FOR doc IN @@collection
		FILTER doc.uuid == '` + uuid + `'
		remove doc IN processDeploymentCollection
	  	RETURN OLD
	`
	//	"@attribute":  customAttribute,
	bindVars := map[string]interface{}{
		"@collection": collection.Name(),
	}
	ctx = context.Background()
	cursor, err := db.Query(ctx, query, bindVars)

	if err != nil {
		logger.Error().Msg("[🛑] failed to remove doc in collection please eneter proper details _key " + query + " not proper")
		return im, errors.New(err, "failed to remove doc in collection please eneter proper details "+query+" _key"+uuid+" not proper and collectionaname"+r.deploymentCollectionName)
	}
	defer cursor.Close()
	//logger.Debug().Msg("Persist Done meta" + "[" + meta.Collection() + "]" + "[" + meta.OldRev + "]" + "[" + meta.Rev + "]")

	response := im.(models.Domain)
	//response.Data = "Delete done [" + meta.ID.Key() + "]" + "[" + meta.ID.Collection() + "]" + "[" + meta.OldRev + "]" + "[" + meta.Rev + "]"
	//var key string
	for {
		process := models.Domain{}
		meta, err := cursor.ReadDocument(ctx, &process)
		if arangodbdriver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return response, err
		}
		// Process the deleted document if needed
		logger.Debug().Msg("[🛑] success _key " + meta.Key + " not proper")

		response = process
	}

	collectiont, err := db.Collection(ctx, "history_collection")
	if err != nil {
		logger.Error().Msgf("[🛑] failed to open collectiont on %s database " + r.databaseName)
		return im, errors.New(err, fmt.Sprintf("failed to open %s collection on %s database", "history_collection",
			r.databaseName))
	}
	backdatatohistory := models.Domaint{}
	backdatatohistory.Domain = response.Domain
	backdatatohistory.Description = response.Description
	backdatatohistory.CreatedBy = response.CreatedBy
	backdatatohistory.Date = response.Date
	backdatatohistory.SubDomains = response.SubDomains
	backdatatohistory.OpType = "D"
	_, err = collectiont.CreateDocument(ctx, backdatatohistory)
	if err != nil {
		logger.Fatal()
		return im, errors.New(err, fmt.Sprintf("failed to create document %s collection on %s database", "history_collection",
			r.databaseName))
	}

	return response, nil
}

func NewJavaClassDepRepository() (IDomainRepository, error) {
	logger.Debug().Msg("Creating Arango Client")
	dbc, err := clients.ArangoDBConnect(env.GetProperties().ArangoAddr, env.GetProperties().ArangoUser, env.GetProperties().ArangoPwd)

	if err != nil {
		logger.Error().Err(errors.New(err, "[🛑] Error connecting DB!"))
		return nil, errors.New(err, "Error connecting DB!")
	}
	logger.Debug().Msg("Arango Client Created")
	return &DomainRepositoryImpl{
		dbClient:                 dbc,
		databaseName:             env.GetProperties().ArangoDbName,
		deploymentCollectionName: env.GetProperties().ArangoCollectionName,
	}, nil
}

func (r *DomainRepositoryImpl) UpdateDomain(uuid string, im models.Domain) (models.Domain, error) {
	ctx := context.Background()
	logger.Debug().Msg("Open DB")
	// Open database
	//ctx := context.Background()
	/*FOR doc IN processDeploymentCollection
	  FILTER doc.uuid == '608906a5-c528-4663-83bf-cf1f5f374e9e'
	  UPDATE doc WITH { data: 'new_value' } IN processDeploymentCollection
	  RETURN NEW*/
	logger.Debug().Msg("Open DB")
	// Open database
	db, err := r.dbClient.Database(ctx, r.databaseName)
	if err != nil {
		logger.Error().Msgf("[🛑] failed to connect %s database " + r.databaseName)
		return im, errors.New(err, fmt.Sprintf("failed to connect %s DB on %s database", r.deploymentCollectionName,
			r.databaseName))
	}

	logger.Info().Msgf("There are %s after Db %s connect", db.Name(), r.databaseName)

	collection, err := db.Collection(ctx, r.deploymentCollectionName)
	if err != nil {
		logger.Error().Msgf("[🛑] failed to open collection on %s database " + r.databaseName)
		return im, errors.New(err, fmt.Sprintf("failed to open %s collection on %s database", r.deploymentCollectionName,
			r.databaseName))
	}

	query := `
		FOR doc IN @@collection
		FILTER doc.uuid == '` + uuid + `'
		RETURN doc._key
	`
	//	"@attribute":  customAttribute,
	bindVars := map[string]interface{}{
		"@collection": collection.Name(),
	}

	ctx = context.Background()
	cursor, err := db.Query(ctx, query, bindVars)

	if err != nil {
		logger.Error().Msg("[🛑] failed to remove doc in collection please eneter proper details _key " + query + " not proper")
		return im, errors.New(err, "failed to remove doc in collection please eneter proper details "+query+" _key"+uuid+" not proper and collectionaname"+r.deploymentCollectionName)
	}
	defer cursor.Close()
	//logger.Debug().Msg("Persist Done meta" + "[" + meta.Collection() + "]" + "[" + meta.OldRev + "]" + "[" + meta.Rev + "]")

	response := models.Domain{}
	var keyl string
	for {

		meta, err := cursor.ReadDocument(ctx, &keyl)
		if arangodbdriver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			logger.Error().Msg("[🛑] unable to find the data associate with given " + uuid + " document in collection")
			return im, errors.New(err, "unable to find the data associate with given "+uuid+" document in collection")
		}
		// Process the deleted document if needed
		logger.Debug().Msg("[🛑] success _key " + meta.Key + " not proper")

		//response = process
	}
	backdata := models.Domain{}
	_, err = collection.ReadDocument(ctx, keyl, &backdata)
	if err != nil {
		//logger.Fatal()
		logger.Error().Msg("[🛑] unable to find the document data associate with given " + uuid + " document in collection")
		return im, errors.New(err, "unable to find the document data associate with given "+uuid+" document in collection")

	}

	isEqual := reflect.DeepEqual(backdata, im)

	if isEqual {
		//response.Domain = "Data present in collection is same with sent data " + uuid
		logger.Error().Msg("[🛑] Data present in collection is same with sent data" + uuid + " no need any update")
		return im, errors.New(err, "Data present in collection is same with sent data"+uuid+" no need any update")

	}

	_, err = collection.ReplaceDocument(ctx, keyl, im)
	if err != nil {
		logger.Fatal()
		return response, err
	}

	collectiont, err := db.Collection(ctx, "history_collection")
	if err != nil {
		logger.Error().Msgf("[🛑] failed to open collectiont on %s database " + r.databaseName)
		return im, errors.New(err, fmt.Sprintf("failed to open %s collection on %s database", "history_collection",
			r.databaseName))
	}
	backdatatohistory := models.Domaint{}
	backdatatohistory.Domain = backdata.Domain
	backdatatohistory.Uuid = backdata.Uuid
	backdatatohistory.Description = backdata.Description
	backdatatohistory.CreatedBy = backdata.CreatedBy
	backdatatohistory.Date = backdata.Date
	backdatatohistory.SubDomains = backdata.SubDomains
	backdatatohistory.OpType = "U"
	_, err = collectiont.CreateDocument(ctx, backdatatohistory)
	if err != nil {
		logger.Fatal()
		return im, errors.New(err, fmt.Sprintf("failed to create document %s collection on %s database", "history_collection",
			r.databaseName))
	}

	//we  can feth the data from saved data also
	/*
				query := `
		        FOR doc IN your_collection_name
		        FILTER doc._key == @key
		        RETURN doc
		    `
		    bindVars := map[string]interface{}{
		        "key": newDocument.Key,
		    }

		    cursor, err := client.Query(ctx, query, bindVars)
		    if err != nil {
		        log.Fatal(err)
		    }

		    var updatedDocument MyDocument
		    for {
		        _, err := cursor.ReadDocument(ctx, &updatedDocument)
		        if driver.IsNoMoreDocuments(err) {
		            break
		        }
		        if err != nil {
		            log.Fatal(err)
		        }
		    }
	*/

	return im, nil
}

func (r *DomainRepositoryImpl) CreateComponent(dm *models.Entitlement) (*models.Entitlement, error) {
	ctx := context.Background()
	//ctx.cre
	logger.Debug().Msg("Open DB")
	// Open database
	db, err := r.dbClient.Database(ctx, r.databaseName)

	if err != nil {
		logger.Error().Msgf("[🛑] failed to open %s database ", r.databaseName)
		return dm, errors.New(err, fmt.Sprintf("failed to open %s database", r.databaseName))
	}

	logger.Debug().Msg("DB Open")
	logger.Debug().Msg("Open Collection")
	// Open collection

	collection, err := db.Collection(ctx, "entitlement")

	if err != nil {
		logger.Error().Msgf("[🛑] failed to open collection on %s database " + r.databaseName)
		return dm, errors.New(err, fmt.Sprintf("failed to open %s collection on %s database", r.deploymentCollectionName,
			r.databaseName))
	}

	//im.Uuid = key

	logger.Debug().Msg("Persist Open")
	// Persistent Store
	metad, derr := collection.CreateDocument(ctx, dm)
	if err != nil {
		logger.Error().Msg("[🛑] failed to create collection")
		return dm, errors.New(derr, "failed to create collection")
	}
	logger.Debug().Msg("Persist Done meta [" + metad.ID.Key() + "]" + "[" + metad.ID.Collection() + "]" + "[" + metad.OldRev + "]" + "[" + metad.Rev + "]")
	var response models.Entitlement
	_, err = collection.ReadDocument(ctx, metad.Key, &response)
	if err != nil {
		logger.Error().Msg("[🛑] failed to read collection" + err.Error())
	}

	//response.Data = metad.ID.Key() + "]" + "[" + metad.ID.Collection() + "]" + "[" + metad.OldRev + "]" + "[" + metad.Rev + "]"

	//response.Data = ""
	//response.ClassName = ""
	return &response, nil
}

func (r *DomainRepositoryImpl) CreateDomain(dm *models.Domain) (*models.Domain, error) {
	ctx := context.Background()
	//ctx.cre
	logger.Debug().Msg("Open DB")
	// Open database
	db, err := r.dbClient.Database(ctx, r.databaseName)

	if err != nil {
		logger.Error().Msgf("[🛑] failed to open %s database ", r.databaseName)
		return dm, errors.New(err, fmt.Sprintf("failed to open %s database", r.databaseName))
	}

	logger.Debug().Msg("DB Open")
	logger.Debug().Msg("Open Collection")
	// Open collection

	collection, err := db.Collection(ctx, r.deploymentCollectionName)

	if err != nil {
		logger.Error().Msgf("[🛑] failed to open collection on %s database " + r.databaseName)
		return dm, errors.New(err, fmt.Sprintf("failed to open %s collection on %s database", r.deploymentCollectionName,
			r.databaseName))
	}

	//im.Uuid = key

	logger.Debug().Msg("Persist Open")
	// Persistent Store
	metad, derr := collection.CreateDocument(ctx, dm)
	if err != nil {
		logger.Error().Msg("[🛑] failed to create collection")
		return dm, errors.New(derr, "failed to create collection")
	}
	logger.Debug().Msg("Persist Done meta [" + metad.ID.Key() + "]" + "[" + metad.ID.Collection() + "]" + "[" + metad.OldRev + "]" + "[" + metad.Rev + "]")
	var response models.Domain
	_, err = collection.ReadDocument(ctx, metad.Key, &response)
	if err != nil {
		logger.Error().Msg("[🛑] failed to read collection" + err.Error())
	}

	//response.Data = metad.ID.Key() + "]" + "[" + metad.ID.Collection() + "]" + "[" + metad.OldRev + "]" + "[" + metad.Rev + "]"

	//response.Data = ""
	//response.ClassName = ""
	return &response, nil
}

func (r *DomainRepositoryImpl) GetAllRoleData() ([]models.Entitlement, int64, error) {
	queryBuilder := query.NewForQuery("entitlement", "doc")
	qs := queryBuilder.Filter("document.entitlement", "NOT LIKE", "'%rrrr%'").Done().Return().String()
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
	var result []models.Entitlement
	//result := []models.Entitlement
	if cursor.Count() == 0 || !cursor.HasMore() {
		return result, cursor.Count(), nil
	}

	for {
		tEntitlement := models.Entitlement{}
		_, err = cursor.ReadDocument(pCtx, &tEntitlement)
		if arangodbdriver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, 0, errors.New(err, fmt.Sprintf("failed to read response from cursor for query [%s]", qs))
		}

		result = append(result, tEntitlement)
	}

	return result, cursor.Count(), nil
}

func (r *DomainRepositoryImpl) GetDomains() (interface{}, int64, error) {
	queryBuilder := query.NewForQuery(r.deploymentCollectionName, "doc")
	qs := queryBuilder.Filter("document.domain.data", "NOT LIKE", "'%rrrr%'").Done().Return().String()
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

	result := []models.Domain{}
	if cursor.Count() == 0 || !cursor.HasMore() {
		return result, cursor.Count(), nil
	}

	for {
		process := models.Domain{}
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

func (r *DomainRepositoryImpl) CreateDomainAndSubDomain(im *[]models.SubDomain, dm *models.Domain) (*models.Domain, error) {
	ctx := context.Background()
	logger.Debug().Msg("Open DB")
	// Open database
	db, err := r.dbClient.Database(ctx, r.databaseName)

	if err != nil {
		logger.Error().Msgf("[🛑] failed to open %s database ", r.databaseName)
		return dm, errors.New(err, fmt.Sprintf("failed to open %s database", r.databaseName))
	}

	logger.Debug().Msg("DB Open")
	logger.Debug().Msg("Open Collection")
	// Open collection

	collection, err := db.Collection(ctx, r.deploymentCollectionName)

	if err != nil {
		logger.Error().Msgf("[🛑] failed to open collection on %s database " + r.databaseName)
		return dm, errors.New(err, fmt.Sprintf("failed to open %s collection on %s database", r.deploymentCollectionName,
			r.databaseName))
	}

	//im.Uuid = key

	logger.Debug().Msg("Persist Open")
	// Persistent Store
	metad, derr := collection.CreateDocument(ctx, dm)
	if err != nil {
		logger.Error().Msg("[🛑] failed to create collection")
		return dm, errors.New(derr, "failed to create collection")
	}
	logger.Debug().Msg("Persist Done meta [" + metad.ID.Key() + "]" + "[" + metad.ID.Collection() + "]" + "[" + metad.OldRev + "]" + "[" + metad.Rev + "]")
	response := dm
	for _, n := range *im { //saving each document
		//key := uuid.New().String()
		//n.Uuid = key
		//fmt.Printf("Hello %s\n", n)
		meta, err := collection.CreateDocument(ctx, n)

		if err != nil {
			logger.Error().Msg("[🛑] failed to create collection")
			return dm, errors.New(err, "failed to create collection")
		}
		logger.Debug().Msg("Persist Done meta [" + meta.ID.Key() + "]" + "[" + meta.ID.Collection() + "]" + "[" + meta.OldRev + "]" + "[" + meta.Rev + "]")

		ecollection, err := db.Collection(ctx, "edgecollection")
		newEdge := models.KnowsEdge{
			EdgeDocument: arangodbdriver.EdgeDocument{
				From: metad.ID,
				To:   meta.ID,
			},
			Relation: "domain",
		}

		dmeta, derr := ecollection.CreateDocument(ctx, newEdge)
		if derr != nil {
			logger.Fatal()
		}
		fmt.Printf("KnowsEdge document created with ID '%s'\n", dmeta.Key)
		response.Domain = meta.ID.Key() + "]" + "[" + meta.ID.Collection() + "]" + "[" + meta.OldRev + "]" + "[" + meta.Rev + "]"

	}

	//response.Data = ""
	//response.ClassName = ""
	return response, nil
}

func (r *DomainRepositoryImpl) GetJavaClass() (interface{}, int64, error) {
	//queryBuilder := query.NewForQuery(r.deploymentCollectionName, "doc")

	//queryin:=queryBuilder.NewForQuery("edgecollection", "edge")
	//qs := queryBuilder.Filter("document.procedures.node", "NOT LIKE", "'%rrrr%'").Done().Return().String()
	//FOR document IN processDeploymentCollection
	// FILTER NOT LIKE(document.procedures.node, '%hemogram%') RETURN document
	//fmt.Printf("queryBuilder %s", qs)

	pCtx := context.Background()
	//ctx := arangodbdriver.WithQueryCount(pCtx)

	// Open database
	db, err := r.dbClient.Database(pCtx, r.databaseName)

	logger.Info().Msgf("There are %s after Db %s connect", db.Name(), r.databaseName)

	if err != nil {
		logger.Warn().Msgf("[⚠️] failed to open %s database.", r.databaseName)
		return nil, 0, errors.New(err, fmt.Sprintf("failed to open [%s] database", r.databaseName))
	}

	result := []models.Domain{}
	var count int64
	result, count, err = fetchDocumentsOfTypeDomain(db, r.deploymentCollectionName)

	if err != nil {
		logger.Warn().Msgf("[⚠️] failed to query [%s] on %s database", "qs2", r.databaseName)
		return nil, 0, errors.New(err, fmt.Sprintf("failed to query [%s] on %s database", "qs2", r.databaseName))
	}
	result2 := []models.SubDomain{}
	//var count2 int64
	//err2 :=
	result2, count, err = fetchDocumentsOfTypeJavaClass(db, r.deploymentCollectionName)

	if err != nil {
		logger.Warn().Msgf("[⚠️] failed to query [%s] on %s database", result2, r.databaseName)
		return nil, 0, errors.New(err, fmt.Sprintf("failed to query [%s] on %s database", "", r.databaseName))
	}
	//rese := []models.DomainSubDomain{}
	var rese []interface{}
	for _, m := range result {
		/*switch v := m.(type) {
		case models.Domain:*/
		if len(m.Uuid) == 0 || m.Uuid == "" || m.Uuid == "null" {
			continue
		} else {
			rese = append(rese, m)
		}
		/*case models.JavaClass:
			continue
		default :
		continue
		}*/
	}
	for _, m := range result2 {
		if len(m.Subdomain) == 0 || m.Subdomain == "" || m.Subdomain == "null" {
			continue
		} else {
			rese = append(rese, m)
		}
	}
	//rese = append(rese, result, result2)
	return rese, count, nil

	/*var qs2 string
	qs2 = "FOR doc IN @@collection LET subquery = (FOR edge IN edgecollection FILTER edge._from == doc._id || edge._to == doc._id return edge._from) RETURN doc"
	//cursor, err := db.Query(nil, query, bindVars)
	bindVars := map[string]interface{}{
		"@collection": "processDeploymentCollection",
	}
	cursor, err := db.Query(ctx, qs2, bindVars)

	if err != nil {
		logger.Warn().Msgf("[⚠️] failed to query [%s] on %s database", qs2, r.databaseName)
		return nil, 0, errors.New(err, fmt.Sprintf("failed to query [%s] on %s database", qs2, r.databaseName))
	}
	defer cursor.Close()

	result := []models.Domaint{}
	if cursor.Count() == 0 || !cursor.HasMore() {
		return result, cursor.Count(), nil
	}

	for {
		process := models.Domaint{}
		_, err = cursor.ReadDocument(pCtx, &process)
		if arangodbdriver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, 0, errors.New(err, fmt.Sprintf("failed to read response from cursor for query [%s]", qs2))
		}

		result = append(result, process)
	}*/

}

func (r *DomainRepositoryImpl) FindDomain(prm string, im *models.Domain) (interface{}, error) {
	ctx := context.Background()
	db, err := r.dbClient.Database(ctx, r.databaseName)

	//db, err := r.dbClient.Database(pCtx, r.databaseName)

	logger.Info().Msgf("There are %s after Db %s connect", db.Name(), r.databaseName)

	collection, err := db.Collection(ctx, r.deploymentCollectionName)

	meta, err := collection.DocumentExists(ctx, prm) //readdocument also we can use

	if err != nil {
		logger.Error().Msgf("[🛑] failed to find document in  %s database " + r.databaseName)
		return im, errors.New(err, fmt.Sprintf("failed to find document in %s collection on %s database", r.deploymentCollectionName,
			r.databaseName))
	}
	//	var flag string = ""
	logger.Debug().Msgf("search is Done meta %t", meta)
	response := im
	if meta == true {
		//flag = "1"
		response.Domain = "1"
		response.Uuid = "flag"
		return response, nil
	} else {
		response.Domain = "0"
		response.Uuid = "1"
		return response, nil
	}

	//im.Data = flag

}

func fetchDocumentsOfTypeJavaClass(db arangodbdriver.Database, collectionName string) ([]models.SubDomain, int64, error) {
	var modelType1Docs []models.SubDomain

	query := "FOR doc IN @@collection LET subquery = (FOR edge IN edgecollection FILTER edge._from == doc._id || edge._to == doc._id return edge._from) RETURN doc"

	bindVars := map[string]interface{}{
		"@collection": collectionName,
	}

	cursor, err := db.Query(nil, query, bindVars)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close()

	for {
		var doc models.SubDomain
		_, err := cursor.ReadDocument(nil, &doc)
		if arangodbdriver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, 0, err
		}

		modelType1Docs = append(modelType1Docs, doc)
	}

	return modelType1Docs, cursor.Count(), nil
}

// Fetch documents of ModelType2 from the collection
func fetchDocumentsOfTypeDomain(db arangodbdriver.Database, collectionName string) ([]models.Domain, int64, error) {
	var modelType2Docs []models.Domain

	query := "FOR doc IN @@collection LET subquery = (FOR edge IN edgecollection FILTER edge._from == doc._id || edge._to == doc._id return edge._from) RETURN doc"
	bindVars := map[string]interface{}{
		"@collection": collectionName,
	}

	cursor, err := db.Query(nil, query, bindVars)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close()

	for {
		var doc models.Domain
		_, err := cursor.ReadDocument(nil, &doc)
		if arangodbdriver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, 0, err
		}

		modelType2Docs = append(modelType2Docs, doc)
	}

	return modelType2Docs, cursor.Count(), err
}

func fetchDomainDetails(db arangodbdriver.Database, uuid string, collectionName string) (key string) {

	// Specify the custom attribute and its value for the query
	//customAttribute := "uuid"
	//customAttributeValue := ""

	// Define the AQL query
	query := `
		FOR doc IN @@collection
		FILTER doc.uuid == '` + uuid + `'
		RETURN doc
	`

	bindVars := map[string]interface{}{
		"@collection": collectionName,
	}

	// Execute the AQL query
	ctx := context.Background()
	cursor, err := db.Query(ctx, query, bindVars)
	if err != nil {
		logger.Error().Msgf("[🛑] failed query to find document in  %s database " + query)
	}
	defer cursor.Close()

	// Process the query result
	for {
		process := models.Domain{}
		_, err := cursor.ReadDocument(ctx, &process)
		if arangodbdriver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			logger.Error().Msgf("[🛑] failed to find document in  %s database " + collectionName)
		}

		fmt.Printf("Found Document: %+v\n", key)
	}

	return key

}
