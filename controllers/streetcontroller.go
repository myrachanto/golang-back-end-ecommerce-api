package controllers

import(
	//"fmt"	
	"net/http"
	"github.com/labstack/echo"
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model"
	"github.com/myrachanto/ecommerce/service"
)
//StreetController ..
var (
	StreetController streetController = streetController{}
)
type streetController struct{ }
/////////controllers/////////////////
func (controller streetController) Create(c echo.Context) error {
	street := &model.Street{}
	if err := c.Bind(street); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	err1 := service.StreetService.Create(street)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, "created successifuly")
}

func (controller streetController) GetAll(c echo.Context) error {
	streets := []model.Street{}
	streets, err3 := service.StreetService.GetAll(streets)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, streets)
} 
func (controller streetController) GetOne(c echo.Context) error {
	id := string(c.Param("id"))
	street, problem := service.StreetService.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, street)	
}

func (controller streetController) Update(c echo.Context) error {
	street :=  &model.Street{}
	if err := c.Bind(street); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	id := string(c.Param("id"))
	problem := service.StreetService.Update(id, street)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusCreated, "Updated successifuly")
}

func (controller streetController) Delete(c echo.Context) error {
	id := string(c.Param("id"))
	success, failure := service.StreetService.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}