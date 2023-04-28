package module

import (
	"context"
	"fmt"
	"os"

	"github.com/Fatwaff/be_tugbes/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect() *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database("db_tugbes")
}

func GetUserFromEmail(email string, db *mongo.Database, col string) (result model.User) {
	collection := db.Collection(col)
	filter := bson.M{"email": email}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Printf("GetUserFromEmail: %v\n", err)
	}
	return result
}

func GetAllDocs(db *mongo.Database, col string, docs interface{}) interface{} {
	collection := db.Collection(col)
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error GetAllDocs in colecction", col, ":", err)
	}
	err = cursor.All(context.TODO(), &docs)
	if err != nil {
		fmt.Println(err)
	}
	return docs
}

func InsertOneDoc(db *mongo.Database, col string, doc interface{}) (insertedID primitive.ObjectID, err error) {
	result, err := db.Collection(col).InsertOne(context.Background(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func EmailExists(db *mongo.Database, col string, email string) bool {
	collection := db.Collection(col)
	filter := bson.M{"email": email}
	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		fmt.Printf("Error EmailExists : %v\n", err)
	}
	return count > 0
}

func SignUp(db *mongo.Database, col string, doc model.User) (insertedID primitive.ObjectID, err error) {
	if doc.FirstName == "" || doc.LastName == "" || doc.Email == "" || doc.Password == "" {
		return insertedID, fmt.Errorf("mohon untuk melengkapi data")
	} 
	if doc.Email == GetUserFromEmail(doc.Email, db, col).Email {
		return insertedID, fmt.Errorf("email sudah terdaftar")
	} 
	return InsertOneDoc(db, col, doc)
}

// func SignUp(db *mongo.Database, col string, firstname string, lastname string, email string, password string) (insertedID primitive.ObjectID, err error) {
// 	var doc model.User
// 	collection := db.Collection(col)
// 	filter := bson.M{"email": email}
// 	err = collection.FindOne(context.TODO(), filter).Decode(&doc)
// 	var er = fmt.Errorf("mongo: no documents in result")
// 	if err != er && err != nil {
// 		fmt.Printf("Gagal Sign Up : %v\n", err)
// 		return
// 	}
// 	if doc.Email == email {
// 		return insertedID, nil
// 	}
// 	doc.FirstName = firstname
// 	doc.LastName = lastname
// 	doc.Email = email
// 	doc.Password = password
// 	return InsertOneDoc(db, col, doc)
// }