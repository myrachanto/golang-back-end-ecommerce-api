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
	VerifyController verifyController = verifyController{}
)
type verifyController struct{ }
/////////controllers/////////////////
func (controller verifyController) Create(c echo.Context) error {
	verify := &model.Verify{}
	if err := c.Bind(verify); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	err1 := service.VerifyService.Create(verify)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, "created successifuly")
}

func (controller verifyController) GetAll(c echo.Context) error {
	verifys := []model.Verify{}
	verifys, err3 := service.VerifyService.GetAll(verifys)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, verifys)
} 
func (controller verifyController) GetOne(c echo.Context) error {
	id := string(c.Param("id"))
	verify, problem := service.VerifyService.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, verify)	
}

func (controller verifyController) Update(c echo.Context) error {
	verify :=  &model.Verify{}
	if err := c.Bind(verify); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	id := string(c.Param("id"))
	problem := service.VerifyService.Update(id, verify)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusCreated, "Updated successifuly")
}

func (controller verifyController) Delete(c echo.Context) error {
	id := string(c.Param("id"))
	success, failure := service.VerifyService.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}