package service

import (
	"github.com/myrachanto/asokomonolith/httperrors"
	"github.com/myrachanto/asokomonolith/model" 
	r "github.com/myrachanto/asokomonolith/repository"
)

var (
	RatingService  = ratingService{}
)
type ratingService struct {
}

func (service ratingService) Create(rating *model.Rating) (*httperrors.HttpError) {
	err1 := r.Ratingrepository.Create(rating)
	 return err1

}

func (service ratingService) GetOne(id string) (*model.Rating, *httperrors.HttpError) {
	rating, err1 := r.Ratingrepository.GetOne(id)
	return rating, err1
}

func (service ratingService) GetAll(ratings []model.Rating) ([]model.Rating, *httperrors.HttpError) {
	ratings, err := r.Ratingrepository.GetAll(ratings)
	return ratings, err
}

func (service ratingService) Update(id string, rating *model.Rating) (*httperrors.HttpError) {
	err1 := r.Ratingrepository.Update(id, rating)
	return err1
}
func (service ratingService) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
		success, failure := r.Ratingrepository.Delete(id)
		return success, failure
}
