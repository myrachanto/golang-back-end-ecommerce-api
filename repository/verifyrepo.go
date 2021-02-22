package repository

import (
	"fmt"
    "go.mongodb.org/mongo-driver/bson"
		"github.com/myrachanto/ecommerce/httperrors"
		"github.com/myrachanto/ecommerce/model" 
)
var (
	Verifyrepository verifyrepository = verifyrepository{}
)

type verifyrepository struct{}

func (r *verifyrepository) Create(verify *model.Verify) (*httperrors.HttpError) {
	if err1 := verify.Validate(); err1 != nil {
		return err1
	}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("verify")
	_, err := collection.InsertOne(ctx, verify)
		if err != nil {
			return httperrors.NewBadRequestError(fmt.Sprintf("Create verify Failed, %d", err))
	}
	DbClose(c)
	return nil
}

func (r *verifyrepository) GetOne(id string) (verify *model.Verify, errors *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("verify")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&verify)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	DbClose(c)
	return verify, nil	
}

func (r *verifyrepository) GetAll(verifys []model.Verify) ([]model.Verify, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("verify")
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
	err := cur.Decode(&verifys)
		if err != nil { 
			return nil,	httperrors.NewNotFoundError("Error while decoding results!")
		}
	// do something with result....
	}
	if err := cur.Err(); err != nil {
		return nil,	httperrors.NewNotFoundError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}	
	DbClose(c)
    return verifys, nil

}

func (r *verifyrepository) Update(id string, verify *model.Verify) (*httperrors.HttpError) {
	uverify := &model.Verify{}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("verify")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&uverify)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	if verify.Question  == "" {
		verify.Question = uverify.Question
	}
	if verify.Answer  == "" {
		verify.Answer = uverify.Answer
	}

	if verify.Hint  == "" {
		verify.Hint = uverify.Hint
	}
	_, err = collection.UpdateOne(ctx, filter, verify)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Update of verify Failed, %d", err))
	} 
	DbClose(c)
	return nil
}
func (r verifyrepository) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("verify")
	filter := bson.M{"_id": id}
	ok, err := collection.DeleteOne(ctx, filter)
	if ok == nil {
		return nil, httperrors.NewNotFoundError(fmt.Sprintf("deletion of %d failed", err))
	}
	DbClose(c)
		return httperrors.NewSuccessMessage("deleted successfully"), nil
	
}

