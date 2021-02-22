package repository

import (
	"fmt"
    "go.mongodb.org/mongo-driver/bson"
		"github.com/myrachanto/ecommerce/httperrors"
		"github.com/myrachanto/ecommerce/model" 
)
var (
	Townrepository townrepository = townrepository{}
)

type townrepository struct{}

func (r *townrepository) Create(town *model.Town) (*httperrors.HttpError) {
	if err1 := town.Validate(); err1 != nil {
		return err1
	}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("town")
	_, err := collection.InsertOne(ctx, town)
		if err != nil {
			return httperrors.NewBadRequestError(fmt.Sprintf("Create town Failed, %d", err))
	}
	DbClose(c)
	return nil
}

func (r *townrepository) GetOne(id string) (town *model.Town, errors *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("town")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&town)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	DbClose(c)
	return town, nil	
}

func (r *townrepository) GetAll(towns []model.Town) ([]model.Town, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("town")
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
	err := cur.Decode(&towns)
		if err != nil { 
			return nil,	httperrors.NewNotFoundError("Error while decoding results!")
		}
	// do something with result....
	}
	if err := cur.Err(); err != nil {
		return nil,	httperrors.NewNotFoundError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}	
	DbClose(c)
    return towns, nil

}

func (r *townrepository) Update(id string, town *model.Town) (*httperrors.HttpError) {
	utown := &model.Town{}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("town")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&utown)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	if town.Name  == "" {
		town.Name = utown.Name
	}
	if town.Title  == "" {
		town.Title = utown.Title
	}

	if town.Description  == "" {
		town.Description = utown.Description
	}
	_, err = collection.UpdateOne(ctx, filter, town)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Update of town Failed, %d", err))
	} 
	DbClose(c)
	return nil
}
func (r townrepository) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("town")
	filter := bson.M{"_id": id}
	ok, err := collection.DeleteOne(ctx, filter)
	if ok == nil {
		return nil, httperrors.NewNotFoundError(fmt.Sprintf("deletion of %d failed", err))
	}
	DbClose(c)
		return httperrors.NewSuccessMessage("deleted successfully"), nil
	
}

