package repository

import (
	"fmt"
    "go.mongodb.org/mongo-driver/bson"
	"github.com/myrachanto/asokomonolith/httperrors"
	"github.com/myrachanto/asokomonolith/model" 
)
var (
	Ratingrepository ratingrepository = ratingrepository{}
)

type ratingrepository struct{}

func (r *ratingrepository) Create(rating *model.Rating) (*httperrors.HttpError) {
	if err1 := rating.Validate(); err1 != nil {
		return err1
	}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("rating")
	_, err := collection.InsertOne(ctx, rating)
		if err != nil {
			return httperrors.NewBadRequestError(fmt.Sprintf("Create rating Failed, %d", err))
	}
	DbClose(c)
	return nil
}

func (r *ratingrepository) GetOne(id string) (rating *model.Rating, errors *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("rating")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&rating)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	DbClose(c)
	return rating, nil	
}

func (r *ratingrepository) GetAll(ratings []model.Rating) ([]model.Rating, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("rating")
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
	err := cur.Decode(&ratings)
		if err != nil { 
			return nil,	httperrors.NewNotFoundError("Error while decoding results!")
		}
	// do something with result....
	}
	if err := cur.Err(); err != nil {
		return nil,	httperrors.NewNotFoundError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}	
	DbClose(c)
    return ratings, nil

}

func (r *ratingrepository) Update(id string, rating *model.Rating) (*httperrors.HttpError) {
	urating := &model.Rating{}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("rating")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&urating)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	if rating.Name  == "" {
		rating.Name = urating.Name
	}
	if rating.Rates  > 0 {
		rating.Rates = urating.Rates
	}

	if rating.Comments  == "" {
		rating.Comments = urating.Comments
	}
	_, err = collection.UpdateOne(ctx, filter, rating)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Update of rating Failed, %d", err))
	} 
	DbClose(c)
	return nil
}
func (r ratingrepository) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("rating")
	filter := bson.M{"_id": id}
	ok, err := collection.DeleteOne(ctx, filter)
	if ok == nil {
		return nil, httperrors.NewNotFoundError(fmt.Sprintf("deletion of %d failed", err))
	}
	DbClose(c)
		return httperrors.NewSuccessMessage("deleted successfully"), nil
}

