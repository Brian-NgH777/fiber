package booking

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(dbName)
	if err != nil {
		return err
	}

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}

func CreateBooking(ctx context.Context, booking *Booking) (*Booking, error) {
	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"bTimeSlotId": 1}, Options: opt}
	collectionBooking := mg.Db.Collection("bookings")
	if _, err := collectionBooking.Indexes().CreateOne(ctx, index); err != nil {
		return nil, err
	}
	collectionTimeslots := mg.Db.Collection("timeslots")

	err := mg.Client.UseSession(ctx, func(sctx mongo.SessionContext) error {
		err := sctx.StartTransaction(options.Transaction().
			SetReadConcern(readconcern.Snapshot()).
			SetWriteConcern(writeconcern.New(writeconcern.WMajority())),
		)
		if err != nil {
			return err
		}

		var result TimeSlot
		err = collectionTimeslots.FindOne(
			ctx,
			bson.M{"_id": booking.TimeSlotID}).Decode(&result)
		if err != nil {
			sctx.AbortTransaction(sctx)
			return err
		}

		count, err := collectionBooking.CountDocuments(ctx, bson.M{"bTimeSlotId": booking.TimeSlotID})
		if err != nil {
			sctx.AbortTransaction(sctx)
			return err
		}

		if count >= result.MaxClient {
			sctx.AbortTransaction(sctx)
			return errors.New("Class is ffan ull")
		}

		booking.CreatedAt = time.Now().UTC()
		booking.UpdatedAt = time.Now().UTC()
		booking.Status = bkPending

		_, err = collectionBooking.InsertOne(ctx, booking)
		if err != nil {
			sctx.AbortTransaction(sctx)
			return err
		}

		if result.Status != tsActive && result.Status != tsCompleted {
			data, err := collectionTimeslots.UpdateOne(
				ctx,
				bson.M{"_id": booking.TimeSlotID},
				bson.D{
					{"$set", bson.D{{"tStatus", tsActive}}},
					{"$set", bson.D{{"updatedAt", time.Now().UTC()}}},
				})
			if err != nil || data.ModifiedCount == 0 {
				sctx.AbortTransaction(sctx)
				return err
			}
		}


		for {
			err = sctx.CommitTransaction(sctx)
			switch e := err.(type) {
			case nil:
				return nil
			case mongo.CommandError:
				if e.HasErrorLabel("UnknownTransactionCommitResult") {
					log.Println("UnknownTransactionCommitResult, retrying commit operation...")
					continue
				}
				log.Println("Error during commit...")
				return e
			default:
				log.Println("Error during commit...")
				return e
			}
		}
	})

	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "bTimeSlotId", Value: booking.TimeSlotID}}
	createdRecord := collectionBooking.FindOne(ctx, filter)

	createdBooking := &Booking{}
	createdRecord.Decode(createdBooking)

	return createdBooking, nil
}
