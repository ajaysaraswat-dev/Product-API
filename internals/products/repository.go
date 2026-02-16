package products

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository interface {
	ListProducts(ctx context.Context)([]Product,error)
	GetProduct(ctx context.Context,id string)(*Product,error)
	CreateProduct(ctx context.Context,product *Product) error
	UpdateProduct(ctx context.Context,id string,product *Product)error
	DeleteProduct(ctx context.Context,id string) error
	
}

type repo struct {
	collection *mongo.Collection
}
func NewRepository(db *mongo.Database) Repository {
	return &repo{
		collection: db.Collection("products"),
	}
}

func (r *repo) ListProducts(ctx context.Context)([]Product,error){
	//call the database to get the list of products and return it to the service layer
	var products []Product
	//mongodb doesnot return the list of products directly it return a cursor which we can iterate to get the list of products
	cursor,err := r.collection.Find(ctx, bson.M{}) //find take two arguments - context and filter. We want to get all the products from the database, so we pass an empty filter.
	if err != nil {
		return nil,err
	}
	if err := cursor.All(ctx,&products); err != nil {
		return nil,err
	}
	return products,nil
}

func(r *repo) GetProduct(ctx context.Context,id string)(*Product,error){
	objId,err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("invalid product id: %s",id)
		return nil,err
	}
	var product Product
	if err := r.collection.FindOne(ctx,bson.M{"_id":objId}).Decode(&product);err != nil {
		log.Printf("failed to get product with id %s: %v",id,err)
		return nil,err
	}
	return &product,nil
}

func (r *repo) CreateProduct(ctx context.Context,product *Product) error {
	_,err := r.collection.InsertOne(ctx,product)
	if err != nil {;
		log.Printf("failed to create product: %v",err)
		return err
	}
	return nil
}

func (r *repo) UpdateProduct(ctx context.Context,id string,product *Product)error{
	objId,err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("failed to get the id: %v",err)
		return err
	}
	update := bson.M{
		"$set":bson.M{
			"name":product.Name,
			"description":product.Description,
			"price":product.Price,
		},
	}
	if _,err := r.collection.UpdateByID(ctx,objId,update); err !=nil {
		log.Printf("failed to update product with id %s: %v",id,err)
		return err
	}
	return nil
}

func(r *repo)DeleteProduct(ctx context.Context,id string) error {
	objId,err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("failed to get the id: %v",err)
		return err
	}
	result,err := r.collection.DeleteOne(ctx,bson.M{"_id": objId})
	if err != nil {
		log.Printf("Failed to delete product with product id %s:%v",id,err)
		return err
	}
	if result.DeletedCount == 0{
		log.Printf("No product found with this id for delete operation: %s",id)
		return nil
	}
	return nil
}