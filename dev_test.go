package Fatwa_1214038

import (
	"fmt"
	"testing"

	"github.com/Fatwaff/be_tugbes/model"
	"github.com/Fatwaff/be_tugbes/module"
)

func TestInserUser(t *testing.T) {
	var data model.User
	data.FirstName = "Naufal Dekha"
	data.LastName = "Widana"
	data.Email = "nopal@gmail.com"
	data.Password = "jklbnm"
	hasil := module.InsertOneDoc(module.MongoConnect(), "user", data)
	fmt.Println(hasil)
}

func TestGetUserFromEmail(t *testing.T) {
	email := "dimas@gmail.com"
	data := module.GetUserFromEmail(email, module.MongoConnect(), "user")
	fmt.Println(data)
}

func TestGetAllDataUser(t *testing.T) {
	var user []model.User
	data := module.GetAllData(module.MongoConnect(), "user", user)
	fmt.Println(data)
}

// func TestInserUser(t *testing.T) {
// 	data := bson.M{
// 		"firstname":    "Naufal Dehka",
// 		"lastname":     "Widana",
// 		"email":     	"nopal@gmail.com",
// 		"password": 	"jklbnm",
// 	}
// 	hasil, err := module.InsertUser(module.MongoConnect(), "user", data)
// 	if err != nil {
// 		t.Errorf("Error inserting data: %v", err)
// 	}
// 	fmt.Println("Data berhasil disimpan dengan id :", hasil.Hex())
// }

// func TestInserUser(t *testing.T) {
// 	firstname := "Fatwa Fatahillah"
// 	lastname := "Fatah"
// 	email := "fatwa@gmail.com"
// 	password := "qwerty"
// 	hasil := module.InsertUser(module.MongoConnect(), "user", firstname, lastname, email, password)
// 	fmt.Println(hasil)
// }

// func TestGetAllPresensiFromKehadiran(t *testing.T) {
// 	kehadiran := "masuk"
// 	data := module.GetAllPresensiFromKehadiran(kehadiran, module.MongoConn, "presensi")
// 	fmt.Println(data)
// }

// func TestGetAllMataKuliah(t *testing.T) {
// 	data := module.GetAllMataKuliah(module.MongoConn, "matkul")
// 	fmt.Println(data)
// }
