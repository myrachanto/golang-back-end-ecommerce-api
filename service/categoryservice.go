package service

import (
	"github.com/myrachanto/asokomonolith/httperrors"
	"github.com/myrachanto/asokomonolith/model" 
	r "github.com/myrachanto/asokomonolith/repository"
)

var (
	CategoryService  = categoryService{}
)

type categoryService struct {
}

func (service categoryService) Create(category *model.Category) (*httperrors.HttpError) {
	err1 := r.Categoryrepository.Create(category)
	 return err1

}

func (service categoryService) GetOne(id string) (*model.Category, *httperrors.HttpError) {
	category, err1 := r.Categoryrepository.GetOne(id)
	return category, err1
}

func (service categoryService) GetAll(categorys []model.Category) ([]model.Category, *httperrors.HttpError) {
	categorys, err := r.Categoryrepository.GetAll(categorys)
	return categorys, err
}

func (service categoryService) Update(id string, category *model.Category) (*httperrors.HttpError) {
	err1 := r.Categoryrepository.Update(id, category)
	return err1
}
func (service categoryService) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
		success, failure := r.Categoryrepository.Delete(id)
		return success, failure
}
