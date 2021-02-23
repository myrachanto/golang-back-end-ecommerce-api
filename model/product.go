package model

import "github.com/myrachanto/ecommerce/httperrors"

type Product struct {
	Name          string   `json:"name,omitempty"`
	Title         string   `json:"title,omitempty"`
	Description   string   `json:"description,omitempty"`
	Code          string   `json:"code,omitempty"`
	Majorcategory string   `json:"majorcat,omitempty"`
	Category      string   `json:"category,omitempty"`
	Subcategory   string   `json:"subcategory,omitempty"`
	Price         float64  `json:"price,omitempty"`
	Picture       string   `json:"picture,omitempty"`
	Tag           []Tag    `json:"tag,omitempty"`
	Rates         []Rating `json:"rates,omitempty"`
	Featured      bool     `json:"featured,omitempty"`
	Promotion     bool     `json:"promotion,omitempty"`
	Hotdeals      bool     `json:"hotdeals,omitempty"`
	Base          `json:"base,omitempty"`
}

func (product Product) Validate() *httperrors.HttpError {
	if product.Name == "" {
		return httperrors.NewNotFoundError("Invalid Name")
	}
	if product.Title == "" {
		return httperrors.NewNotFoundError("Invalid title")
	}
	if product.Description == "" {
		return httperrors.NewNotFoundError("Invalid Description")
	}
	return nil
}
