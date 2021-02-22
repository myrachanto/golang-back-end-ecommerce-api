package model


import "github.com/myrachanto/ecommerce/httperrors"

type Subcategory struct {
	Id string `bson:"_id"`
	Name string `bson:"name"`
	Title string `bson:"title"`
	Description string `bson:"description"`
	Base
}
func (subcategory Subcategory) Validate() *httperrors.HttpError{
	if subcategory.Name == "" {
		return httperrors.NewNotFoundError("Invalid Name")
	}
	if subcategory.Title == "" {
		return httperrors.NewNotFoundError("Invalid title")
	}
	if subcategory.Description == "" {
		return httperrors.NewNotFoundError("Invalid Description")
	}
	return nil
}