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

type ShoppingItemHandler struct {
	collection *mongo.Collection
}

func NewShoppingItemHandler(collection *mongo.Collection) *ShoppingItemHandler {
	return &ShoppingItemHandler{collection: collection}
}

// Get all items in a shopping list
func (handler *ShoppingItemHandler) GetAllItemsInShoppingList(c echo.Context) error {
	var items []models.ShoppingItem

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid shopping list ID")
	}

	// Find all ShoppingItem objects in the database that belong to the specified ShoppingList
	cursor, err := handler.collection.Find(
		context.Background(),
		bson.M{"shopping_list_id": id},
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &items)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError,
			"Unable to fetch items in the shopping list")
	}

	if len(items) == 0 {
		return c.JSON(http.StatusOK, []models.ShoppingItem{})
	}

	return c.JSON(http.StatusOK, items)
}

// Create a new ShoppingItem object
func (handler *ShoppingItemHandler) Create(c echo.Context) error {
	item := new(models.ShoppingItem)
	if err := c.Bind(item); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Unable to bind request body")
	}

	result, err := handler.collection.InsertOne(context.Background(), item)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	item.ID = result.InsertedID.(primitive.ObjectID)
	return c.JSON(http.StatusCreated, item)
}

// Mark a shopping item as completed
func (handler *ShoppingItemHandler) Complete(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid item ID")
	}

	update := bson.M{
		"$set": bson.M{"completed": true},
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
		return echo.NewHTTPError(http.StatusNotFound, "Item not found")
	}

	return c.NoContent(http.StatusOK)
}
