package model
import "github.com/myrachanto/ecommerce/httperrors"

type Nortification struct {
	Name        string  `bson:"name"`
	Title       string  `bson:"title"`
	Description string  `bson:"description"`
	Read bool
	Base
}

func (nortification Nortification) Validate() *httperrors.HttpError {
	if nortification.Name == "" {
		return httperrors.NewNotFoundError("Invalid Name")
	}
	if nortification.Title == "" {
		return httperrors.NewNotFoundError("Invalid title")
	}
	if nortification.Description == "" {
		return httperrors.NewNotFoundError("Invalid Description")
	}
	return nil
}
