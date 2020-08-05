package controllers

import(
	//"fmt"	
	"net/http"
	"github.com/labstack/echo"
	"github.com/myrachanto/asokomonolith/httperrors"
	"github.com/myrachanto/asokomonolith/model"
	"github.com/myrachanto/asokomonolith/service"
)
 
var (
	DivisionController divisionController = divisionController{}
)
type divisionController struct{ }
/////////controllers/////////////////
func (controller divisionController) Create(c echo.Context) error {
	division := &model.Division{}
	if err := c.Bind(division); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	err1 := service.DivisionService.Create(division)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, "created successifuly")
}

func (controller divisionController) GetAll(c echo.Context) error {
	divisions := []model.Division{}
	divisions, err3 := service.DivisionService.GetAll(divisions)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, divisions)
} 
func (controller divisionController) GetOne(c echo.Context) error {
	id := string(c.Param("id"))
	division, problem := service.DivisionService.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, division)	
}

func (controller divisionController) Update(c echo.Context) error {
	division :=  &model.Division{}
	if err := c.Bind(division); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	id := string(c.Param("id"))
	problem := service.DivisionService.Update(id, division)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusCreated, "Updated successifuly")
}

func (controller divisionController) Delete(c echo.Context) error {
	id := string(c.Param("id"))
	success, failure := service.DivisionService.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}