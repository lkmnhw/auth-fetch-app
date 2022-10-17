package databases

import (
	"context"

	_ "github.com/go-sql-driver/mysql" //sql driver
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database holds database connection
type Database struct {
	MClient *mongo.Client
	MDBase  *mongo.Database
	User    string
}

// Connect to database
func Connect(ctx context.Context, mongodb, database, user string) (*Database, error) {
	clientOptions := options.Client().ApplyURI(mongodb)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return &Database{}, errors.Wrap(err, "Failed to connect MongoDB database")
	}
	base := client.Database(database)

	return &Database{
			MClient: client,
			MDBase:  base,
			User:    user,
		},
		nil
}

// Close the database connection
func (db Database) Close() error {
	return db.MClient.Disconnect(context.Background())
}

// FindOne using object id from database
func (db *Database) FindOne(ctx context.Context, col string, filter bson.M) (*mongo.SingleResult, error) {
	r := db.MClient.Database(db.MDBase.Name()).Collection(col).FindOne(ctx, filter)
	return r, nil
}

// InsertOne to database
func (db *Database) InsertOne(ctx context.Context, col string, data interface{}) (*mongo.InsertOneResult, error) {
	r, err := db.MClient.Database(db.MDBase.Name()).Collection(col).InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	return r, err
}
