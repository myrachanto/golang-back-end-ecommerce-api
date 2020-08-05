package repository

import (
	"fmt"
    "go.mongodb.org/mongo-driver/bson"
	"github.com/myrachanto/asokomonolith/httperrors"
	"github.com/myrachanto/asokomonolith/model" 
)
var (
	Subcategoryrepository subcategoryrepository = subcategoryrepository{}
)

type subcategoryrepository struct{}

func (r *subcategoryrepository) Create(subcategory *model.Subcategory) (*httperrors.HttpError) {
	if err1 := subcategory.Validate(); err1 != nil {
		return err1
	}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("subcategory")
	_, err := collection.InsertOne(ctx, subcategory)
		if err != nil {
			return httperrors.NewBadRequestError(fmt.Sprintf("Create subcategory Failed, %d", err))
	}
	DbClose(c)
	return nil
}

func (r *subcategoryrepository) GetOne(id string) (subcategory *model.Subcategory, errors *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("subcategory")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&subcategory)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	DbClose(c)
	return subcategory, nil	
}

func (r *subcategoryrepository) GetAll(subcategorys []model.Subcategory) ([]model.Subcategory, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("subcategory")
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
	err := cur.Decode(&subcategorys)
		if err != nil { 
			return nil,	httperrors.NewNotFoundError("Error while decoding results!")
		}
	// do something with result....
	}
	if err := cur.Err(); err != nil {
		return nil,	httperrors.NewNotFoundError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}	
	DbClose(c)
    return subcategorys, nil

}

func (r *subcategoryrepository) Update(id string, subcategory *model.Subcategory) (*httperrors.HttpError) {
	usubcategory := &model.Subcategory{}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("subcategory")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&usubcategory)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	if subcategory.Name  == "" {
		subcategory.Name = usubcategory.Name
	}
	if subcategory.Title  == "" {
		subcategory.Title = usubcategory.Title
	}

	if subcategory.Description  == "" {
		subcategory.Description = usubcategory.Description
	}
	_, err = collection.UpdateOne(ctx, filter, subcategory)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Update of subcategory Failed, %d", err))
	} 
	DbClose(c)
	return nil
}
func (r subcategoryrepository) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("subcategory")
	filter := bson.M{"_id": id}
	ok, err := collection.DeleteOne(ctx, filter)
	if ok == nil {
		return nil, httperrors.NewNotFoundError(fmt.Sprintf("deletion of %d failed", err))
	}
	DbClose(c)
		return httperrors.NewSuccessMessage("deleted successfully"), nil
	
}

