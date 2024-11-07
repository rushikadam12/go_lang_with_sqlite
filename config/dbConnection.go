package config

import (
	"crud_with_sqlite/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

type DB struct {
	*gorm.DB
}

func NewDB() *DB {
	db, err := gorm.Open("sqlite3", "/home/rushikeshkadam/code_space/golang_with_sqlite/test.db")

	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.User{})

	return &DB{db}
}

func (db *DB) GetUsers()([]models.User,error){
	var users []models.User

	err:=db.Find(&users).Error;
	if err!=nil{
		return nil,err
	}
	return users,nil
}
func (db *DB)CreateUser(user *models.User)error{

	err:=db.Create(user).Error;

	if err!=nil{
		return err
	}

	return nil
}