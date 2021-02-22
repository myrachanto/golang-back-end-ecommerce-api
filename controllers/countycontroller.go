package controllers

import(
	//"fmt"	
	"net/http"
	"github.com/labstack/echo"
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model"
	"github.com/myrachanto/ecommerce/service"
)
 //CountyController ...
var (
	CountyController countyController = countyController{}
)
type countyController struct{ }
/////////controllers/////////////////
func (controller countyController) Create(c echo.Context) error {
	county := &model.County{}
	if err := c.Bind(county); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	err1 := service.CountyService.Create(county)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, "created successifuly")
}

func (controller countyController) GetAll(c echo.Context) error {
	countys := []model.County{}
	countys, err3 := service.CountyService.GetAll(countys)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, countys)
} 
func (controller countyController) GetOne(c echo.Context) error {
	id := string(c.Param("id"))
	county, problem := service.CountyService.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, county)	
}

func (controller countyController) Update(c echo.Context) error {
	county :=  &model.County{}
	if err := c.Bind(county); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	id := string(c.Param("id"))
	problem := service.CountyService.Update(id, county)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusCreated, "Updated successifuly")
}

func (controller countyController) Delete(c echo.Context) error {
	id := string(c.Param("id"))
	success, failure := service.CountyService.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}