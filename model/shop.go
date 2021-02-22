package model

import "github.com/myrachanto/ecommerce/httperrors"

type Shop struct {
	Id          string  `bson:"_id"`
	Name        string  `bson:"name"`
	Title       string  `bson:"title"`
	Description string  `bson:"description"`
	Industry string
	County string
	Town string
	Street string
	Location Location
	Tag []Tag
	Products []Product
	Rates []Rating
	Verified []Verified
	Reliable []Veted
	Base
}
type Verified struct{
	Verified bool
	KRApin string
	NationId string
	Passport string
	Picture string
}
type Veted struct{
	Verified bool
	KRApin string
	NationId string
	Passport string
	Picture string
}
type Location struct{
	Longtitude float64
	Latitude float64
}

func (shop Shop) Validate() *httperrors.HttpError {
	if shop.Name == "" {
		return httperrors.NewNotFoundError("Invalid Name")
	}
	if shop.Title == "" {
		return httperrors.NewNotFoundError("Invalid title")
	}
	if shop.Description == "" {
		return httperrors.NewNotFoundError("Invalid Description")
	}
	// if shop.Verified.KRApin == shop.Reliable.KRApin {
	// 	return httperrors.NewNotFoundError("cannot use the same person to enable reliability")
	// }
	// if shop.Verified.NationId == shop.Reliable.NationId {
	// 	return httperrors.NewNotFoundError("cannot use the same person to enable reliability")
	// }
	// if shop.Verified.Passport == shop.Reliable.Passport {
	// 	return httperrors.NewNotFoundError("cannot use the same person to enable reliability")
	// }
	return nil
}
