package service

import (
	"errors"
	"fmt"
	"foodway/internal/domain"
	"foodway/pkg/db"
	jwtf "foodway/pkg/jwt"
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
func addUserToDB(user domain.UserInfoPhone) (domain.JsonUserInfo, error) {
	nwUserInfo := NewUserInfo(user.Phone)
	jsUserInfoResp := NewJsonUserInfo(nwUserInfo.Token, nwUserInfo.RefreshToken, nwUserInfo.Id)

	return jsUserInfoResp, db.AddUser(nwUserInfo)
}

// checkHaveUser проверка того что указанный номер есть в базе
func checkHaveUser(phone string) (db.UserInfo, error) {
	ok := db.GetUserByPhone(phone)
	if ok.Id != 0 {
		return ok, errors.New(fmt.Sprintf("%s user have", phone))
	}

	return ok, nil
}

// Главная функция регистрации
func Registration(user domain.UserInfoPhone) (*domain.JsonUserInfo, error) {
	log := logger.GetLogger()

	log.Infof("Start registration user %s ", user.Phone)

	// Первый этап валидация данных
	err := validation(user)
	if err != nil {
		log.Errorf("Validation user: %s error \n", user.Phone)
		return nil, err
	}

	_, err = checkHaveUser(user.Phone)
	if err != nil {
		return nil, err
	}

	// Проходит валидация идет добавление в базу данных
	resp, err := addUserToDB(user)
	if err != nil {
		log.Errorf("Add user: %s to data base error \n", user.Phone)
		return nil, err
	}

	return &resp, nil
}

// Создает после обнволяет jwt
func Jwt(user domain.JsonUserInfo) (*domain.JsonUserInfo, error) {
	log := logger.GetLogger()

	isExpr, err := jwtf.TokenExpired(user.Jwt)
	if err != nil {
		return nil, err
	}
	if isExpr {
		isExpr, err := jwtf.TokenExpired(user.Refresh)
		if err != nil {
			return nil, err
		}
		if !isExpr {
			jwt := jwtf.GenerateJwtById(user.ID)
			userUpd, err := db.GetUser(user.ID)
			if err != nil {
				return nil, err
			}
			log.Infof("jwt for user: %d updated ", user.ID)
			userUpd.Token = jwt

			err = db.UpdateUser(userUpd)

			user.Jwt = jwt
			return &user, nil
		}
	}

	return &user, nil
}

// Главная функция авторизации
func Authorization(user domain.JsonUserInfo) (*domain.JsonUserInfo, error) {
	log := logger.GetLogger()

	userInfo, err := db.GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if userInfo.Id == 0 {
		return nil, errors.New(fmt.Sprintf("User %d not found", user.ID))
	}

	nwUser, err := Jwt(user)
	if err != nil {
		return nil, err
	}
	log.Infof("Jwt - %s ", nwUser.Jwt)

	return nwUser, err
}

// NewUserInfo Создает юзера по телефону
func NewUserInfo(phone string) db.UserInfo {
	user := db.UserInfo{}
	user.Phone = phone
	user.Id = uuid.New().ID()
	user.RefreshToken = jwtf.GenerateRefeshTokenById(user.Id)
	user.Token = jwtf.GenerateJwtById(user.Id)

	return user
}

func NewJsonUserInfo(jwt, refresh string, id uint32) domain.JsonUserInfo {
	user := domain.JsonUserInfo{}
	user.Jwt = jwt
	user.Refresh = refresh
	user.ID = id

	return user
}
