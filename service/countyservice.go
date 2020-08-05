package service

import (
	"github.com/myrachanto/asokomonolith/httperrors"
	"github.com/myrachanto/asokomonolith/model" 
	r "github.com/myrachanto/asokomonolith/repository"
)

var (
	CountyService  = countyService{}
)

type countyService struct {
}

func (service countyService) Create(county *model.County) (*httperrors.HttpError) {
	err1 := r.Countyrepository.Create(county)
	 return err1

}

func (service countyService) GetOne(id string) (*model.County, *httperrors.HttpError) {
	county, err1 := r.Countyrepository.GetOne(id)
	return county, err1
}

func (service countyService) GetAll(countys []model.County) ([]model.County, *httperrors.HttpError) {
	countys, err := r.Countyrepository.GetAll(countys)
	return countys, err
}

func (service countyService) Update(id string, county *model.County) (*httperrors.HttpError) {
	err1 := r.Countyrepository.Update(id, county)
	return err1
}
func (service countyService) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
		success, failure := r.Countyrepository.Delete(id)
		return success, failure
}
