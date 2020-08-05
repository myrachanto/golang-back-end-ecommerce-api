package model

import(
	"github.com/myrachanto/asokomonolith/httperrors"
)

type Street struct {
	Id string `bson:"_id"`
	Name string `bson:"name"`
	Title string `bson:"title"`
	Description string `bson:"description"`
	Population float64 `bson:"population"`
	Base
}
func (street Street) Validate() *httperrors.HttpError{
	if street.Name == "" {
		return httperrors.NewNotFoundError("Invalid Name")
	}
	if street.Title == "" {
		return httperrors.NewNotFoundError("Invalid title")
	}
	if street.Description == "" {
		return httperrors.NewNotFoundError("Invalid Description")
	}
	return nil
}