package controllers

import(
	"fmt"	
	"net/http"
	"github.com/labstack/echo"
	"github.com/myrachanto/asokomonolith/httperrors"
	"github.com/myrachanto/asokomonolith/model"
	"github.com/myrachanto/asokomonolith/service"
)
 
var (
	UserController userController = userController{}
)
type userController struct{ }
/////////controllers/////////////////
func (controller userController) Create(c echo.Context) error {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	err1 := service.UserService.Create(user)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	} 
	return c.JSON(http.StatusCreated, "created successifuly")
}
func (controller userController) Login(c echo.Context) error {
	user := &model.LoginUser{}
	auth := &model.Auth{}
	if err := c.Bind(user); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	
	auth, problem := service.UserService.Login(user)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, auth)	
}
func (controller userController) Logout(c echo.Context) error {
	token := string(c.Param("token"))
	_,problem := service.UserService.Logout(token)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, "succeessifully logged out")	
}
func (controller userController) GetAll(c echo.Context) error {
	Users := []model.User{}
	users, err3 := service.UserService.GetAll(Users)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, users)
} 
func (controller userController) GetOne(c echo.Context) error {
	id := string(c.Param("id"))
	user, problem := service.UserService.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, user)	
}

func (controller userController) Update(c echo.Context) error {
	user :=  &model.User{}
	if err := c.Bind(user); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	fmt.Println(user)
	id := string(c.Param("id"))
	problem := service.UserService.Update(id, user)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusCreated, "Updated successifuly")
}
func (controller userController) AUpdate(c echo.Context) error {
	user :=  &model.User{}
	if err := c.Bind(user); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	fmt.Println(user)
	id := string(c.Param("id"))
	problem := service.UserService.AUpdate(id, user)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusCreated, "Updated successifuly")
}
func (controller userController) Delete(c echo.Context) error {
	id := string(c.Param("id"))
	success, failure := service.UserService.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}