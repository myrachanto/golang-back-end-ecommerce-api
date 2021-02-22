package model
import "github.com/myrachanto/ecommerce/httperrors"

type Rating struct {
	Id          string  `bson:"_id"`
	Name        string  `bson:"name"`
	Rates       float64  `bson:"rates"`
	Comments string  `bson:"comments"`
	Base
}

func (rating Rating) Validate() *httperrors.HttpError {
	if rating.Name == "" {
		return httperrors.NewNotFoundError("Invalid Name")
	}
	if rating.Rates > 0 {
		return httperrors.NewNotFoundError("Invalid rate")
	}
	if rating.Comments == "" {
		return httperrors.NewNotFoundError("Invalid Comments")
	}
	return nil
}
