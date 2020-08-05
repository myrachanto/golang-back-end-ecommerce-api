package model

import(
	"github.com/myrachanto/asokomonolith/httperrors"
)

type Town struct {
	Id string `bson:"_id"`
	Name string `bson:"name"`
	Title string `bson:"title"`
	Description string `bson:"description"`
	Population float64 `bson:"population"`
	Picture string `bson:"picture"`
	Base
}
func (town Town) Validate() *httperrors.HttpError{
	if town.Name == "" {
		return httperrors.NewNotFoundError("Invalid Name")
	}
	if town.Title == "" {
		return httperrors.NewNotFoundError("Invalid title")
	}
	if town.Description == "" {
		return httperrors.NewNotFoundError("Invalid Description")
	}
	return nil
}