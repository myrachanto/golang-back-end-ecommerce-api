package service

import (
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model" 
	r "github.com/myrachanto/ecommerce/repository"
)

var (
	InvoiceService  = invoiceService{}
)

type invoiceService struct {
}

func (service invoiceService) Create(invoice *model.Invoice) (*httperrors.HttpError) {
	err1 := r.Invoicerepository.Create(invoice)
	 return err1

}

func (service invoiceService) GetOne(id string) (*model.Invoice, *httperrors.HttpError) {
	invoice, err1 := r.Invoicerepository.GetOne(id)
	return invoice, err1
}

func (service invoiceService) GetAll(invoices []model.Invoice) ([]model.Invoice, *httperrors.HttpError) {
	invoices, err := r.Invoicerepository.GetAll(invoices)
	return invoices, err
}

// func (service invoiceService) Update(id string, invoice *model.Invoice) (*httperrors.HttpError) {
// 	err1 := r.Invoicerepository.Update(id, invoice)
// 	return err1
// }
// func (service invoiceService) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
// 		success, failure := r.Invoicerepository.Delete(id)
// 		return success, failure
// }
