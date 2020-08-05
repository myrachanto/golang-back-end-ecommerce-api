package model

import(
	"github.com/myrachanto/asokomonolith/httperrors"
)

type Majorcategory struct {
	Id string`bson:"_id"`
	Name string `bson:"name"`
	Title string `bson:"title"`
	Description string `bson:"description"`
	Base
}
func (majorcategory Majorcategory) Validate() *httperrors.HttpError{
	if majorcategory.Name == "" {
		return httperrors.NewNotFoundError("Invalid Name")
	}
	if majorcategory.Title == "" {
		return httperrors.NewNotFoundError("Invalid title")
	}
	if majorcategory.Description == "" {
		return httperrors.NewNotFoundError("Invalid Description")
	}
	return nil
}