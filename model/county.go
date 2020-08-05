package model

import "github.com/myrachanto/asokomonolith/httperrors"

type County struct {
	Id          string  `bson:"_id"`
	Name        string  `bson:"name"`
	Title       string  `bson:"title"`
	Description string  `bson:"description"`
	Population  float64 `bson:"float"`
	Base
}

func (county County) Validate() *httperrors.HttpError {
	if county.Name == "" {
		return httperrors.NewNotFoundError("Invalid Name")
	}
	if county.Title == "" {
		return httperrors.NewNotFoundError("Invalid title")
	}
	if county.Description == "" {
		return httperrors.NewNotFoundError("Invalid Description")
	}
	return nil
}
