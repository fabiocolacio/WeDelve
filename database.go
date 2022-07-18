package main

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrNoModifications error = errors.New("No records were modified")
	ErrMissingField    error = errors.New("A required field is missing")
)

type Database struct {
	client *mongo.Client
	db     *mongo.Database
	users  *mongo.Collection
}

type RecordId interface{}

func OpenDatabase(uri string) (Database, error) {
	var (
		db  Database
		err error
	)

	ctx := context.TODO()
	opts := options.Client().ApplyURI(uri)

	if db.client, err = mongo.Connect(ctx, opts); err != nil {
		return db, err
	}

	db.db = db.client.Database("wedelve")
	db.users = db.db.Collection("users")

	return db, nil
}

func (db Database) Close() error {
	return db.client.Disconnect(context.TODO())
}

func (db Database) InsertUser(user User) (RecordId, error) {
	var (
		result *mongo.InsertOneResult
		id     RecordId
		err    error
	)

	user.Password = ""

	if result, err = db.users.InsertOne(context.TODO(), user); err == nil {
		id = result.InsertedID
	}

	return id, err
}

func (db Database) UpdateChallenge(user User) (RecordId, error) {
	var (
		result *mongo.UpdateResult
		id     RecordId
		err    error
	)

	if user.Challenge == nil {
		return id, ErrMissingField
	}

	filter := bson.M{"name": user.Name}
	update := bson.D{{"$set", bson.M{"challenge": user.Challenge}}}

	if result, err = db.users.UpdateOne(context.TODO(), filter, update); err == nil {
		id = result.UpsertedID
		if result.ModifiedCount < 1 {
			err = ErrNoModifications
		}
	}

	return id, err
}

func (db Database) GetUserByName(name string) (user User, err error) {
	result := db.users.FindOne(context.TODO(), bson.M{"name": name})
	err = result.Decode(&user)
	return
}
