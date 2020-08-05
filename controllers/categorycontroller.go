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
	CategoryController categoryController = categoryController{}
)
type categoryController struct{ }
/////////controllers/////////////////
func (controller categoryController) Create(c echo.Context) error {
	category := &model.Category{}
	if err := c.Bind(category); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	err1 := service.CategoryService.Create(category)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, "created successifuly")
}

func (controller categoryController) GetAll(c echo.Context) error {
	categorys := []model.Category{}
	categorys, err3 := service.CategoryService.GetAll(categorys)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, categorys)
} 
func (controller categoryController) GetOne(c echo.Context) error {
	id := string(c.Param("id"))
	category, problem := service.CategoryService.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, category)	
}

func (controller categoryController) Update(c echo.Context) error {
	category :=  &model.Category{}
	if err := c.Bind(category); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	id := string(c.Param("id"))
	problem := service.CategoryService.Update(id, category)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusCreated, "updated successifuly")
}

func (controller categoryController) Delete(c echo.Context) error {
	id := string(c.Param("id"))
	success, failure := service.CategoryService.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}