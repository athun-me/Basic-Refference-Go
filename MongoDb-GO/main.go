package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Item represents an example item structure
type Item struct {
	ID    string `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Price int    `json:"price,omitempty" bson:"price,omitempty"`
}

var client *mongo.Client
var itemCollection *mongo.Collection

func init() {
	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Could not connect to the MongoDB server:", err)
	}

	// Access the collection
	itemCollection = client.Database("testdb").Collection("items")
}

// CreateItem adds a new item to the database
func CreateItem(c *gin.Context) {
	var newItem Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert the new item into the collection
	result, err := itemCollection.InsertOne(context.Background(), newItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newItem.ID = result.InsertedID.(string)
	c.JSON(http.StatusCreated, newItem)
}

// GetItem retrieves a single item by ID
func GetItem(c *gin.Context) {
	id := c.Param("id")

	var item Item
	err := itemCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&item)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

// UpdateItem updates an existing item by ID
func UpdateItem(c *gin.Context) {
	id := c.Param("id")

	var updatedItem Item
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := itemCollection.ReplaceOne(context.Background(), bson.M{"_id": id}, updatedItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, updatedItem)
}

// DeleteItem removes an item by ID
func DeleteItem(c *gin.Context) {
	id := c.Param("id")

	result, err := itemCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}

func main() {
	r := gin.Default()

	r.POST("/items", CreateItem)
	r.GET("/items/:id", GetItem)
	r.PUT("/items/:id", UpdateItem)
	r.DELETE("/items/:id", DeleteItem)

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(r.Run(":8080"))
}
