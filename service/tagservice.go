package service

import (
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model" 
	r "github.com/myrachanto/ecommerce/repository"
)

var (
	TagService  = tagService{}
)

type tagService struct {
}

func (service tagService) Create(tag *model.Tag) (*httperrors.HttpError) {
	err1 := r.Tagrepository.Create(tag)
	 return err1

}

func (service tagService) GetOne(id string) (*model.Tag, *httperrors.HttpError) {
	tag, err1 := r.Tagrepository.GetOne(id)
	return tag, err1
}

func (service tagService) GetAll(tags []model.Tag) ([]model.Tag, *httperrors.HttpError) {
	tags, err := r.Tagrepository.GetAll(tags)
	return tags, err
}

func (service tagService) Update(id string, tag *model.Tag) (*httperrors.HttpError) {
	err1 := r.Tagrepository.Update(id, tag)
	return err1
}
func (service tagService) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
		success, failure := r.Tagrepository.Delete(id)
		return success, failure
}
