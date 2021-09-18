package booking

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
}


func(mg *MongoInstance) CountUser(ctx context.Context, user *User) (int64, error) {
	collectionUser := mg.Db.Collection("users")
	count, err := collectionUser.CountDocuments(ctx, bson.M{"_id": user.ID})
	if err != nil {
		return 0, err
	}

	return count, nil
}
