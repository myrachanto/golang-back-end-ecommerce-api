package service

import (
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model" 
	r "github.com/myrachanto/ecommerce/repository"
)

var (
	IndustryService  = industryService{}
)

type industryService struct {
}

func (service industryService) Create(industry *model.Industry) (*httperrors.HttpError) {
	err1 := r.Industryrepository.Create(industry)
	 return err1

}

func (service industryService) GetOne(id string) (*model.Industry, *httperrors.HttpError) {
	industry, err1 := r.Industryrepository.GetOne(id)
	return industry, err1
}

func (service industryService) GetAll(industrys []model.Industry) ([]model.Industry, *httperrors.HttpError) {
	industrys, err := r.Industryrepository.GetAll(industrys)
	return industrys, err
}

func (service industryService) Update(id string, industry *model.Industry) (*httperrors.HttpError) {
	err1 := r.Industryrepository.Update(id, industry)
	return err1
}
func (service industryService) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
		success, failure := r.Industryrepository.Delete(id)
		return success, failure
}
