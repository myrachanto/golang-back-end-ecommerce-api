package repository

import (
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
	jwt "github.com/dgrijalva/jwt-go"
    "go.mongodb.org/mongo-driver/bson"
		"github.com/myrachanto/ecommerce/httperrors"
		"github.com/myrachanto/ecommerce/model" 
)
//Userrepository repository
var (
	Userrepository userrepository = userrepository{}
)

type userrepository struct{}

func (r *userrepository) Create(user *model.User) (*httperrors.HttpError) {
	if err1 := user.Validate(); err1 != nil {
		return err1
	}
	ok, err1 := user.ValidatePassword(user.Password)
	if !ok {
		return err1
	}
	ok = user.ValidateEmail(user.Email)
	if !ok {
		return httperrors.NewNotFoundError("Your email format is wrong!")
	}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	filter := bson.M{"email": user.Email,}
	auser := &model.User{}
	collection1 := db.Collection("user")
	err4 := collection1.FindOne(ctx, filter).Decode(&auser)
	if err4 == nil {
		return  httperrors.NewBadRequestError(fmt.Sprintf("User with this email exist @ - , %d", err4))
	}
	hashpassword, err2 := user.HashPassword(user.Password)
	if err2 != nil {
		return err2
	}
	user.Password = hashpassword
	
	collection := db.Collection("user")
	_, err := collection.InsertOne(ctx, user)
		if err != nil {
			return httperrors.NewBadRequestError(fmt.Sprintf("Create user Failed, %d", err))
	}
	DbClose(c)
	return nil
}

func (r *userrepository) Login(user *model.LoginUser) (*model.Auth, *httperrors.HttpError) {
	if err := user.Validate(); err != nil {
		return nil,err
	}
	c, t := Mongoclient();if t != nil {
		return nil,t
	}
	db, e := Mongodb();if e != nil {
		return nil,e
	}
	collection := db.Collection("user")
	filter := bson.M{"email": user.Email,}
	auser := &model.User{}
	err := collection.FindOne(ctx, filter).Decode(&auser)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("User with this email exist @ - , %d", err))
	}
	ok := user.Compare(user.Password, auser.Password)
	if !ok {
		return nil, httperrors.NewNotFoundError("wrong email password combo!")
	}
	tk := &model.Token{
		UserID: auser.Id,
		UName: auser.UName,
		Admin: auser.Admin,
		Supervisor:  auser.Supervisor,
		Employee:  auser.Employee,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: model.ExpiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	err2 := godotenv.Load()
	if err2 != nil {
		log.Fatal("Error loading key")
	}
	encyKey := os.Getenv("EncryptionKey")
	tokenString, error := token.SignedString([]byte(encyKey))
	if error != nil {
		fmt.Println(error)
	}
	auths := &model.Auth{UserID:auser.Id, Admin:auser.Admin, UName:auser.UName, Supervisor:auser.Supervisor,Employee:auser.Employee, Token:tokenString}
  //  fmt.Println(auths)
	_, err3 := collection.InsertOne(ctx, auths)
		if err3 != nil {
			return nil,httperrors.NewBadRequestError(fmt.Sprintf("Create user Failed, %d", err))
	}
	
	// filter1 := bson.D{}
	// auth := &model.Auth{}
	// collection2 := db.Collection("auth")
	// err4 := collection2.FindOne(ctx, filter1).Decode(&auth)
	// fmt.Println(filter1)
	// if err4 != nil {
	// 	fmt.Println(err4)
	// 	return nil, httperrors.NewBadRequestError("something went wrong authorizing!")
	// }
	DbClose(c)
	return auths, nil
}
func (r *userrepository) Logout(token string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("auth")
	filter1 := bson.M{"token": token,}
	_, err3 := collection.DeleteOne(ctx, filter1)
	if err3 != nil {
		return nil, httperrors.NewBadRequestError("something went wrong login out!")
	}
	DbClose(c)
 return httperrors.NewSuccessMessage("something went wrong login out!"), nil
}
func (r *userrepository) GetOne(id string) (user *model.User, errors *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("user")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	DbClose(c)
	return user, nil	
}

func (r *userrepository) GetAll(users []model.User) ([]model.User, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("user")
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
	err := cur.Decode(&users)
		if err != nil { 
			return nil,	httperrors.NewNotFoundError("Error while decoding results!")
		}
	// do something with result....
	}
	if err := cur.Err(); err != nil {
		return nil,	httperrors.NewNotFoundError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}	
	DbClose(c)
    return users, nil

}

func (r *userrepository) Update(id string, user *model.User) (*httperrors.HttpError) {
	uuser := &model.User{}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	ok, err1 := user.ValidatePassword(user.Password)
	if !ok {
		return err1
	}
	ok = user.ValidateEmail(user.Email)
	if !ok {
		return httperrors.NewNotFoundError("Your email format is wrong!")
	}
	hashpassword, err2 := user.HashPassword(user.Password)
	if err2 != nil {
		return err2
	}
	user.Password = hashpassword
	collection := db.Collection("user")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&uuser)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}

	if user.FName  == "" {
		user.FName = uuser.FName
	}
	if user.LName  == "" {
		user.LName = uuser.LName
	}
	if user.UName  == "" {
		user.UName = uuser.UName
	}
	if user.Phone  == "" {
		user.Phone = uuser.Phone
	}
	if user.Address  == "" {
		user.Address = uuser.Address
	}
	if user.Picture  == "" {
		user.Picture = uuser.Picture
	}
	//////////////////////////////////////////////////////////////////////
	/////////////////get the admin authorisation to handle this///////////////
	if user.Admin  == true {
		user.Admin = uuser.Admin
	}
	if user.Supervisor  == true {
		user.Supervisor = uuser.Supervisor
	}
	if user.Employee  == true {
		user.Employee = uuser.Employee
	}
	if user.Email  == "" {
		user.Email = uuser.Email
	}
	if hashpassword == "" {
		user.Password = uuser.Password
	}
	_, err = collection.UpdateOne(ctx, filter, user)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Update of user Failed, %d", err))
	} 
	DbClose(c)
	return nil
}

func (r *userrepository) AUpdate(id string, user *model.User) (*httperrors.HttpError) {
	uuser := &model.User{}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	ok, err1 := user.ValidatePassword(user.Password)
	if !ok {
		return err1
	}
	ok = user.ValidateEmail(user.Email)
	if !ok {
		return httperrors.NewNotFoundError("Your email format is wrong!")
	}
	hashpassword, err2 := user.HashPassword(user.Password)
	if err2 != nil {
		return err2
	}
	user.Password = hashpassword
	collection := db.Collection("user")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&uuser)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	if user.FName  == "" {
		user.FName = uuser.FName
	}
	if user.LName  == "" {
		user.LName = uuser.LName
	}
	if user.UName  == "" {
		user.UName = uuser.UName
	}
	if user.Phone  == "" {
		user.Phone = uuser.Phone
	}
	if user.Address  == "" {
		user.Address = uuser.Address
	}
	if user.Picture  == "" {
		user.Picture = uuser.Picture
	}
	//////////////////////////////////////////////////////////////////////
	/////////////////get the admin authorisation to handle this///////////////
	if !user.Admin {
		user.Admin = uuser.Admin
	}
	if !user.Supervisor {
		user.Supervisor = uuser.Supervisor
	}
	if !user.Employee {
		user.Employee = uuser.Employee
	}
	if user.Email  == "" {
		user.Email = uuser.Email
	}
	if hashpassword == "" {
		user.Password = uuser.Password
	}
	_, err = collection.UpdateOne(ctx, filter, user)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Update of user Failed, %d", err))
	} 
	DbClose(c)
	return nil
}
func (r userrepository) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("user")
	filter := bson.M{"_id": id}
	ok, err := collection.DeleteOne(ctx, filter)
	if ok == nil {
		return nil, httperrors.NewNotFoundError(fmt.Sprintf("deletion of %d failed", err))
	}
	DbClose(c)
		return httperrors.NewSuccessMessage("deleted successfully"), nil
	
}
