package repository

import (
	"fmt"
    "go.mongodb.org/mongo-driver/bson"
	"github.com/myrachanto/asokomonolith/httperrors"
	"github.com/myrachanto/asokomonolith/model" 
)
var (
	Nortificationrepository nortificationrepository = nortificationrepository{}
)

type nortificationrepository struct{}

func (r *nortificationrepository) Create(nortification *model.Nortification) (*httperrors.HttpError) {
	if err1 := nortification.Validate(); err1 != nil {
		return err1
	}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("nortification")
	_, err := collection.InsertOne(ctx, nortification)
		if err != nil {
			return httperrors.NewBadRequestError(fmt.Sprintf("Create nortification Failed, %d", err))
	}
	DbClose(c)
	return nil
}

func (r *nortificationrepository) GetOne(id string) (nortification *model.Nortification, errors *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("nortification")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&nortification)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	DbClose(c)
	return nortification, nil	
}

func (r *nortificationrepository) GetAll(nortifications []model.Nortification) ([]model.Nortification, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("nortification")
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
	err := cur.Decode(&nortifications)
		if err != nil { 
			return nil,	httperrors.NewNotFoundError("Error while decoding results!")
		}
	// do something with result....
	}
	if err := cur.Err(); err != nil {
		return nil,	httperrors.NewNotFoundError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}	
	DbClose(c)
    return nortifications, nil

}

func (r *nortificationrepository) Update(id string, nortification *model.Nortification) (*httperrors.HttpError) {
	unortification := &model.Nortification{}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("nortification")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&unortification)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	if nortification.Name  == "" {
		nortification.Name = unortification.Name
	}
	if nortification.Title  == "" {
		nortification.Title = unortification.Title
	}

	if nortification.Description  == "" {
		nortification.Description = unortification.Description
	}
	_, err = collection.UpdateOne(ctx, filter, nortification)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Update of nortification Failed, %d", err))
	} 
	DbClose(c)
	return nil
}
func (r nortificationrepository) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("nortification")
	filter := bson.M{"_id": id}
	ok, err := collection.DeleteOne(ctx, filter)
	if ok == nil {
		return nil, httperrors.NewNotFoundError(fmt.Sprintf("deletion of %d failed", err))
	}
	DbClose(c)
		return httperrors.NewSuccessMessage("deleted successfully"), nil
}

