package model


import "github.com/myrachanto/ecommerce/httperrors"

type Majorcategory struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Description string `json:"description"`
	Code string `json:"code"`
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