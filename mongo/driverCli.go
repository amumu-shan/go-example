package _mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	mongoURI = "mongodb://admin:123456@127.0.0.1:27017"
)

type User struct {
	FullName string `bson:"fullName"`
	Age      int    `bson:"age"`
}

type mongoCollection struct {
	client *mongo.Collection
}

var UserDB *mongoCollection

func New() error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		fmt.Println(err)
		return nil
	}
	usersCollection := client.Database("test").Collection("users")
	UserDB = &mongoCollection{client: usersCollection}
	return err
}

func (m *mongoCollection) Insert() {

	user := bson.D{{"fullName", "User1"}, {"age", 30}}

	result, err := m.client.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("inserted id=", result.InsertedID)

	users := []interface{}{
		bson.D{{"fullName", "User2"}, {"age", 25}},
		bson.D{{"fullName", "User3"}, {"age", 20}},
		bson.D{{"fullName", "User4"}, {"age", 28}},
	}
	results, err := m.client.InsertMany(context.TODO(), users)
	if err != nil {
		fmt.Println(results)
		return
	}
	fmt.Println("inserted many id=", result.InsertedID)
}

func (m *mongoCollection) Query() {
	filter := bson.D{
		{
			"$and",
			bson.A{
				bson.D{
					{"age", bson.D{{"$gt", 25}}},
				},
			},
		},
	}
	cursor, err := m.client.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []bson.M
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		panic(err)
	}

	for _, result := range results {
		fmt.Println(result)
	}
	fmt.Println("===findOne")
	var result bson.M
	if err = m.client.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println(result)

}
func (m *mongoCollection) UpdateById() {
	user := bson.D{{"fullName", "User 5"}, {"age", 22}}
	insertResult, err := m.client.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	update := bson.D{
		{"$set", bson.D{{"fullName", "User five"}}},
		{"$inc", bson.D{{"age", 1}}},
	}
	updateResult, err := m.client.UpdateByID(context.TODO(), insertResult.InsertedID, update)
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of documents updated:", updateResult.ModifiedCount)
}
func (m *mongoCollection) UpdateByOne() {
	filter := bson.D{
		{
			"$and",
			bson.A{
				bson.D{
					{"age", bson.D{{"$gt", 25}, {"$lt", 40}}},
				},
			},
		},
	}
	update := bson.D{
		{"$set", bson.D{{"age", 40}}},
	}
	updateResult, err := m.client.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of documents updated:", updateResult.ModifiedCount)
}
func (m *mongoCollection) UpdateByMany() {
	filter := bson.D{
		{
			"$and",
			bson.A{
				bson.D{
					{"age", bson.D{{"$gt", 25}, {"$lt", 50}}},
				},
			},
		},
	}
	update := bson.D{
		{"$set", bson.D{{"age", 50}}},
	}
	updateResult, err := m.client.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of documents updated:", updateResult.ModifiedCount)
}
func (m *mongoCollection) ReplaceOne() {
	filter := bson.D{{"fullName", "User2"}}
	replacement := bson.D{
		{"firstName", "json"},
		{"lastName", "Doe"},
		{"age", 30},
		{"emailAddress", "ps@qq.com"},
	}
	result, err := m.client.ReplaceOne(context.TODO(), filter, replacement)
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of documents replaced:", result.ModifiedCount)
}
func (m *mongoCollection) DeleteOne() {
	filter := bson.D{
		{"$and", bson.A{
			bson.D{
				{"age", bson.D{{"$gt", 25}}},
			},
		}},
	}
	result, err := m.client.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of documents deleted:", result.DeletedCount)
}
