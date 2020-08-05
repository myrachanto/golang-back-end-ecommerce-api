package model

import "github.com/myrachanto/asokomonolith/httperrors"

type Tag struct {
	Id          string  `bson:"_id"`
	Name        string  `bson:"name"`
	Title       string  `bson:"title"`
	Description string  `bson:"description"`
	Base
}

func (tag Tag) Validate() *httperrors.HttpError {
	if tag.Name == "" {
		return httperrors.NewNotFoundError("Invalid Name")
	}
	if tag.Title == "" {
		return httperrors.NewNotFoundError("Invalid title")
	}
	if tag.Description == "" {
		return httperrors.NewNotFoundError("Invalid Description")
	}
	return nil
}
