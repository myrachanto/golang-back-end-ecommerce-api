package model

import "github.com/myrachanto/ecommerce/httperrors"

type Product struct {
	Name        string  `json:"name"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Code       string  `json:"code"`
	Majorcategory string  `json:"majorcat"`
	Category string  `json:"category"`
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
