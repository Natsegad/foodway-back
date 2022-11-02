package db

import (
	"fmt"
	"foodway/internal/cfg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserInfo struct {
	Id           uint32 `json:"id" gorm:"primaryKey"`
	Phone        string `json:"phone"`
	Token        string `json:"jwt"`
	RefreshToken string `json:"refresh_jwt"`
}

var DataBase *gorm.DB

func InitDb() {
	dbURL := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s", cfg.Cfg.DbUser, cfg.Cfg.DbPassword, cfg.Cfg.DbPort, cfg.Cfg.DbName)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error initialize DB %s \n", err.Error())
		return
	}
	fmt.Println("DB initialize!")
	DataBase = db
}

func AutoMigrate() error {
	err := DataBase.AutoMigrate(&UserInfo{})
	if err != nil {
		fmt.Printf("Error migrate DB %s \n", err.Error())
		return err
	}

	return nil
}

func AddUser(user UserInfo) error {
	ok := DataBase.Create(&user)
	if ok.Error != nil {
		fmt.Printf("Error AddUser %s \n", ok.Error.Error())
		return ok.Error
	}

	fmt.Printf("User added to database %s \n", user.Phone)
	return nil
}

func DeleteUser(id uint64) error {
	user := UserInfo{}
	ok := DataBase.Find(&user, id)
	if ok.Error != nil {
		fmt.Printf("Error AddUser %s \n", ok.Error.Error())
		return ok.Error
	}

	DataBase.Delete(&user)

	return nil
}

func GetUser(id uint32) (UserInfo, error) {
	user := UserInfo{}
	ok := DataBase.Find(&user, id)
	if ok.Error != nil {
		fmt.Printf("Error AddUser %s \n", ok.Error.Error())
		return user, ok.Error
	}

	return user, nil
}

func GetAllUsers() ([]UserInfo, error) {
	var users []UserInfo
	ok := DataBase.Find(&users)
	if ok.Error != nil {
		fmt.Printf("Error AddUser %s \n", ok.Error.Error())
		return nil, ok.Error
	}

	return users, nil
}

func GetUserByPhone(phone string) UserInfo {
	ret := UserInfo{}
	users, err := GetAllUsers()
	if err != nil {
		fmt.Printf("Error GetAllUsers in GetUserByEmail \n")
		return ret
	}

	for _, v := range users {
		if v.Phone == phone {
			return v
		}
	}

	return ret
}

func UpdateUser(user UserInfo) error {
	userBase := UserInfo{}
	ok := DataBase.Find(&userBase, user.Id)
	if ok.Error != nil {
		fmt.Printf("Error AddUser %s \n", ok.Error.Error())
		return ok.Error
	}

	userBase.Phone = user.Phone
	userBase.Token = user.Token

	DataBase.Save(&userBase)

	return nil
}
