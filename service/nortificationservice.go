package service

import (
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model" 
	r "github.com/myrachanto/ecommerce/repository"
)

var (
	NortificationService  = nortificationService{}
)

type nortificationService struct {
}

func (service nortificationService) Create(nortification *model.Nortification) (*httperrors.HttpError) {
	err1 := r.Nortificationrepository.Create(nortification)
	 return err1

}

func (service nortificationService) GetOne(id string) (*model.Nortification, *httperrors.HttpError) {
	nortification, err1 := r.Nortificationrepository.GetOne(id)
	return nortification, err1
}

func (service nortificationService) GetAll(nortifications []model.Nortification) ([]model.Nortification, *httperrors.HttpError) {
	nortifications, err := r.Nortificationrepository.GetAll(nortifications)
	return nortifications, err
}

func (service nortificationService) Update(id string, nortification *model.Nortification) (*httperrors.HttpError) {
	err1 := r.Nortificationrepository.Update(id, nortification)
	return err1
}
func (service nortificationService) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
		success, failure := r.Nortificationrepository.Delete(id)
		return success, failure
}
