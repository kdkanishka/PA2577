package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/kdkanishka/shoppinglist-service/models"
)

type ShoppingListHandler struct {
	collection *mongo.Collection
}

func NewShoppingListHandler(collection *mongo.Collection) *ShoppingListHandler {
	return &ShoppingListHandler{collection: collection}
}

// Create a new ShoppingList object
func (handler *ShoppingListHandler) Create(c echo.Context) error {
	list := new(models.ShoppingList) // Create a new ShoppingList object to hold the request body
	if err := c.Bind(list); err != nil {
		//bad request if the body can't be bound to the ShoppingList object
		return echo.NewHTTPError(http.StatusBadRequest, "Unable to bind request body")
	}

	// Insert the ShoppingList object into the database
	result, err := handler.collection.InsertOne(context.Background(), list)
	if err != nil {
		// Internal server error if the database operation fails
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Return newly created ShoppingList object including the generated ID
	list.ID = result.InsertedID.(primitive.ObjectID)
	return c.JSON(http.StatusCreated, list)
}

// Get all ShoppingList objects
func (handler *ShoppingListHandler) GetAll(c echo.Context) error {
	var lists []models.ShoppingList

	// Find all ShoppingList objects in the database
	cursor, err := handler.collection.Find(context.Background(), bson.M{})
	if err != nil {
		// Internal server error if the database operation fails
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer cursor.Close(context.Background())

	// Iterate over the cursor and decode each document into a ShoppingList object
	for cursor.Next(context.Background()) {
		var list models.ShoppingList
		if err := cursor.Decode(&list); err != nil {
			// Internal server error if decoding fails
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		lists = append(lists, list)
	}

	if len(lists) == 0 {
		return c.JSON(http.StatusNotFound, []models.ShoppingList{})
	}

	return c.JSON(http.StatusOK, lists)
}

// Update a ShoppingList object
func (handler *ShoppingListHandler) Update(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		// Bad request if the ID is not a valid ObjectID
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	shoppingList := new(models.ShoppingList) // Create a new ShoppingList object to hold the request body
	if err := c.Bind(shoppingList); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	update := bson.M{
		"$set": bson.M{
			"name":        shoppingList.Name,
			"description": shoppingList.Description,
		},
	}

	result, err := handler.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		update,
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if result.MatchedCount == 0 {
		return echo.NewHTTPError(http.StatusNotFound,
			"Unable to update, Shopping list not found")
	}

	return c.JSON(http.StatusOK, shoppingList)

}
