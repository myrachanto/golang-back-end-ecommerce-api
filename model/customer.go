package model

import (
	"time"
	"github.com/myrachanto/ecommerce/httperrors"
)

type Customer struct {
	Id      string    `json:"_id" bson:"_id"`
	Name    string    `bson:"name"`
	Company string    `bson:"company"`
	Phone   string    `bson:"phone"`
	Address string    `bson:"address"`
	Email   string    `bson:"email;unique"`
	Invoice []Invoice `bson:"invoices"` //has many invoices
	Base
}

type Invoice struct {
	CustomerID  uint64        `bson:"customer_id"`
	User    User      
	Title       string        `bson:"title"`
	Dated       time.Time     `bson:"date"`
	Due_date    time.Time     `bson:"due_date"`
	Discount    float64       `bson:"discount"`
	Sub_total   float64       `bson:"sub_total"`
	Tax   float64       `bson:"tax"`
	Total       float64       `bson:"total"`
	Balance 	float64
	GracePeriod *time.Time
	ExpiryDate *time.Time
	Nortify    *time.Time
	Facilitator string
	Base
}

func (customer Customer) Validate() *httperrors.HttpError {
	if customer.Name == "" {
		return httperrors.NewNotFoundError("Invalid Name")
	}
	if customer.Company == "" {
		return httperrors.NewNotFoundError("Invalid Company")
	}
	if customer.Phone == "" {
		return httperrors.NewNotFoundError("Invalid Phone")
	}
	if customer.Email == "" {
		return httperrors.NewNotFoundError("Invalid Email")
	}
	if customer.Address == "" {
		return httperrors.NewNotFoundError("Invalid Address")
	}
	return nil
}
// func (invoice Invoice) Validate() *httperrors.HttpError {
// 	if invoice.User.Id == "" {
// 		return httperrors.NewNotFoundError("Invalid User")
// 	}
// 	if invoice.Title == "" {
// 		return httperrors.NewNotFoundError("Invalid Title")
// 	}
// 	if invoice.Total == "" {
// 		return httperrors.NewNotFoundError("Invalid Total")
// 	}
// 	if invoice.ExpiryDate == "" {
// 		return httperrors.NewNotFoundError("Invalid Email")
// 	}
// 	if invoice.Address == "" {
// 		return httperrors.NewNotFoundError("Invalid Address")
// 	}
// 	return nil
// }