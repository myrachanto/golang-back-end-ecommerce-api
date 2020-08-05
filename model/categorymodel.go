package model

import(
	"github.com/myrachanto/asokomonolith/httperrors"
)

type Category struct {
	Id string `bson:"_id"`
	Name string `bson:"name"`
	Title string `bson:"title"`
	Description string `bson:"description"`
	Base
}
func (category Category) Validate() *httperrors.HttpError{
	if category.Name == "" {
		return httperrors.NewNotFoundError("Invalid Name")
	}
	if category.Title == "" {
		return httperrors.NewNotFoundError("Invalid title")
	}
	if category.Description == "" {
		return httperrors.NewNotFoundError("Invalid Description")
	}
	return nil
}