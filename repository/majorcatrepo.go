package repository

import (
	"fmt"
	"strconv"
    "go.mongodb.org/mongo-driver/bson"
		"github.com/myrachanto/ecommerce/httperrors"
		"github.com/myrachanto/ecommerce/model" 
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
	code, err1 := Majorcategoryrepository.genecode()
	if err1 != nil {
		return err1
	}
	majorcategory.Code = code
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

func (r *majorcategoryrepository) GetAll() ([]*model.Majorcategory, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	majorcategorys := []*model.Majorcategory{}
	collection := db.Collection("majorcategory")
	filter := bson.M{}
	cur, err := collection.Find(ctx, filter)
	if err != nil { 
		return nil,	httperrors.NewNotFoundError("no results found")
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var majorcategory model.Majorcategory
		err := cur.Decode(&majorcategory)
			if err != nil { 
				return nil,	httperrors.NewNotFoundError("Error while decoding results!")
			}
	 majorcategorys = append(majorcategorys, &majorcategory)
	}
	if err := cur.Err(); err != nil {
		return nil,	httperrors.NewNotFoundError("Error with cursor!")
	}	
	DbClose(c)
    return majorcategorys, nil

}

func (r *majorcategoryrepository) Update(code string, majorcategory *model.Majorcategory) (*httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	result, err3 := Majorcategoryrepository.getuno(code)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(result)
	if majorcategory.Name == ""{
		majorcategory.Name = result.Name
	}
	if majorcategory.Title == ""{
		majorcategory.Title = result.Title
	}
	if majorcategory.Description == ""{
		majorcategory.Description = result.Description
	}
	if majorcategory.Code == ""{
		majorcategory.Code = result.Code
	}
	collection := db.Collection("majorcategory")
	filter := bson.M{"code": code}
	fmt.Println(filter)
	fmt.Println(majorcategory)
	update := bson.M{"$set": majorcategory}
	_, err := collection.UpdateOne(ctx, filter, update)
		if err != nil {
		return	httperrors.NewNotFoundError("Error updating!")
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
func (r majorcategoryrepository)genecode()(string, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return "", t
	}
	db, e := Mongodb();if e != nil {
		return "", e
	}
	collection := db.Collection("majorcategory")
	filter := bson.M{}
	count, err := collection.CountDocuments(ctx, filter)
	co := count + 1
	if err != nil { 
		return "",	httperrors.NewNotFoundError("no results found")
	}
	code := "MajorcatCode"+strconv.FormatUint(uint64(co), 10)

	DbClose(c)
	return code, nil
}
func (r majorcategoryrepository)getuno(code string)(result *model.Majorcategory, err *httperrors.HttpError){
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("majorcategory")
	filter := bson.M{"code": code}
	err1 := collection.FindOne(ctx, filter).Decode(&result)
	if err1 != nil {
		return nil, httperrors.NewNotFoundError("no results found")
	}
	DbClose(c)
	return result, nil	
}
