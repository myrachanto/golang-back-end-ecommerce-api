package model

import "github.com/myrachanto/asokomonolith/httperrors"

type Product struct {
	Id          string  `bson:"_id"`
	Name        string  `bson:"name"`
	Title       string  `bson:"title"`
	Description string  `bson:"description"`
	Majorcategory string
	Category string
	Subcategory string
	Price float64
	Tag []Tag
	Rates []Rating
	Featured bool
	Promotion bool
	Hotdeals bool
	Base
}

func (product Product) Validate() *httperrors.HttpError {
	if product.Name == "" {
		return httperrors.NewNotFoundError("Invalid Name")
	}
	if product.Title == "" {
		return httperrors.NewNotFoundError("Invalid title")
	}
	if product.Description == "" {
		return httperrors.NewNotFoundError("Invalid Description")
	}
	return nil
}
