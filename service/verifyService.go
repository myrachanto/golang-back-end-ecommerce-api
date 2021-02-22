package service

import (
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model" 
	r "github.com/myrachanto/ecommerce/repository"
)

var (
	VerifyService  = verifyService{}
)
type verifyService struct {
}

func (service verifyService) Create(verify *model.Verify) (*httperrors.HttpError) {
	err1 := r.Verifyrepository.Create(verify)
	 return err1

}

func (service verifyService) GetOne(id string) (*model.Verify, *httperrors.HttpError) {
	verify, err1 := r.Verifyrepository.GetOne(id)
	return verify, err1
}

func (service verifyService) GetAll(verifys []model.Verify) ([]model.Verify, *httperrors.HttpError) {
	verifys, err := r.Verifyrepository.GetAll(verifys)
	return verifys, err
}

func (service verifyService) Update(id string, verify *model.Verify) (*httperrors.HttpError) {
	err1 := r.Verifyrepository.Update(id, verify)
	return err1
}
func (service verifyService) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
		success, failure := r.Verifyrepository.Delete(id)
		return success, failure
}
