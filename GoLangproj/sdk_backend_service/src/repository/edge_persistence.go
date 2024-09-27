package repository

import (
	"context"
	"fmt"
	"log"
	"sdk_backend_service/src/clients"
	"sdk_backend_service/src/config/env"
	"sdk_backend_service/src/models"

	driver "github.com/arangodb/go-driver"
)

func dbEdgeConn() (driver.Database, error) {
	/*conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
	})
	if err != nil {
		log.Fatal(err)
	}*/

	dbc, err := clients.ArangoDBConnect(env.GetProperties().ArangoAddr, env.GetProperties().ArangoUser, env.GetProperties().ArangoPwd)

	/*client, err := driver.NewClient(driver.ClientConfig{
		Connection: conn,
	})
	if err != nil {
		log.Fatal(err)
	}*/

	// Database and collection
	ctx := context.Background()
	db, err := dbc.Database(ctx, "processDeployment")
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
func createEdge() {
	ctx := context.Background()
	db, err := dbEdgeConn()
	edgeCollection, err := db.Collection(ctx, "edgecollection")
	if err != nil {
		log.Fatal(err)
	}

	// Create a KnowsEdge
	newEdge := models.KnowsEdge{
		Relation: "Acquaintance",
	}

	// CreateDocument: Create a new edge
	meta, err := edgeCollection.CreateDocument(ctx, newEdge)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("KnowsEdge document created with ID '%s'\n", meta.Key)
}

func readEdge() {
	// ReadDocument: Retrieve the created edge
	ctx := context.Background()
	db, err := dbEdgeConn()
	edgeCollection, err := db.Collection(ctx, "edgecollection")
	var retrievedEdge models.KnowsEdge
	_, err = edgeCollection.ReadDocument(ctx, "meta.Key", &retrievedEdge)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Retrieved KnowsEdge document: %+v\n", retrievedEdge)
}
func updateEdge() {
	// UpdateDocument: Update the edge's relation
	ctx := context.Background()
	db, err := dbEdgeConn()
	edgeCollection, err := db.Collection(ctx, "edgecollection")
	var retrievedEdge models.KnowsEdge
	//_, err = edgeCollection.ReadDocument(ctx, meta.Key, &retrievedEdge)
	retrievedEdge.Relation = "Friend"
	_, err = edgeCollection.UpdateDocument(ctx, "meta.Key", retrievedEdge)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("KnowsEdge document updated")
}

// DeleteDocument: Delete the edge
func edgeDelete() {
	ctx := context.Background()
	db, err := dbEdgeConn()
	edgeCollection, err := db.Collection(ctx, "edgecollection")
	//var retrievedEdge models.KnowsEdge
	_, err = edgeCollection.RemoveDocument(ctx, "meta.Key")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("KnowsEdge document deleted")
}
