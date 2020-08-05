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
	MajorcategoryController majorcategoryController = majorcategoryController{}
)
type majorcategoryController struct{ }
/////////controllers/////////////////
func (controller majorcategoryController) Create(c echo.Context) error {
	majorcategory := &model.Majorcategory{}
	if err := c.Bind(majorcategory); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	err1 := service.MajorcategoryService.Create(majorcategory)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, "created successifuly")
}

func (controller majorcategoryController) GetAll(c echo.Context) error {
	majorcategorys := []model.Majorcategory{}
	majorcategorys, err3 := service.MajorcategoryService.GetAll(majorcategorys)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, majorcategorys)
} 
func (controller majorcategoryController) GetOne(c echo.Context) error {
	id := string(c.Param("id"))
	majorcategory, problem := service.MajorcategoryService.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, majorcategory)	
}

func (controller majorcategoryController) Update(c echo.Context) error {
	majorcategory :=  &model.Majorcategory{}
	if err := c.Bind(majorcategory); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	id := string(c.Param("id"))
	problem := service.MajorcategoryService.Update(id, majorcategory)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusCreated, "Updated successifuly")
}

func (controller majorcategoryController) Delete(c echo.Context) error {
	id := string(c.Param("id"))
	success, failure := service.MajorcategoryService.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}