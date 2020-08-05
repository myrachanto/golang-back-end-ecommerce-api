package service

import (
	"github.com/myrachanto/asokomonolith/httperrors"
	"github.com/myrachanto/asokomonolith/model" 
	r "github.com/myrachanto/asokomonolith/repository"
)

var (
	ProductService  = productService{}
)
type productService struct {
}

func (service productService) Create(product *model.Product) (*httperrors.HttpError) {
	err1 := r.Productrepository.Create(product)
	 return err1

}

func (service productService) GetOne(id string) (*model.Product, *httperrors.HttpError) {
	product, err1 := r.Productrepository.GetOne(id)
	return product, err1
}

func (service productService) GetAll(products []model.Product) ([]model.Product, *httperrors.HttpError) {
	products, err := r.Productrepository.GetAll(products)
	return products, err
}

func (service productService) Update(id string, product *model.Product) (*httperrors.HttpError) {
	err1 := r.Productrepository.Update(id, product)
	return err1
}
func (service productService) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
		success, failure := r.Productrepository.Delete(id)
		return success, failure
}
