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
	ShopController shopController = shopController{}
)
type shopController struct{ }
/////////controllers/////////////////
func (controller shopController) Create(c echo.Context) error {
	shop := &model.Shop{}
	if err := c.Bind(shop); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	err1 := service.ShopService.Create(shop)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, "created successifuly")
}

func (controller shopController) GetAll(c echo.Context) error {
	shops := []model.Shop{}
	shops, err3 := service.ShopService.GetAll(shops)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, shops)
} 
func (controller shopController) GetOne(c echo.Context) error {
	id := string(c.Param("id"))
	shop, problem := service.ShopService.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, shop)	
}

func (controller shopController) Update(c echo.Context) error {
	shop :=  &model.Shop{}
	if err := c.Bind(shop); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	id := string(c.Param("id"))
	problem := service.ShopService.Update(id, shop)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusCreated, "updated successifuly")
}

func (controller shopController) Delete(c echo.Context) error {
	id := string(c.Param("id"))
	success, failure := service.ShopService.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}