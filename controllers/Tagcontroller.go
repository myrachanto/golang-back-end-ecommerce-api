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
	TagController tagController = tagController{}
)
type tagController struct{ }
/////////controllers/////////////////
func (controller tagController) Create(c echo.Context) error {
	tag := &model.Tag{}
	if err := c.Bind(tag); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	err1 := service.TagService.Create(tag)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, "created successifuly")
}

func (controller tagController) GetAll(c echo.Context) error {
	tags := []model.Tag{}
	tags, err3 := service.TagService.GetAll(tags)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, tags)
} 
func (controller tagController) GetOne(c echo.Context) error {
	id := string(c.Param("id"))
	tag, problem := service.TagService.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, tag)	
}

func (controller tagController) Update(c echo.Context) error {
	tag :=  &model.Tag{}
	if err := c.Bind(tag); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	id := string(c.Param("id"))
	problem := service.TagService.Update(id, tag)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusCreated, "Updated successifuly")
}

func (controller tagController) Delete(c echo.Context) error {
	id := string(c.Param("id"))
	success, failure := service.TagService.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}