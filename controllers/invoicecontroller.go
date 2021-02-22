package controllers

import(
	//"fmt"	
	"net/http"
	"github.com/labstack/echo"
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model"
	"github.com/myrachanto/ecommerce/service"
)
 //InvoiceController ...
var (
	InvoiceController invoiceController = invoiceController{}
)
type invoiceController struct{ }
/////////controllers/////////////////
func (controller invoiceController) Create(c echo.Context) error {
	invoice := &model.Invoice{}
	if err := c.Bind(invoice); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	err1 := service.InvoiceService.Create(invoice)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, "created successifuly")
}

func (controller invoiceController) GetAll(c echo.Context) error {
	invoices := []model.Invoice{}
	invoices, err3 := service.InvoiceService.GetAll(invoices)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, invoices)
} 
func (controller invoiceController) GetOne(c echo.Context) error {
	id := string(c.Param("id"))
	invoice, problem := service.InvoiceService.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, invoice)	
}

// func (controller invoiceController) Update(c echo.Context) error {
// 	invoice :=  &model.Invoice{}
// 	if err := c.Bind(invoice); err != nil {
// 		httperror := httperrors.NewBadRequestError("Invalid json body")
// 		return c.JSON(httperror.Code, httperror)
// 	}	
// 	id := string(c.Param("id"))
// 	problem := service.InvoiceService.Update(id, invoice)
// 	if problem != nil {
// 		return c.JSON(problem.Code, problem)
// 	}
// 	return c.JSON(http.StatusCreated, "Updated successifuly")
// }

// func (controller invoiceController) Delete(c echo.Context) error {
// 	id := string(c.Param("id"))
// 	success, failure := service.InvoiceService.Delete(id)
// 	if failure != nil {
// 		return c.JSON(failure.Code, failure)
// 	}
// 	return c.JSON(success.Code, success)
		
// }