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
	IndustryController industryController = industryController{}
)
type industryController struct{ }
/////////controllers/////////////////
func (controller industryController) Create(c echo.Context) error {
	industry := &model.Industry{}
	if err := c.Bind(industry); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	 err1 := service.IndustryService.Create(industry)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, "created successifuly")
}

func (controller industryController) GetAll(c echo.Context) error {
	industrys := []model.Industry{}
	industrys, err3 := service.IndustryService.GetAll(industrys)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, industrys)
} 
func (controller industryController) GetOne(c echo.Context) error {
	id := string(c.Param("id"))
	industry, problem := service.IndustryService.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, industry)	
}

func (controller industryController) Update(c echo.Context) error {
	industry :=  &model.Industry{}
	if err := c.Bind(industry); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	id := string(c.Param("id"))
	problem := service.IndustryService.Update(id, industry)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusCreated, "Updated successifuly")
}

func (controller industryController) Delete(c echo.Context) error {
	id := string(c.Param("id"))
	success, failure := service.IndustryService.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}