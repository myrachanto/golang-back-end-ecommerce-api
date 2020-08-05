package repository

import (
	"fmt"
    "go.mongodb.org/mongo-driver/bson"
	"github.com/myrachanto/asokomonolith/httperrors"
	"github.com/myrachanto/asokomonolith/model" 
)
var (
	Divisionrepository divisionrepository = divisionrepository{}
)

type divisionrepository struct{}

func (r *divisionrepository) Create(division *model.Division) (*httperrors.HttpError) {
	if err1 := division.Validate(); err1 != nil {
		return err1
	}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("division")
	_, err := collection.InsertOne(ctx, division)
		if err != nil {
			return httperrors.NewBadRequestError(fmt.Sprintf("Create division Failed, %d", err))
	}
	DbClose(c)
	return nil
}

func (r *divisionrepository) GetOne(id string) (division *model.Division, errors *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("division")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&division)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	DbClose(c)
	return division, nil	
}

func (r *divisionrepository) GetAll(divisions []model.Division) ([]model.Division, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("division")
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
	err := cur.Decode(&divisions)
		if err != nil { 
			return nil,	httperrors.NewNotFoundError("Error while decoding results!")
		}
	// do something with result....
	}
	if err := cur.Err(); err != nil {
		return nil,	httperrors.NewNotFoundError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}	
	DbClose(c)
    return divisions, nil

}

func (r *divisionrepository) Update(id string, division *model.Division) (*httperrors.HttpError) {
	udivision := &model.Division{}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("division")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&udivision)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	if division.Name  == "" {
		division.Name = udivision.Name
	}
	if division.Title  == "" {
		division.Title = udivision.Title
	}

	if division.Description  == "" {
		division.Description = udivision.Description
	}
	_, err = collection.UpdateOne(ctx, filter, division)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Update of division Failed, %d", err))
	} 
	DbClose(c)
	return nil
}
func (r divisionrepository) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("division")
	filter := bson.M{"_id": id}
	ok, err := collection.DeleteOne(ctx, filter)
	if ok == nil {
		return nil, httperrors.NewNotFoundError(fmt.Sprintf("deletion of %d failed", err))
	}
	
	DbClose(c)
	return httperrors.NewSuccessMessage("deleted successfully"), nil
}

