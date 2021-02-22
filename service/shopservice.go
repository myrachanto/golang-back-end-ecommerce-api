package service

import (
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model" 
	r "github.com/myrachanto/ecommerce/repository"
)

var (
	ShopService  = shopService{}
)

type shopService struct {
}

func (service shopService) Create(shop *model.Shop) (*httperrors.HttpError) {
	err1 := r.Shoprepository.Create(shop)
	 return err1

}

func (service shopService) GetOne(id string) (*model.Shop, *httperrors.HttpError) {
	shop, err1 := r.Shoprepository.GetOne(id)
	return shop, err1
}

func (service shopService) GetAll(shops []model.Shop) ([]model.Shop, *httperrors.HttpError) {
	shops, err := r.Shoprepository.GetAll(shops)
	return shops, err
}

func (service shopService) Update(id string, shop *model.Shop) (*httperrors.HttpError) {
	err1 := r.Shoprepository.Update(id, shop)
	return err1
}
func (service shopService) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
		success, failure := r.Shoprepository.Delete(id)
		return success, failure
}
