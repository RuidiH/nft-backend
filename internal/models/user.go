// backend/internal/models/user.go

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User represents a user entity
type User struct {
    ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name string             `bson:"name" json:"name"`
}
