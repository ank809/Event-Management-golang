package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Event struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `json:"name"`
	DateTime  primitive.DateTime `json:"datetime"`
	Venue     string             `json:"venue"`
	Attendess int                `json:"attendees"`
}
