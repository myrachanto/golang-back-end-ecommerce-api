package controllers

import(
	//"fmt"	
	"net/http"
	"github.com/labstack/echo"
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model"
	"github.com/myrachanto/ecommerce/service"
)
 //NortificationController ...
var (
	NortificationController nortificationController = nortificationController{}
)
type nortificationController struct{ }
/////////controllers/////////////////
func (controller nortificationController) Create(c echo.Context) error {
	nortification := &model.Nortification{}
	if err := c.Bind(nortification); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	err1 := service.NortificationService.Create(nortification)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, "created successifuly")
}

func (controller nortificationController) GetAll(c echo.Context) error {
	nortifications := []model.Nortification{}
	nortifications, err3 := service.NortificationService.GetAll(nortifications)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, nortifications)
} 
func (controller nortificationController) GetOne(c echo.Context) error {
	id := string(c.Param("id"))
	nortification, problem := service.NortificationService.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, nortification)	
}

func (controller nortificationController) Update(c echo.Context) error {
	nortification :=  &model.Nortification{}
	if err := c.Bind(nortification); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	id := string(c.Param("id"))
	problem := service.NortificationService.Update(id, nortification)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusCreated, "Updated successifuly")
}

func (controller nortificationController) Delete(c echo.Context) error {
	id := string(c.Param("id"))
	success, failure := service.NortificationService.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}