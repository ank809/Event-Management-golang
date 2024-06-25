package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Event struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `json:"name"`
	DateTime  string             `json:"datetime"`
	Venue     string             `json:"venue"`
	Attendees int                `json:"attendees"`
	Organizer string             `json:"organizer"`
}
