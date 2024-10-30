package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ShoppingList struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
}

type ShoppingItem struct {
	ID             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ShoppingListID primitive.ObjectID `json:"shopping_list_id" bson:"shopping_list_id"`
	Name           string             `json:"name" bson:"name"`
	Quantity       int                `json:"quantity" bson:"quantity"`
	Completed      bool               `json:"completed" bson:"completed"`
}

type ShoppingListWithItems struct {
	ShoppingList
	Items []ShoppingItem `json:"items"`
}
