package models

type User struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	Username     string             `json:"username" bson:"username" `
	Password     string             `json:"password" bson:"password"`
	Email        string             `json:"email" bson:"email"`
}
