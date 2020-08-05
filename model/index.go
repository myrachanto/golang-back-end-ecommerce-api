package model

import(
	"time"
)
type Base struct{
	Created_At time.Time `bson:"created_at"`
	Updated_At time.Time `bson:"updated_at"`
	Delete_At *time.Time `bson:"deleted_at"`

}