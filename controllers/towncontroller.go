package controllers

import(
	//"fmt"	
	"net/http"
	"github.com/labstack/echo"
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model"
	"github.com/myrachanto/ecommerce/service"
)
 //TownController ..
var (
	TownController townController = townController{}
)
type townController struct{ }
/////////controllers/////////////////
func (controller townController) Create(c echo.Context) error {
	town := &model.Town{}
	if err := c.Bind(town); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	err1 := service.TownService.Create(town)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, "created successifuly")
}

func (controller townController) GetAll(c echo.Context) error {
	towns := []model.Town{}
	towns, err3 := service.TownService.GetAll(towns)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, towns)
} 
func (controller townController) GetOne(c echo.Context) error {
	id := string(c.Param("id"))
	town, problem := service.TownService.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, town)	
}

func (controller townController) Update(c echo.Context) error {
	town :=  &model.Town{}
	if err := c.Bind(town); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	id := string(c.Param("id"))
	problem := service.TownService.Update(id, town)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusCreated, "Updated successifuly")
}

func (controller townController) Delete(c echo.Context) error {
	id := string(c.Param("id"))
	success, failure := service.TownService.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}