package service

import (
	"github.com/myrachanto/asokomonolith/httperrors"
	"github.com/myrachanto/asokomonolith/model" 
	r "github.com/myrachanto/asokomonolith/repository"
)

var (
	UserService  = userService{}
)

type userService struct {
}

func (service userService) Create(user *model.User) (*httperrors.HttpError) {
	err1 := r.Userrepository.Create(user)
	 return err1

}

func (service userService) Login(auser *model.LoginUser) (*model.Auth, *httperrors.HttpError) {
	user, err1 :=  r.Userrepository.Login(auser)
	if err1 != nil {
		return nil, err1
	}
	return user, nil
}
func (service userService) Logout(token string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
	success, failure := r.Userrepository.Logout(token)
	return success, failure
}
func (service userService) GetOne(id string) (*model.User, *httperrors.HttpError) {
	user, err1 := r.Userrepository.GetOne(id)
	return user, err1
}

func (service userService) GetAll(users []model.User) ([]model.User, *httperrors.HttpError) {
	users, err := r.Userrepository.GetAll(users)
	return users, err
}

func (service userService) Update(id string, user *model.User) (*httperrors.HttpError) {
	err1 := r.Userrepository.Update(id, user)
	return err1
}

func (service userService) AUpdate(id string, user *model.User) (*httperrors.HttpError) {
	err1 := r.Userrepository.AUpdate(id, user)
	return err1
}
func (service userService) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
		success, failure := r.Userrepository.Delete(id)
		return success, failure
}
