package service

import (
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model" 
	r "github.com/myrachanto/ecommerce/repository"
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

func (service majorcategoryService) GetAll() ([]*model.Majorcategory, *httperrors.HttpError) {
	majorcategorys, err := r.Majorcategoryrepository.GetAll()
	return majorcategorys, err
}

func (service majorcategoryService) Update(code string, majorcategory *model.Majorcategory) (*httperrors.HttpError) {
	err1 := r.Majorcategoryrepository.Update(code, majorcategory)
	return err1
}
func (service majorcategoryService) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
		success, failure := r.Majorcategoryrepository.Delete(id)
		return success, failure
}
