package repository

import (
	"fmt"
    "go.mongodb.org/mongo-driver/bson"
		"github.com/myrachanto/ecommerce/httperrors"
		"github.com/myrachanto/ecommerce/model" 
)
var (
	Tagrepository tagrepository = tagrepository{}
)

type tagrepository struct{}

func (r *tagrepository) Create(tag *model.Tag) (*httperrors.HttpError) {
	if err1 := tag.Validate(); err1 != nil {
		return err1
	}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("tag")
	_, err := collection.InsertOne(ctx, tag)
		if err != nil {
			return httperrors.NewBadRequestError(fmt.Sprintf("Create tag Failed, %d", err))
	}
	DbClose(c)
	return nil
}

func (r *tagrepository) GetOne(id string) (tag *model.Tag, errors *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("tag")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&tag)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	DbClose(c)
	return tag, nil	
}

func (r *tagrepository) GetAll(tags []model.Tag) ([]model.Tag, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("tag")
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
	err := cur.Decode(&tags)
		if err != nil { 
			return nil,	httperrors.NewNotFoundError("Error while decoding results!")
		}
	// do something with result....
	}
	if err := cur.Err(); err != nil {
		return nil,	httperrors.NewNotFoundError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}	
	DbClose(c)
    return tags, nil

}

func (r *tagrepository) Update(id string, tag *model.Tag) (*httperrors.HttpError) {
	utag := &model.Tag{}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("tag")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&utag)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	if tag.Name  == "" {
		tag.Name = utag.Name
	}
	if tag.Title  == "" {
		tag.Title = utag.Title
	}

	if tag.Description  == "" {
		tag.Description = utag.Description
	}
	_, err = collection.UpdateOne(ctx, filter, tag)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Update of tag Failed, %d", err))
	} 
	DbClose(c)
	return nil
}
func (r tagrepository) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("tag")
	filter := bson.M{"_id": id}
	ok, err := collection.DeleteOne(ctx, filter)
	if ok == nil {
		return nil, httperrors.NewNotFoundError(fmt.Sprintf("deletion of %d failed", err))
	}
	DbClose(c)
		return httperrors.NewSuccessMessage("deleted successfully"), nil
	
}

