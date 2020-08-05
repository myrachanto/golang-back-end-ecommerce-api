package repository

import (
	"fmt"
    "go.mongodb.org/mongo-driver/bson"
	"github.com/myrachanto/asokomonolith/httperrors"
	"github.com/myrachanto/asokomonolith/model" 
)
var (
	Categoryrepository categoryrepository = categoryrepository{}
)

type categoryrepository struct{}

func (r *categoryrepository) Create(category *model.Category) (*httperrors.HttpError) {
	if err1 := category.Validate(); err1 != nil {
		return err1
	}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	
	collection := db.Collection("category")
	fmt.Println(collection,c,db,category)
	_, err := collection.InsertOne(ctx, category)
		if err != nil {
			return httperrors.NewBadRequestError(fmt.Sprintf("Create category Failed, %d", err))
	}
	DbClose(c)
	return nil
}

func (r *categoryrepository) GetOne(id string) (category *model.Category, errors *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("category")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&category)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	DbClose(c)
	return category, nil	
}

func (r *categoryrepository) GetAll(categorys []model.Category) ([]model.Category, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("category")
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
	err := cur.Decode(&categorys)
		if err != nil { 
			return nil,	httperrors.NewNotFoundError("Error while decoding results!")
		}
	// do something with result....
	}
	if err := cur.Err(); err != nil {
		return nil,	httperrors.NewNotFoundError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}	
	DbClose(c)
    return categorys, nil

}

func (r *categoryrepository) Update(id string, category *model.Category) (*httperrors.HttpError) {
	ucategory := &model.Category{}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("category")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&ucategory)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	if category.Name  == "" {
		category.Name = ucategory.Name
	}
	if category.Title  == "" {
		category.Title = ucategory.Title
	}

	if category.Description  == "" {
		category.Description = ucategory.Description
	}
	_, err = collection.UpdateOne(ctx, filter, category)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Update of category Failed, %d", err))
	} 
	DbClose(c)
	return nil
}
func (r categoryrepository) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("category")
	filter := bson.M{"_id": id}
	ok, err := collection.DeleteOne(ctx, filter)
	if ok == nil {
		return nil, httperrors.NewNotFoundError(fmt.Sprintf("deletion of %d failed", err))
	}
	DbClose(c)
		return httperrors.NewSuccessMessage("deleted successfully"), nil
}

