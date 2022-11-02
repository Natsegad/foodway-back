package domain

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"regexp"
)

type Config struct {
	Port string
	IP   string

	DbName     string
	DbPassword string
	DbUser     string
	DbPort     string
}

type UserInfoPhone struct {
	Phone string `json:"phone"`
}

func (ths *UserInfoPhone) Validation() error {
	return validation.ValidateStruct(
		ths,
		validation.Field(&ths.Phone, validation.Required, validation.Length(11, 11), is.Digit),
	)
}

type UserInfo struct {
	ID       uint32 `json:"user_id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

func (ths *UserInfo) Validate() error {
	return validation.ValidateStruct(
		ths,
		validation.Field(&ths.Name, is.ASCII),
		validation.Field(&ths.Password, validation.Match(regexp.MustCompile(`[a-zA-Z0-9]]`))),
	)
}

type Response struct {
	IsPhone bool        `json:"is_phone"`
	Data    interface{} `json:"data"`
}

type JsonUserInfo struct {
	Jwt     string `json:"jwt"`
	Refresh string `json:"refresh"`
	ID      uint32 `json:"id"`
}

func (ths *Response) Parse() error {
	if ths.IsPhone {
		phoneStruct := ths.Data.(UserInfoPhone)

		fmt.Println(phoneStruct)

		return nil
	} else {
		Struct := ths.Data.(UserInfo)

		fmt.Println(Struct)

		return nil
	}
}
