package Fatwa_1214038

import (
	"fmt"
	"testing"

	"github.com/Fatwaff/be_tugbes/model"
	"github.com/Fatwaff/be_tugbes/module"
)

var db = module.MongoConnect()

func TestGetUserFromEmail(t *testing.T) {
	email := "fatwaff@gmail.com"
	hasil, err := module.GetUserFromEmail(email, db, "user")
	if err != nil {
		t.Errorf("Error TestGetUserFromEmail: %v", err)
	} else {
		fmt.Println(hasil)
	}
}

func TestGetAllDoc(t *testing.T) {
	var docs []model.User
	hasil := module.GetAllDocs(db, "user", docs)
	fmt.Println(hasil)
}

func TestInsertOneDoc(t *testing.T) {
 	var doc model.User
	doc.FirstName = "Fatwa Fatahillah"
	doc.LastName = "Fatah"
	doc.Email = "fatwahh@gmail.com"
	doc.Password = "jklbnm"
	if doc.FirstName == "" || doc.LastName == "" || doc.Email == "" || doc.Password == "" {
		t.Errorf("mohon untuk melengkapi data")
	} else {
		insertedID, err := module.InsertOneDoc(db, "user", doc)
		if err != nil {
			t.Errorf("Error inserting document: %v", err)
			fmt.Println("Data tidak berhasil disimpan")
		} else {
		fmt.Println("Data berhasil disimpan dengan id :", insertedID.Hex())
		}
	}
}

// func TestSignUp(t *testing.T) {
// 	var doc model.User
// 	doc.FirstName = "Farel Nouval"
// 	doc.LastName = "Widana"
// 	doc.Email = "farel@gmail.com"
// 	doc.Password = "fghjkl"
// 	if doc.FirstName == "" || doc.LastName == "" || doc.Email == "" || doc.Password == "" {
// 		t.Errorf("mohon untuk melengkapi data")
// 	} else if module.EmailExists(db, "user", doc.Email){
// 		t.Errorf("Email sudah terdaftar")
// 	} else {
// 		insertedID, err := module.InsertOneDoc(db, "user", doc)
// 		if err != nil {
// 			t.Errorf("Error inserting document: %v", err)
// 			fmt.Println("Data tidak berhasil disimpan")
// 		} else {
// 		fmt.Println("Data berhasil disimpan dengan id :", insertedID.Hex())
// 		}
// 	}
// }

func TestSignUp(t *testing.T) {
	var doc model.User
	doc.FirstName = "Fatwa Fatahillah"
	doc.LastName = "Fatah"
	doc.Email = "fatwaff@gmail.com"
	doc.Password = "fghjklio"
	doc.Confirmpassword = "fghjklio"
	insertedID, err := module.SignUp(db, "user", doc)
	if err != nil {
		t.Errorf("Error inserting document: %v", err)
	} else {
	fmt.Println("Data berhasil disimpan dengan id :", insertedID.Hex())
	}
}

func TestLogIn(t *testing.T) {
	var doc model.User
	doc.Email = "fatwaff@gmail.com"
	doc.Password = "fghjklio"
	user, err := module.LogIn(db, "user", doc)
	if err != nil {
		t.Errorf("Error getting document: %v", err)
	} else {
		fmt.Println("Welcome :", user)
	}
}

// func TestSignUp(t *testing.T) {
// 	firstName := "Farel Naufal"
// 	lastName := "Daswara"
// 	email := "farel@gmail.com"
// 	password := "iopjkl"
//  	insertedID, err := module.SignUp(module.MongoConnect(), "user", firstName, lastName, email, password)
// 	var er = fmt.Errorf("mongo: no documents in result")
//  	if err == er && err != nil {
// 		fmt.Println(err)
// 		fmt.Println(er)
//  		t.Errorf("Error inserting document: %v", err)
//  	} 
// 	if insertedID.Hex() == "000000000000000000000000" && err == nil{
// 		fmt.Println("Email sudah terdaftar")
// 	} else {
//  	fmt.Println("Data berhasil disimpan dengan id :", insertedID.Hex())
// 	}
// }