package controllers

import(
	"fmt"	
	"net/http"
	"github.com/labstack/echo"
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model"
	"github.com/myrachanto/ecommerce/service"
)
 //CustomerController ...
var (
	CustomerController customerController = customerController{}
)
type customerController struct{ }
/////////controllers/////////////////
func (controller customerController) Create(c echo.Context) error {
	customer := &model.Customer{}
	if err := c.Bind(customer); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	fmt.Println(customer)
	err1 := service.CustomerService.Create(customer)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, "created successifuly")
}

func (controller customerController) GetAll(c echo.Context) error {
	Customers := []model.Customer{}
	customers, err3 := service.CustomerService.GetAll(Customers)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, customers)
} 
func (controller customerController) GetOne(c echo.Context) error {
	id := string(c.Param("id"))
	customer, problem := service.CustomerService.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, customer)	
}

func (controller customerController) Update(c echo.Context) error {
	customer :=  &model.Customer{}
	if err := c.Bind(customer); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	fmt.Println(customer)
	id := string(c.Param("id"))
	problem := service.CustomerService.Update(id, customer)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusCreated, "Updated successifuly")
}

func (controller customerController) Delete(c echo.Context) error {
	id := string(c.Param("id"))
	success, failure := service.CustomerService.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}