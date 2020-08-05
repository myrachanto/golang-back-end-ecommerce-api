package service

import (
	"github.com/myrachanto/asokomonolith/httperrors"
	"github.com/myrachanto/asokomonolith/model" 
	r "github.com/myrachanto/asokomonolith/repository"
)

var (
	SubcategoryService  = subcategoryService{}
)

type subcategoryService struct {
}

func (service subcategoryService) Create(subcategory *model.Subcategory) (*httperrors.HttpError) {
	err1 := r.Subcategoryrepository.Create(subcategory)
	 return err1

}

func (service subcategoryService) GetOne(id string) (*model.Subcategory, *httperrors.HttpError) {
	subcategory, err1 := r.Subcategoryrepository.GetOne(id)
	return subcategory, err1
}

func (service subcategoryService) GetAll(subcategorys []model.Subcategory) ([]model.Subcategory, *httperrors.HttpError) {
	subcategorys, err := r.Subcategoryrepository.GetAll(subcategorys)
	return subcategorys, err
}

func (service subcategoryService) Update(id string, subcategory *model.Subcategory) (*httperrors.HttpError) {
	err1 := r.Subcategoryrepository.Update(id, subcategory)
	return err1
}
func (service subcategoryService) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
		success, failure := r.Subcategoryrepository.Delete(id)
		return success, failure
}
