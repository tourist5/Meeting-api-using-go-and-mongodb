package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type Meeting struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"title" bson:"title,omitempty"`
	Participants *Participants`json:"participants" bson:"participants,omitempty"`
	Start Time string          `json:"start time" bson:"start time,omitempty"`
	End Time string          `json:"end time" bson:"start time,omitempty"`
	Timestamp string          `json:"timestamp" bson:"timestamp,omitempty"`
}

type Participants struct {
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	email  string `json:"email,omitempty" bson:"email,omitempty"`
	RSVP   string `json:"rsvp,omitempty" bson:"rsvp,omitempty"`
}