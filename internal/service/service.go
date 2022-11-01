package service

import (
	"foodway/internal/domain"
	"foodway/pkg/logger"
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

	// Проходит валидация идет добавление в базу данных
	err = addUserToDB(user)
	if err != nil {
		log.Errorf("Add user: %s to data base error \n", user.Phone)
		return err
	}

	return nil
}
