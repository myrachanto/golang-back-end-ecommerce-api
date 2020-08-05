package repository

import (
	"fmt"
    "go.mongodb.org/mongo-driver/bson"
	"github.com/myrachanto/asokomonolith/httperrors"
	"github.com/myrachanto/asokomonolith/model" 
)
var (
	Customerrepository customerrepository = customerrepository{}
)

type customerrepository struct{}

func (r *customerrepository) Create(customer *model.Customer) (*httperrors.HttpError) {
	if err1 := customer.Validate(); err1 != nil {
		return err1
	}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("customer")
	_, err := collection.InsertOne(ctx, customer)
		if err != nil {
			return httperrors.NewBadRequestError(fmt.Sprintf("Create customer Failed, %d", err))
	}
	DbClose(c)
	return nil
}

func (r *customerrepository) GetOne(id string) (customer *model.Customer, errors *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("customer")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&customer)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	DbClose(c)
	return customer, nil	
}

func (r *customerrepository) GetAll(customers []model.Customer) ([]model.Customer, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("customer")
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
	err := cur.Decode(&customers)
		if err != nil { 
			return nil,	httperrors.NewNotFoundError("Error while decoding results!")
		}
	// do something with result....
	}
	if err := cur.Err(); err != nil {
		return nil,	httperrors.NewNotFoundError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}	
	DbClose(c)
    return customers, nil

}

func (r *customerrepository) Update(id string, customer *model.Customer) (*httperrors.HttpError) {
	ucustomer := &model.Customer{}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("customer")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&ucustomer)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	if customer.Name  == "" {
		customer.Name = ucustomer.Name
	}
	if customer.Company  == "" {
		customer.Company = ucustomer.Company
	}
	if customer.Phone  == "" {
		customer.Phone = ucustomer.Phone
	}
	if customer.Email  == "" {
		customer.Email = ucustomer.Email
	}
	if customer.Address  == "" {
		customer.Address = ucustomer.Address
	}
	_, err = collection.UpdateOne(ctx, filter, customer)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Update of customer Failed, %d", err))
	} 
	DbClose(c)
	return nil
}
func (r customerrepository) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("customer")
	filter := bson.M{"_id": id}
	ok, err := collection.DeleteOne(ctx, filter)
	if ok == nil {
		return nil, httperrors.NewNotFoundError(fmt.Sprintf("deletion of %d failed", err))
	}
	DbClose(c)
		return httperrors.NewSuccessMessage("deleted successfully"), nil
}
