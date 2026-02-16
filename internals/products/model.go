package products

import "go.mongodb.org/mongo-driver/v2/bson"

type Product struct {
	ID          bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Price       float64 `json:"price" bson:"price"`
}
