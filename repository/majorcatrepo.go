package repository

import (
	"fmt"
    "go.mongodb.org/mongo-driver/bson"
	"github.com/myrachanto/asokomonolith/httperrors"
	"github.com/myrachanto/asokomonolith/model" 
)
var (
	Majorcategoryrepository majorcategoryrepository = majorcategoryrepository{}
)

type majorcategoryrepository struct{}

func (r *majorcategoryrepository) Create(majorcategory *model.Majorcategory) (*httperrors.HttpError) {
	if err1 := majorcategory.Validate(); err1 != nil {
		return err1
	}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("majorcategory")
	_, err := collection.InsertOne(ctx, majorcategory)
		if err != nil {
			return httperrors.NewBadRequestError(fmt.Sprintf("Create majorcategory Failed, %d", err))
	}
	DbClose(c)
	return nil
}

func (r *majorcategoryrepository) GetOne(id string) (majorcategory *model.Majorcategory, errors *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("majorcategory")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&majorcategory)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	DbClose(c)
	return majorcategory, nil	
}

func (r *majorcategoryrepository) GetAll(majorcategorys []model.Majorcategory) ([]model.Majorcategory, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("majorcategory")
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
	err := cur.Decode(&majorcategorys)
		if err != nil { 
			return nil,	httperrors.NewNotFoundError("Error while decoding results!")
		}
	// do something with result....
	}
	if err := cur.Err(); err != nil {
		return nil,	httperrors.NewNotFoundError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}	
	DbClose(c)
    return majorcategorys, nil

}

func (r *majorcategoryrepository) Update(id string, majorcategory *model.Majorcategory) (*httperrors.HttpError) {
	umajorcategory := &model.Majorcategory{}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("majorcategory")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&umajorcategory)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	if majorcategory.Name  == "" {
		majorcategory.Name = umajorcategory.Name
	}
	if majorcategory.Title  == "" {
		majorcategory.Title = umajorcategory.Title
	}

	if majorcategory.Description  == "" {
		majorcategory.Description = umajorcategory.Description
	}
	_, err = collection.UpdateOne(ctx, filter, majorcategory)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Update of majorcategory Failed, %d", err))
	} 
	DbClose(c)
	return nil
}
func (r majorcategoryrepository) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("majorcategory")
	filter := bson.M{"_id": id}
	ok, err := collection.DeleteOne(ctx, filter)
	if ok == nil {
		return nil, httperrors.NewNotFoundError(fmt.Sprintf("deletion of %d failed", err))
	}
	DbClose(c)
		return httperrors.NewSuccessMessage("deleted successfully"), nil
}

