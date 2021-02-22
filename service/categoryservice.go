package service

import (
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model" 
	r "github.com/myrachanto/ecommerce/repository"
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

func (service categoryService) GetAll(code string) ([]*model.Category, *httperrors.HttpError) {
	categorys, err := r.Categoryrepository.GetAll(code)
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
