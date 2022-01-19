package schemas

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Avatar   string             `bson:"avatar,omitempty"`
}
