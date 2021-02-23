package repository

import (
	"fmt"
	"strconv"
    "go.mongodb.org/mongo-driver/bson"
		"github.com/myrachanto/ecommerce/httperrors"
		"github.com/myrachanto/ecommerce/model"  
)
//productrepository ...
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
	code, err1 := Productrepository.genecode()
	if err1 != nil {
		return err1
	}
	product.Code = code
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

func (r *productrepository) GetAll(code string) ([]*model.Product, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	products := []*model.Product{}
	collection := db.Collection("product")
	filter := bson.M{"category": code}
	fmt.Println(filter)
	cur, err := collection.Find(ctx, filter)
	if err != nil { 
		return nil,	httperrors.NewNotFoundError("no results found")
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var product model.Product
		err := cur.Decode(&product)
			if err != nil { 
				return nil,	httperrors.NewNotFoundError("Error while decoding results!")
			}
	 products = append(products, &product)
	}
	if err := cur.Err(); err != nil {
		return nil,	httperrors.NewNotFoundError("Error with cursor!")
	}	
	DbClose(c)
    return products, nil

}

func (r *productrepository) Update(code string, product *model.Product) (*httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	result, err3 := Productrepository.getuno(code)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(result)
	if product.Name == ""{
		product.Name = result.Name
	}
	if product.Title == ""{
		product.Title = result.Title
	}
	if product.Description == ""{
		product.Description = result.Description
	}
	if product.Code == ""{
		product.Code = result.Code
	}
	collection := db.Collection("product")
	filter := bson.M{"code": code}
	fmt.Println(filter)
	fmt.Println(product)
	update := bson.M{"$set": product}
	_, err := collection.UpdateOne(ctx, filter, update)
		if err != nil {
		return	httperrors.NewNotFoundError("Error updating!")
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
func (r productrepository)genecode()(string, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return "", t
	}
	db, e := Mongodb();if e != nil {
		return "", e
	}
	collection := db.Collection("product")
	filter := bson.M{}
	count, err := collection.CountDocuments(ctx, filter)
	co := count + 1
	if err != nil { 
		return "",	httperrors.NewNotFoundError("no results found")
	}
	code := "ProductCode"+strconv.FormatUint(uint64(co), 10)

	DbClose(c)
	return code, nil
}
func (r productrepository)getuno(code string)(result *model.Product, err *httperrors.HttpError){
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("product")
	filter := bson.M{"code": code}
	err1 := collection.FindOne(ctx, filter).Decode(&result)
	if err1 != nil {
		return nil, httperrors.NewNotFoundError("no results found")
	}
	DbClose(c)
	return result, nil	
}
