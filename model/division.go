package model


import "github.com/myrachanto/ecommerce/httperrors"

type Division struct {
	Id string `bson:"_id"`
	Name string `bson:"name"`
	Title string `bson:"title"`
	Description string `bson:"description"`
	Base
}
func (division Division) Validate() *httperrors.HttpError{
	if division.Name == "" {
		return httperrors.NewNotFoundError("Invalid Name")
	}
	if division.Title == "" {
		return httperrors.NewNotFoundError("Invalid title")
	}
	if division.Description == "" {
		return httperrors.NewNotFoundError("Invalid Description")
	}
	return nil
}