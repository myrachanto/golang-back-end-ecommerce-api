package controllers

import(
	//"fmt"	
	"net/http"
	"github.com/labstack/echo"
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model"
	"github.com/myrachanto/ecommerce/service"
)
 //RatingController ...
var (
	RatingController ratingController = ratingController{}
)
type ratingController struct{ }
/////////controllers/////////////////
func (controller ratingController) Create(c echo.Context) error {
	rating := &model.Rating{}
	if err := c.Bind(rating); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	err1 := service.RatingService.Create(rating)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, "created successifuly")
}

func (controller ratingController) GetAll(c echo.Context) error {
	ratings := []model.Rating{}
	ratings, err3 := service.RatingService.GetAll(ratings)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, ratings)
} 
func (controller ratingController) GetOne(c echo.Context) error {
	id := string(c.Param("id"))
	rating, problem := service.RatingService.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, rating)	
}

func (controller ratingController) Update(c echo.Context) error {
	rating :=  &model.Rating{}
	if err := c.Bind(rating); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	id := string(c.Param("id"))
	problem := service.RatingService.Update(id, rating)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusCreated, "Updated successifuly")
}

func (controller ratingController) Delete(c echo.Context) error {
	id := string(c.Param("id"))
	success, failure := service.RatingService.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}