// api/api.go
package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"refactor/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getShipmentsHandler(c *gin.Context) {
	placementsDB, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer placementsDB.Disconnect(context.Background())
	collection := placementsDB.Database("Raaho").Collection("shipment")

	var shipments []model.Shipment
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error retrieving shipments"})
		return
	}
	fmt.Println(cursor, err)

	if err := cursor.All(context.Background(), &shipments); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error decoding shipments"})
		return
	}
	c.JSON(http.StatusOK, shipments)
}

func createShipmentHandler(c *gin.Context) {
	placementsDB, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer placementsDB.Disconnect(context.Background())
	collection := placementsDB.Database("Raaho").Collection("shipment")

	var shipment model.Shipment
	err = c.BindJSON(&shipment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shipment.ID = primitive.NewObjectID()
	// Add timestamps
	currentTime := time.Now().UTC()
	shipment.CreatedAt = currentTime
	shipment.UpdatedAt = currentTime

	_, err = collection.InsertOne(context.Background(), shipment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error inserting shipment"})
		return
	}

	c.JSON(http.StatusCreated, shipment)
}

func StartServer() {
	r := gin.Default()
	// Define API routes and handlers
	r.GET("/shipments", getShipmentsHandler)
	r.POST("/shipments", createShipmentHandler)

	if err := r.Run(":8080"); err != nil {
		// Handle errors
		panic(err)
	}
}
