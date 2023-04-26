package module

import (
	"context"
	"fmt"
	"os"

	"github.com/Fatwaff/be_tugbes/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

// var MongoInfo = atdb.DBInfo{
// 	DBString: MongoString,
// 	DBName:   "db_tugbes",
// }

// var MongoConn = atdb.MongoConnect(MongoInfo)

func MongoConnect() *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database("db_tugbes")
}

func InsertOneDoc(db *mongo.Database, col string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(col).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func GetUserFromEmail(email string, db *mongo.Database, col string) (result model.User) {
	table := db.Collection(col)
	filter := bson.M{"email": email}
	err := table.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Printf("GetUserFromEmail: %v\n", err)
	}
	return result
}

func GetAllData(db *mongo.Database, col string, data interface{}) interface{} {
	table := db.Collection(col)
	filter := bson.M{}
	cursor, err := table.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllData :", col, err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

// func InsertUser(db *mongo.Database, col string, doc interface{}) (insertedID primitive.ObjectID, err error) {
// 	result, err := db.Collection(col).InsertOne(context.Background(), doc)
// 	if err != nil {
// 		fmt.Printf("InsertPresensi: %v\n", err)
// 		return
// 	}
// 	insertedID = result.InsertedID.(primitive.ObjectID)
// 	return insertedID, nil
// }

// func InsertUser(db *mongo.Database, col string, firstname string, lastname string, email string, password string) (InsertedID interface{}) {
// 	var field model.User
// 	field.FirstName = firstname
// 	field.LastName = lastname
// 	field.Email = email
// 	field.Password = password
// 	return InsertOneDoc(db, col, field)
// }

// func GetAllPresensiFromKehadiran(kehadiran string, db *mongo.Database, col string) (aprs []model.Presensi) {
// 	presensi := db.Collection(col)
// 	filter := bson.M{"kehadiran": kehadiran}
// 	cursor, err := presensi.Find(context.TODO(), filter)
// 	if err != nil {
// 		fmt.Printf("GetAllPresensiFromKehadiran: %v\n", err)
// 	}
// 	err = cursor.All(context.TODO(), &aprs)
// 	if err != nil{
// 		fmt.Println(err)
// 	}
// 	return 
// }

// func GetAllMataKuliah(db *mongo.Database, col string) (data []model.MataKuliah) {
// 	dataMataKuliah := db.Collection(col)
// 	filter := bson.M{}
// 	cursor, err := dataMataKuliah.Find(context.TODO(), filter)
// 	if err != nil {
// 		fmt.Println("GetAllMataKuliah :", err)
// 	}
// 	err = cursor.All(context.TODO(), &data)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return data
// }
