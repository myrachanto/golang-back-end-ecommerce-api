package model

import(
	"github.com/myrachanto/asokomonolith/httperrors"
)

type Industry struct {
	Id string `bson:"_id"`
	Name string `bson:"name"`
	Title string `bson:"title"`
	Description string `bson:"description"`
	Picture string `bson:"picture"`
	Base
}
func (industry Industry) Validate() *httperrors.HttpError{
	if industry.Name == "" {
		return httperrors.NewNotFoundError("Invalid Name")
	}
	if industry.Title == "" {
		return httperrors.NewNotFoundError("Invalid title")
	}
	if industry.Description == "" {
		return httperrors.NewNotFoundError("Invalid Description")
	}
	return nil
}