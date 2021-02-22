package service

import (
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model" 
	r "github.com/myrachanto/ecommerce/repository"
)

var (
	DivisionService  = divisionService{}
)

type divisionService struct {
}

func (service divisionService) Create(division *model.Division) (*httperrors.HttpError) {
	err1 := r.Divisionrepository.Create(division)
	 return err1

}

func (service divisionService) GetOne(id string) (*model.Division, *httperrors.HttpError) {
	division, err1 := r.Divisionrepository.GetOne(id)
	return division, err1
}

func (service divisionService) GetAll(divisions []model.Division) ([]model.Division, *httperrors.HttpError) {
	divisions, err := r.Divisionrepository.GetAll(divisions)
	return divisions, err
}

func (service divisionService) Update(id string, division *model.Division) (*httperrors.HttpError) {
	err1 := r.Divisionrepository.Update(id, division)
	return err1
}
func (service divisionService) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
		success, failure := r.Divisionrepository.Delete(id)
		return success, failure
}
