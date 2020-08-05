package service

import (
	"github.com/myrachanto/asokomonolith/httperrors"
	"github.com/myrachanto/asokomonolith/model" 
	r "github.com/myrachanto/asokomonolith/repository"
)

var (
	TownService  = townService{}
)

type townService struct {
}

func (service townService) Create(town *model.Town) (*httperrors.HttpError) {
	err1 := r.Townrepository.Create(town)
	 return err1

}

func (service townService) GetOne(id string) (*model.Town, *httperrors.HttpError) {
	town, err1 := r.Townrepository.GetOne(id)
	return town, err1
}

func (service townService) GetAll(towns []model.Town) ([]model.Town, *httperrors.HttpError) {
	towns, err := r.Townrepository.GetAll(towns)
	return towns, err
}

func (service townService) Update(id string, town *model.Town) (*httperrors.HttpError) {
	err1 := r.Townrepository.Update(id, town)
	return err1
}
func (service townService) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
		success, failure := r.Townrepository.Delete(id)
		return success, failure
}
