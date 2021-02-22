package controllers

import(
	"fmt"	
	"net/http"
	"os"
	"io"
	"github.com/labstack/echo"
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model"
	"github.com/myrachanto/ecommerce/service"
)
//UserController ...
var (
	UserController userController = userController{}
)
type userController struct{ }
/////////controllers/////////////////
func (controller userController) Create(c echo.Context) error {

	user := &model.User{}
	user.FName = c.FormValue("fname")
	user.LName = c.FormValue("lname")
	user.UName = c.FormValue("uname")
	user.Phone = c.FormValue("phone")
	user.Address = c.FormValue("address")
	user.Email = c.FormValue("email")
	user.Password = c.FormValue("password")
	// user.Business = c.FormValue("business")

	pic, err2 := c.FormFile("picture")
	if pic != nil {
		//    fmt.Println(pic.Filename)
			if err2 != nil {
					httperror := httperrors.NewBadRequestError("Invalid picture")
					return c.JSON(httperror.Code, err2)
				}	
			src, err := pic.Open()
			if err != nil {
				httperror := httperrors.NewBadRequestError("the picture is corrupted")
				return c.JSON(httperror.Code, err)
			}	
			defer src.Close()
			// filePath := "./public/imgs/users/"
			filePath := "./public/imgs/users/" + pic.Filename
			filePath1 := "/imgs/users/" + pic.Filename
			// Destination
			dst, err4 := os.Create(filePath)
			if err4 != nil {
				httperror := httperrors.NewBadRequestError("the Directory mess")
				return c.JSON(httperror.Code, err4)
			}
			defer dst.Close()
			// Copy
			if _, err = io.Copy(dst, src); err != nil {
				if err2 != nil {
					httperror := httperrors.NewBadRequestError("error filling")
					return c.JSON(httperror.Code, httperror)
				}
			}
			
		user.Picture = filePath1
		err1 := service.UserService.Create(user)
		if err1 != nil {
			return c.JSON(err1.Code, err1)
		}
		if _, err = io.Copy(dst, src); err != nil {
			if err2 != nil {
				httperror := httperrors.NewBadRequestError("error filling")
				return c.JSON(httperror.Code, httperror)
			}
		}
		return c.JSON(http.StatusCreated, "user created succesifully")
	}
	err1 := service.UserService.Create(user)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	} 
	return c.JSON(http.StatusCreated, "user created succesifully")
}
func (controller userController) Login(c echo.Context) error {
	user := &model.LoginUser{}
	auth := &model.Auth{}
	
	user.Email = c.FormValue("email")
	user.Password = c.FormValue("password")
	auth, problem := service.UserService.Login(user)
	if problem != nil {
		fmt.Println(problem)
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