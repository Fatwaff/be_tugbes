package Fatwa_1214038

import (
	"fmt"
	"testing"

	"github.com/Fatwaff/be_tugbes/model"
	"github.com/Fatwaff/be_tugbes/module"
)

var db = module.MongoConnect()

func TestGetUserFromEmail(t *testing.T) {
	email := "abc@gmail.com"
	hasil := module.GetUserFromEmail(email, db, "user")
	fmt.Println(hasil)
}

func TestGetAllDoc(t *testing.T) {
	var docs []model.Tess
	hasil := module.GetAllDocs(db, "usgf", docs)
	fmt.Println(hasil)
}

func TestInsertOneDoc(t *testing.T) {
 	var doc model.User
	doc.FirstName = "Naufal Dekha"
	doc.LastName = "Widana"
	doc.Email = "nopal@gmail.com"
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

func TestSignUp(t *testing.T) {
	var doc model.User
	doc.FirstName = "Farel Nouval"
	doc.LastName = "Widana"
	doc.Email = "farel@gmail.com"
	doc.Password = "fghjkl"
	if doc.FirstName == "" || doc.LastName == "" || doc.Email == "" || doc.Password == "" {
		t.Errorf("mohon untuk melengkapi data")
	} else if module.EmailExists(db, "user", doc.Email){
		t.Errorf("Email sudah terdaftar")
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

func TestSignUp2(t *testing.T) {
	var doc model.User
	doc.FirstName = "Farel Nouval"
	doc.LastName = "Widana"
	doc.Email = "abc@gmail.com"
	doc.Password = "fghjkl"
	insertedID, err := module.SignUp(db, "user", doc)
	if err != nil {
		t.Errorf("Error inserting document: %v", err)
	} else {
	fmt.Println("Data berhasil disimpan dengan id :", insertedID.Hex())
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