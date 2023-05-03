package module

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/Fatwaff/be_tugbes/model"
	"github.com/badoux/checkmail"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/argon2"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect() *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database("db_tugbes")
}

func GetUserFromEmail(email string, db *mongo.Database, col string) (result model.User, err error) {
	collection := db.Collection(col)
	filter := bson.M{"email": email}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result, fmt.Errorf("email tidak ditemukan")
		}
		return result, fmt.Errorf("kesalahan server")
	}
	return result, nil
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
		// fmt.Printf("InsertOneDoc: %v\n", err)
		return insertedID, fmt.Errorf("kesalahan server")
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

// func EmailExists(db *mongo.Database, col string, email string) bool {
// 	collection := db.Collection(col)
// 	filter := bson.M{"email": email}
// 	count, err := collection.CountDocuments(context.Background(), filter)
// 	if err != nil {
// 		fmt.Printf("Error EmailExists : %v\n", err)
// 	}
// 	return count > 0
// }

func SignUp(db *mongo.Database, col string, insertedDoc model.User) (insertedID primitive.ObjectID, err error) {
	if insertedDoc.FirstName == "" || insertedDoc.LastName == "" || insertedDoc.Email == "" || insertedDoc.Password == "" {
		return insertedID, fmt.Errorf("mohon untuk melengkapi data")
	} 
	if err = checkmail.ValidateFormat(insertedDoc.Email); err != nil {
		return insertedID, fmt.Errorf("email tidak valid")
	} 
	userExists, _ := GetUserFromEmail(insertedDoc.Email, db, col)
	if insertedDoc.Email == userExists.Email {
		return insertedID, fmt.Errorf("email sudah terdaftar")
	} 
	if insertedDoc.Confirmpassword != insertedDoc.Password {
		return insertedID, fmt.Errorf("konfirmasi password salah")
	}
	if strings.Contains(insertedDoc.Password, " ") {
		return insertedID, fmt.Errorf("password tidak boleh mengandung spasi")
	}
	if len(insertedDoc.Password) < 8 {
		return insertedID, fmt.Errorf("password terlalu pendek")
	}
	salt := make([]byte, 16)
	_, err = rand.Read(salt)
	if err != nil {
		return insertedID, fmt.Errorf("kesalahan server")
	}
	hashedPassword := argon2.IDKey([]byte(insertedDoc.Password), salt, 1, 64*1024, 4, 32)
	insertedDoc.Password = hex.EncodeToString(hashedPassword)
	insertedDoc.Salt = hex.EncodeToString(salt)
	insertedDoc.Confirmpassword = ""
	return InsertOneDoc(db, col, insertedDoc)
}

func LogIn(db *mongo.Database, col string, insertedDoc model.User) (userName string, err error) {
	if insertedDoc.Email == "" || insertedDoc.Password == "" {
		return userName, fmt.Errorf("mohon untuk melengkapi data")
	} 
	if err = checkmail.ValidateFormat(insertedDoc.Email); err != nil {
		return userName, fmt.Errorf("email tidak valid")
	} 
	existsDoc, err := GetUserFromEmail(insertedDoc.Email, db, col)
	if err != nil {
		return 
	}
	salt, err := hex.DecodeString(existsDoc.Salt)
	if err != nil {
		return userName, fmt.Errorf("kesalahan server")
	}
	hash := argon2.IDKey([]byte(insertedDoc.Password), salt, 1, 64*1024, 4, 32)
	if hex.EncodeToString(hash) != existsDoc.Password {
		// fmt.Println("insert :", hex.EncodeToString(hash))
		// fmt.Println("exist :", existsDoc.Password)
		// fmt.Println("salt :", salt)
		return userName, fmt.Errorf("password salah")
	}
	return existsDoc.FirstName + " " + existsDoc.LastName, nil
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