package service

import (
	"github.com/myrachanto/asokomonolith/httperrors"
	"github.com/myrachanto/asokomonolith/model" 
	r "github.com/myrachanto/asokomonolith/repository"
)

var (
	MajorcategoryService  = majorcategoryService{}
)

type majorcategoryService struct {
}

func (service majorcategoryService) Create(majorcategory *model.Majorcategory) (*httperrors.HttpError) {
	err1 := r.Majorcategoryrepository.Create(majorcategory)
	 return err1

}

func (service majorcategoryService) GetOne(id string) (*model.Majorcategory, *httperrors.HttpError) {
	majorcategory, err1 := r.Majorcategoryrepository.GetOne(id)
	return majorcategory, err1
}

func (service majorcategoryService) GetAll(majorcategorys []model.Majorcategory) ([]model.Majorcategory, *httperrors.HttpError) {
	majorcategorys, err := r.Majorcategoryrepository.GetAll(majorcategorys)
	return majorcategorys, err
}

func (service majorcategoryService) Update(id string, majorcategory *model.Majorcategory) (*httperrors.HttpError) {
	err1 := r.Majorcategoryrepository.Update(id, majorcategory)
	return err1
}
func (service majorcategoryService) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
		success, failure := r.Majorcategoryrepository.Delete(id)
		return success, failure
}
