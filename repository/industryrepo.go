package repository

import (
	"fmt"
    "go.mongodb.org/mongo-driver/bson"
		"github.com/myrachanto/ecommerce/httperrors"
		"github.com/myrachanto/ecommerce/model"  
)
var (
	Industryrepository industryrepository = industryrepository{}
)

type industryrepository struct{}

func (r *industryrepository) Create(industry *model.Industry) (*httperrors.HttpError) {
	if err1 := industry.Validate(); err1 != nil {
		return err1
	}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("industry")
	_, err := collection.InsertOne(ctx, industry)
		if err != nil {
			return httperrors.NewBadRequestError(fmt.Sprintf("Create industry Failed, %d", err))
	}
	DbClose(c)
	return nil
}

func (r *industryrepository) GetOne(id string) (industry *model.Industry, errors *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("industry")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&industry)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	DbClose(c)
	return industry, nil	
}

func (r *industryrepository) GetAll(industrys []model.Industry) ([]model.Industry, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("industry")
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
	err := cur.Decode(&industrys)
		if err != nil { 
			return nil,	httperrors.NewNotFoundError("Error while decoding results!")
		}
	// do something with result....
	}
	if err := cur.Err(); err != nil {
		return nil,	httperrors.NewNotFoundError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}	
	DbClose(c)
    return industrys, nil

}

func (r *industryrepository) Update(id string, industry *model.Industry) (*httperrors.HttpError) {
	uindustry := &model.Industry{}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("industry")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&uindustry)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	if industry.Name  == "" {
		industry.Name = uindustry.Name
	}
	if industry.Title  == "" {
		industry.Title = uindustry.Title
	}

	if industry.Description  == "" {
		industry.Description = uindustry.Description
	}
	_, err = collection.UpdateOne(ctx, filter, industry)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Update of industry Failed, %d", err))
	} 
	DbClose(c)
	return nil
}
func (r industryrepository) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("industry")
	filter := bson.M{"_id": id}
	ok, err := collection.DeleteOne(ctx, filter)
	if ok == nil {
		return nil, httperrors.NewNotFoundError(fmt.Sprintf("deletion of %d failed", err))
	}
	DbClose(c)
		return httperrors.NewSuccessMessage("deleted successfully"), nil
}

