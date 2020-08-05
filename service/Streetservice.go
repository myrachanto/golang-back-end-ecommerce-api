package service

import (
	"github.com/myrachanto/asokomonolith/httperrors"
	"github.com/myrachanto/asokomonolith/model" 
	r "github.com/myrachanto/asokomonolith/repository"
)

var (
	StreetService  = streetService{}
)


type streetService struct {
}

func (service streetService) Create(street *model.Street) (*httperrors.HttpError) {
	err1 := r.Streetrepository.Create(street)
	 return err1

}

func (service streetService) GetOne(id string) (*model.Street, *httperrors.HttpError) {
	street, err1 := r.Streetrepository.GetOne(id)
	return street, err1
}

func (service streetService) GetAll(streets []model.Street) ([]model.Street, *httperrors.HttpError) {
	streets, err := r.Streetrepository.GetAll(streets)
	return streets, err
}

func (service streetService) Update(id string, street *model.Street) (*httperrors.HttpError) {
	err1 := r.Streetrepository.Update(id, street)
	return err1
}
func (service streetService) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
		success, failure := r.Streetrepository.Delete(id)
		return success, failure
}
