package booking

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance
// config mongo
const (
	dbName   = "pttrainer"
	mongoURI = "mongodb+srv://pttrainer:NmSJWnnBmApV5sEu@maincluster.gkfe6.mongodb.net/pttrainer?authSource=admin&replicaSet=atlas-u555p7-shard-0&w=majority&readPreference=primary&appname=MongoDB%20Compass&retryWrites=true&ssl=true"
	// mongoURI = "mongodb://localhost:27017/" + dbName
)

// status timeslots
const (
	tsActive  = "active"
)

// status bookings
const (
	bkPending = "pending"
)

type Booking struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ClientID   primitive.ObjectID `json:"bClientId,omitempty" bson:"bClientId,omitempty"`
	TimeSlotID primitive.ObjectID `json:"bTimeSlotId,omitempty" bson:"bTimeSlotId,omitempty"`
	Status     string             `json:"bStatus,omitempty" bson:"bStatus,omitempty"`
	CreatedAt  time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt  time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
