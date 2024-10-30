package main

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/kdkanishka/shoppinglist-service/handlers"
	"github.com/labstack/echo"
)

func main() {
	// Read the configuration via environment variables
	mongoURI := os.Getenv("MONGO_URI")
	log.Printf("Mongo URI: %s", mongoURI)

	// Setup the db connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Create collections
	db := client.Database("shoppinglist")
	shoppingListCollection := db.Collection("shoppinglists")
	shoppingItemsCollection := db.Collection("shoppingitems")

	// Create handlers
	shoppingListHandler := handlers.NewShoppingListHandler(shoppingListCollection)
	shoppingItemHandler := handlers.NewShoppingItemHandler(shoppingItemsCollection)

	// Create the echo instance
	e := echo.New()

	// shopping list routes
	e.POST("/shoppinglists", shoppingListHandler.Create)
	e.GET("/shoppinglists", shoppingListHandler.GetAll)
	e.PUT("/shoppinglists/:id", shoppingListHandler.Update)

	// shopping item routes
	e.POST("/shoppingitems", shoppingItemHandler.Create)
	e.GET("/shoppinglists/:id/shoppingitems", shoppingItemHandler.GetAllItemsInShoppingList)
	e.PUT("/shoppingitems/:id/complete", shoppingItemHandler.Complete)

	e.Logger.Fatal(e.Start(":8080"))

}
