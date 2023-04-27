package Fatwa_1214038

import (
	"fmt"
	"testing"

	"github.com/Fatwaff/be_tugbes/model"
	"github.com/Fatwaff/be_tugbes/module"
)

// func TestInserOneDoc(t *testing.T) {
// 	var data model.User
// 	data.FirstName = "Fatwa Fatahillah"
// 	data.LastName = "Fatah"
// 	data.Email = "fatwa@gmail.com"
// 	data.Password = "jklbnm"
// 	hasil := module.InsertOneDoc(module.MongoConnect(), "user", data)
// 	fmt.Println(hasil)
// }

// Batas suci

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

 func TestInserUser(t *testing.T) {
 	var data model.User
	data.FirstName = "Naufal Dekha"
	data.LastName = "Widana"
	data.Email = "nopal@gmail.com"
	data.Password = "jklbnm"
 	hasil, err := module.InsertOneDoc(module.MongoConnect(), "user", data)
 	if err != nil {
 		t.Errorf("Error inserting data: %v", err)
 	}
 	fmt.Println("Data berhasil disimpan dengan id :", hasil.Hex())
}

