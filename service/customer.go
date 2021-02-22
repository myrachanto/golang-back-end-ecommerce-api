package service

import (
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model" 
	r "github.com/myrachanto/ecommerce/repository"
)

var (
	CustomerService  = customerService{}
)
type customerService struct {
}

func (service customerService) Create(customer *model.Customer) (*httperrors.HttpError) {
	err1 := r.Customerrepository.Create(customer)
	 return err1

}

func (service customerService) GetOne(id string) (*model.Customer, *httperrors.HttpError) {
	customer, err1 := r.Customerrepository.GetOne(id)
	return customer, err1
}

func (service customerService) GetAll(customers []model.Customer) ([]model.Customer, *httperrors.HttpError) {
	customers, err := r.Customerrepository.GetAll(customers)
	return customers, err
}

func (service customerService) Update(id string, customer *model.Customer) (*httperrors.HttpError) {
	err1 := r.Customerrepository.Update(id, customer)
	return err1
}
func (service customerService) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
		success, failure := r.Customerrepository.Delete(id)
		return success, failure
}
