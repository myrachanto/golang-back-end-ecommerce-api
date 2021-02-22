package repository

import (
	"fmt"
    "go.mongodb.org/mongo-driver/bson"
		"github.com/myrachanto/ecommerce/httperrors"
		"github.com/myrachanto/ecommerce/model" 
)
var (
	Streetrepository streetrepository = streetrepository{}
)

type streetrepository struct{}

func (r *streetrepository) Create(street *model.Street) (*httperrors.HttpError) {
	if err1 := street.Validate(); err1 != nil {
		return err1
	}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("street")
	_, err := collection.InsertOne(ctx, street)
		if err != nil {
			return httperrors.NewBadRequestError(fmt.Sprintf("Create street Failed, %d", err))
	}
	DbClose(c)
	return nil
}

func (r *streetrepository) GetOne(id string) (street *model.Street, errors *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("street")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&street)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	DbClose(c)
	return street, nil	
}

func (r *streetrepository) GetAll(streets []model.Street) ([]model.Street, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("street")
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
	err := cur.Decode(&streets)
		if err != nil { 
			return nil,	httperrors.NewNotFoundError("Error while decoding results!")
		}
	// do something with result....
	}
	if err := cur.Err(); err != nil {
		return nil,	httperrors.NewNotFoundError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}	
	DbClose(c)
    return streets, nil

}

func (r *streetrepository) Update(id string, street *model.Street) (*httperrors.HttpError) {
	ustreet := &model.Street{}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("street")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&ustreet)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	if street.Name  == "" {
		street.Name = ustreet.Name
	}
	if street.Title  == "" {
		street.Title = ustreet.Title
	}

	if street.Description  == "" {
		street.Description = ustreet.Description
	}
	_, err = collection.UpdateOne(ctx, filter, street)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Update of street Failed, %d", err))
	} 
	DbClose(c)
	return nil
}
func (r streetrepository) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("street")
	filter := bson.M{"_id": id}
	ok, err := collection.DeleteOne(ctx, filter)
	if ok == nil {
		return nil, httperrors.NewNotFoundError(fmt.Sprintf("deletion of %d failed", err))
	}
	DbClose(c)
		return httperrors.NewSuccessMessage("deleted successfully"), nil
}

