package repository

import (
	"fmt"
    "go.mongodb.org/mongo-driver/bson"
	"github.com/myrachanto/asokomonolith/httperrors"
	"github.com/myrachanto/asokomonolith/model" 
)
var (
	Shoprepository shoprepository = shoprepository{}
)

type shoprepository struct{}

func (r *shoprepository) Create(shop *model.Shop) (*httperrors.HttpError) {
	if err1 := shop.Validate(); err1 != nil {
		return err1
	}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("shop")
	_, err := collection.InsertOne(ctx, shop)
		if err != nil {
			return httperrors.NewBadRequestError(fmt.Sprintf("Create shop Failed, %d", err))
	}
	DbClose(c)
	return nil
}

func (r *shoprepository) GetOne(id string) (shop *model.Shop, errors *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("shop")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&shop)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	DbClose(c)
	return shop, nil	
}

func (r *shoprepository) GetAll(shops []model.Shop) ([]model.Shop, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("shop")
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
	err := cur.Decode(&shops)
		if err != nil { 
			return nil,	httperrors.NewNotFoundError("Error while decoding results!")
		}
	// do something with result....
	}
	if err := cur.Err(); err != nil {
		return nil,	httperrors.NewNotFoundError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}	
	DbClose(c)
    return shops, nil

}

func (r *shoprepository) Update(id string, shop *model.Shop) (*httperrors.HttpError) {
	ushop := &model.Shop{}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("shop")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&ushop)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	
	if shop.Name  == "" {
		shop.Name = ushop.Name
	}
	if shop.Title  == "" {
		shop.Title = ushop.Title
	}
	if shop.Description  == "" {
		shop.Description = ushop.Description
	}
	if shop.Industry  == "" {
		shop.Industry = ushop.Industry
	}
	if shop.County  == "" {
		shop.County = ushop.County
	}
	if shop.Town  == "" {
		shop.Town = ushop.Town
	}
	if shop.Street  == "" {
		shop.Street = ushop.Street
	}
	if len(shop.Products)  > 0  {
		shop.Products = ushop.Products
	}
	if len(shop.Tag)  > 0  {
		shop.Tag = ushop.Tag
	}
	if len(shop.Rates)  > 0  {
		shop.Rates = ushop.Rates
	}
	if len(shop.Verified)  > 0  {
		shop.Verified = ushop.Verified
	}
	if len(shop.Reliable)  > 0  {
		shop.Reliable = ushop.Reliable
	}
	_, err = collection.UpdateOne(ctx, filter, shop)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Update of shop Failed, %d", err))
	} 
	DbClose(c)
	return nil
}
func (r shoprepository) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("shop")
	filter := bson.M{"_id": id}
	ok, err := collection.DeleteOne(ctx, filter)
	if ok == nil {
		return nil, httperrors.NewNotFoundError(fmt.Sprintf("deletion of %d failed", err))
	}
	DbClose(c)
		return httperrors.NewSuccessMessage("deleted successfully"), nil
}

