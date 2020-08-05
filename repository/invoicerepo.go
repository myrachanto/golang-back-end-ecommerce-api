package repository

import (
	"fmt"
    "go.mongodb.org/mongo-driver/bson"
	"github.com/myrachanto/asokomonolith/httperrors"
	"github.com/myrachanto/asokomonolith/model" 
)
var (
	Invoicerepository invoicerepository = invoicerepository{}
)

type invoicerepository struct{}

func (r *invoicerepository) Create(invoice *model.Invoice) (*httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("invoice")
	_, err := collection.InsertOne(ctx, invoice)
		if err != nil {
			return httperrors.NewBadRequestError(fmt.Sprintf("Create invoice Failed, %d", err))
	}
	DbClose(c)
	return nil
}

func (r *invoicerepository) GetOne(id string) (invoice *model.Invoice, errors *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("invoice")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&invoice)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	DbClose(c)
	return invoice, nil	
}

func (r *invoicerepository) GetAll(invoices []model.Invoice) ([]model.Invoice, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("invoice")
	filter := bson.M{}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	if err != nil { 
		return nil,	httperrors.NewNotFoundError("no results found")
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
	err := cur.Decode(&invoices)
		if err != nil { 
			return nil,	httperrors.NewNotFoundError("Error while decoding results!")
		}
	// do something with result....
	}
	if err := cur.Err(); err != nil {
		return nil,	httperrors.NewNotFoundError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}	
	DbClose(c)
    return invoices, nil

}

// func (r *invoicerepository) Update(id string, invoice *model.Invoice) (*httperrors.HttpError) {
// 	uinvoice := &model.Invoice{}
// 	c, t := Mongoclient();if t != nil {
// 		return t
// 	}
// 	db, e := Mongodb();if e != nil {
// 		return e
// 	}
// 	collection := db.Collection("invoice")
// 	filter := bson.M{"_id": id}
// 	err := collection.FindOne(ctx, filter).Decode(&uinvoice)
// 	if err != nil {
// 		return httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
// 	}
// 	if invoice.Name  == "" {
// 		invoice.Name = uinvoice.Name
// 	}
// 	if invoice.Title  == "" {
// 		invoice.Title = uinvoice.Title
// 	}

// 	if invoice.Description  == "" {
// 		invoice.Description = uinvoice.Description
// 	}
// 	_, err = collection.UpdateOne(ctx, filter, invoice)
// 	if err != nil {
// 		return httperrors.NewBadRequestError(fmt.Sprintf("Update of invoice Failed, %d", err))
// 	} 
// 	return nil
// }
// func (r invoicerepository) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
// 	c, t := Mongoclient();if t != nil {
// 		return nil, t
// 	}
// 	db, e := Mongodb();if e != nil {
// 		return nil, e
// 	}
// 	collection := db.Collection("invoice")
// 	filter := bson.M{"_id": id}
// 	ok, err := collection.DeleteOne(ctx, filter)
// 	if ok == nil {
// 		return nil, httperrors.NewNotFoundError(fmt.Sprintf("deletion of %d failed", err))
// 	}else{
// 		return httperrors.NewSuccessMessage("deleted successfully"), nil
// 	}
// }
