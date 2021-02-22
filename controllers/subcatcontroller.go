package controllers

import(
	//"fmt"	
	"net/http"
	"github.com/labstack/echo"
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model"
	"github.com/myrachanto/ecommerce/service"
)
 //SubcategoryController ...
var (
	SubcategoryController subcategoryController = subcategoryController{}
)
type subcategoryController struct{ }
/////////controllers/////////////////
func (controller subcategoryController) Create(c echo.Context) error {
	subcategory := &model.Subcategory{}
	if err := c.Bind(subcategory); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	err1 := service.SubcategoryService.Create(subcategory)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, "created successifuly")
}

func (controller subcategoryController) GetAll(c echo.Context) error {
	subcategorys := []model.Subcategory{}
	subcategorys, err3 := service.SubcategoryService.GetAll(subcategorys)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, subcategorys)
} 
func (controller subcategoryController) GetOne(c echo.Context) error {
	id := string(c.Param("id"))
	subcategory, problem := service.SubcategoryService.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, subcategory)	
}

func (controller subcategoryController) Update(c echo.Context) error {
	subcategory :=  &model.Subcategory{}
	if err := c.Bind(subcategory); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	id := string(c.Param("id"))
	problem := service.SubcategoryService.Update(id, subcategory)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusCreated, "Updated successifuly")
}

func (controller subcategoryController) Delete(c echo.Context) error {
	id := string(c.Param("id"))
	success, failure := service.SubcategoryService.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}