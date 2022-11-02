package service

import (
	"errors"
	"fmt"
	"foodway/internal/domain"
	"foodway/pkg/db"
	"foodway/pkg/jwt"
	"foodway/pkg/logger"
	"github.com/google/uuid"
)

// Валидация. Проверка данных соответствию стандартов
func validation(user domain.UserInfoPhone) error {
	log := logger.GetLogger()

	err := user.Validation()
	if err != nil {
		log.Errorf("Error Validation user: phone = %s. Error: %s\n", user.Phone, err.Error())
		return err
	}

	return nil
}

// Добавление в бд
func addUserToDB(user domain.UserInfoPhone) error {
	return db.AddUser(NewUserInfo(user.Phone))
}

// checkHaveUser проверка того что указанный номер есть в базе
func checkHaveUser(user domain.UserInfoPhone) error {
	ok := db.GetUserByPhone(user.Phone)
	if ok.Id != 0 {
		return errors.New(fmt.Sprintf("%s user have", user.Phone))
	}

	return nil
}

// Главная функция регистрации
func Registration(user domain.UserInfoPhone) error {
	log := logger.GetLogger()

	log.Infof("Start registration user %s ", user.Phone)

	// Первый этап валидация данных
	err := validation(user)
	if err != nil {
		log.Errorf("Validation user: %s error \n", user.Phone)
		return err
	}

	err = checkHaveUser(user)
	if err != nil {
		return err
	}

	// Проходит валидация идет добавление в базу данных
	err = addUserToDB(user)
	if err != nil {
		log.Errorf("Add user: %s to data base error \n", user.Phone)
		return err
	}

	return nil
}

// NewUserInfo Создает юзера по телефону
func NewUserInfo(phone string) db.UserInfo {
	user := db.UserInfo{}
	user.Phone = phone
	user.RefreshToken = ""
	user.Id = uuid.New().ID()
	user.Token = jwt.GenerateJwtById(user.Id)

	return user
}
