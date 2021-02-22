package repository

import (
	"fmt"
	"strconv"
    "go.mongodb.org/mongo-driver/bson"
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model" 
)
//Categoryrepository ...
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
	code, err1 := Categoryrepository.genecode()
	if err1 != nil {
		return err1
	}
	category.Code = code
	collection := db.Collection("category")
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

func (r *categoryrepository) GetAll(code string) ([]*model.Category, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	categorys := []*model.Category{}
	collection := db.Collection("category")
	filter := bson.M{"majorcat": code}
	fmt.Println(filter)
	cur, err := collection.Find(ctx, filter)
	if err != nil { 
		return nil,	httperrors.NewNotFoundError("no results found")
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var category model.Category
		err := cur.Decode(&category)
			if err != nil { 
				return nil,	httperrors.NewNotFoundError("Error while decoding results!")
			}
	 categorys = append(categorys, &category)
	}
	if err := cur.Err(); err != nil {
		return nil,	httperrors.NewNotFoundError("Error with cursor!")
	}	
	DbClose(c)
    return categorys, nil

}

func (r *categoryrepository) Update(code string, category *model.Category) (*httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	result, err3 := Categoryrepository.getuno(code)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(result)
	if category.Name == ""{
		category.Name = result.Name
	}
	if category.Title == ""{
		category.Title = result.Title
	}
	if category.Description == ""{
		category.Description = result.Description
	}
	if category.Code == ""{
		category.Code = result.Code
	}
	collection := db.Collection("category")
	filter := bson.M{"code": code}
	fmt.Println(filter)
	fmt.Println(category)
	update := bson.M{"$set": category}
	_, err := collection.UpdateOne(ctx, filter, update)
		if err != nil {
		return	httperrors.NewNotFoundError("Error updating!")
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
func (r categoryrepository)genecode()(string, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return "", t
	}
	db, e := Mongodb();if e != nil {
		return "", e
	}
	collection := db.Collection("category")
	filter := bson.M{}
	count, err := collection.CountDocuments(ctx, filter)
	co := count + 1
	if err != nil { 
		return "",	httperrors.NewNotFoundError("no results found")
	}
	code := "CategoryCode"+strconv.FormatUint(uint64(co), 10)

	DbClose(c)
	return code, nil
}
func (r categoryrepository)getuno(code string)(result *model.Category, err *httperrors.HttpError){
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("category")
	filter := bson.M{"code": code}
	err1 := collection.FindOne(ctx, filter).Decode(&result)
	if err1 != nil {
		return nil, httperrors.NewNotFoundError("no results found")
	}
	DbClose(c)
	return result, nil	
}
