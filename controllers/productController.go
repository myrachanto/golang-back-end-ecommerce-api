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
	ProductController productController = productController{}
)
type productController struct{ }
/////////controllers/////////////////
func (controller productController) Create(c echo.Context) error {
	product := &model.Product{}
	if err := c.Bind(product); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	err1 := service.ProductService.Create(product)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, "created successifuly")
}

func (controller productController) GetAll(c echo.Context) error {
	products := []model.Product{}
	products, err3 := service.ProductService.GetAll(products)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, products)
} 
func (controller productController) GetOne(c echo.Context) error {
	id := string(c.Param("id"))
	product, problem := service.ProductService.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, product)	
}

func (controller productController) Update(c echo.Context) error {
	product :=  &model.Product{}
	if err := c.Bind(product); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	id := string(c.Param("id"))
	problem := service.ProductService.Update(id, product)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusCreated, "Updated successifuly")
}

func (controller productController) Delete(c echo.Context) error {
	id := string(c.Param("id"))
	success, failure := service.ProductService.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}