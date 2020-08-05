package repository

import (
	"fmt"
    "go.mongodb.org/mongo-driver/bson"
	"github.com/myrachanto/asokomonolith/httperrors"
	"github.com/myrachanto/asokomonolith/model" 
)
var (
	Productrepository productrepository = productrepository{}
)

type productrepository struct{}

func (r *productrepository) Create(product *model.Product) (*httperrors.HttpError) {
	if err1 := product.Validate(); err1 != nil {
		return err1
	}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("product")
	_, err := collection.InsertOne(ctx, product)
		if err != nil {
			return httperrors.NewBadRequestError(fmt.Sprintf("Create product Failed, %d", err))
	}
	DbClose(c)
	return nil
}

func (r *productrepository) GetOne(id string) (product *model.Product, errors *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("product")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	DbClose(c)
	return product, nil	
}

func (r *productrepository) GetAll(products []model.Product) ([]model.Product, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("product")
	filter := bson.M{}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	if err != nil { 
		return nil,	httperrors.NewNotFoundError("no results found")
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
	err := cur.Decode(&products)
		if err != nil { 
			return nil,	httperrors.NewNotFoundError("Error while decoding results!")
		}
	// do something with result....
	}
	if err := cur.Err(); err != nil {
		return nil,	httperrors.NewNotFoundError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}	
	DbClose(c)
    return products, nil

}

func (r *productrepository) Update(id string, product *model.Product) (*httperrors.HttpError) {
	uproduct := &model.Product{}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("product")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&uproduct)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	if product.Name  == "" {
		product.Name = uproduct.Name
	}
	if product.Title  == "" {
		product.Title = uproduct.Title
	}
	if product.Description  == "" {
		product.Description = uproduct.Description
	}
	if product.Majorcategory  == "" {
		product.Majorcategory = uproduct.Majorcategory
	}
	if product.Category  == "" {
		product.Category = uproduct.Category
	}
	if product.Subcategory  == "" {
		product.Subcategory = uproduct.Subcategory
	}
	if product.Price  > 0 {
		product.Price = uproduct.Price
	}
	if len(product.Tag)  > 0  {
		product.Tag = uproduct.Tag
	}
	if len(product.Rates)  > 0  {
		product.Rates = uproduct.Rates
	}
	_, err = collection.UpdateOne(ctx, filter, product)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Update of product Failed, %d", err))
	} 
	DbClose(c)
	return nil
}
func (r productrepository) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("product")
	filter := bson.M{"_id": id}
	ok, err := collection.DeleteOne(ctx, filter)
	if ok == nil {
		return nil, httperrors.NewNotFoundError(fmt.Sprintf("deletion of %d failed", err))
	}
	DbClose(c)
		return httperrors.NewSuccessMessage("deleted successfully"), nil
}
