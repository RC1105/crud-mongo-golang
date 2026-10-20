package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString = "mongodb://localhost:27017"
var dbName = "school"
var colName = "students"

var collection *mongo.Collection

type Student struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" bson:"name,omitempty"`
	Age   int                `json:"age,omitempty" bson:"age,omitempty"`
	Marks int                `json:"marks,omitempty" bson:"marks,omitempty"`
}

func init() {
	// client option
	clientOption := options.Client().
		ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(
		context.TODO(),
		clientOption,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	// collection instance
	collection = client.
		Database(dbName).
		Collection(colName)

	fmt.Println("Collection instance is ready")
}
func main() {
	lst := []Student{}

	/*st := Student{
		Name:  "Vimal",
		Age:   20,
		Marks: 99,
	}

	_, err := collection.InsertOne(context.TODO(), st)
	if err != nil {
		fmt.Print("failed")
	} else {
		fmt.Print("Insert Successful")
	}
	fmt.Print("Finally")*/
	//collection.UpdateMany(context.TODO(), bson.M{"age": bson.M{"$lte": 30, "$gt": 20}}, bson.M{"$set": bson.M{"age": 25}})
	collection.DeleteOne(context.TODO(), bson.M{"age": 25, "name": "Raghav"})
	res, err1 := collection.Find(context.TODO(), bson.M{"age": bson.M{"$lte": 30, "$gt": 20}})
	if err1 != nil {
		log.Fatal("Error in retrieving")
	} else {
		res.All(context.TODO(), &lst)
	}
	fmt.Print(lst)

}
func main1() {
	fmt.Println("Application started")
	//ls := []Student{}
	/*for i := 0; i < 3; i++ {
		stud := Student{
			Name:  "Vimal" + string(i),
			Age:   i*10 + 7,
			Marks: 90 - i,
		}
		ls = append(ls, stud)
	}
	collection.InsertOne(context.TODO(), ls[0])*/
	filter := bson.M{"age": 24}
	collection.UpdateOne(context.TODO(), filter, bson.M{"$set": bson.M{"age": 25}})
	/*stud := Student{
		Name:  "Vijay",
		Age:   28,
		Marks: 90,
	}*/

	/*result, err := collection.InsertOne(
		context.TODO(),
		stud,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Student inserted successfully!")
	fmt.Println("Inserted ID:", result.InsertedID)*/
	lstStudents := []bson.M{}
	res, _ := collection.Find(context.TODO(), bson.M{"age": 25})
	res.All(context.TODO(), &lstStudents)
	fmt.Println(lstStudents)
	//	upd := bson.M{"$set": bson.M{"Pass": "Yes"}}
	/*
		_, err := collection.DeleteMany(context.TODO(), bson.M{"marks": nil})
		if err == nil {
			fmt.Print("check")
		}*/

}
