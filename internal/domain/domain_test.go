package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserInfo_Validate(t *testing.T) {
	user := UserInfo{
		1,
		"123",
		"ASD",
		"89634239857",
	}

	assert.NoError(t, user.Validate())
}

func TestUserInfoPhone_Validation(t *testing.T) {
	user := UserInfoPhone{"89634239857"}

	assert.NoError(t, user.Validation())
}

//func TestResponse_Parse(t *testing.T) {
//
//	tests := []struct {
//		name     string
//		Function func()
//	}{}
//
//	for _, v := range tests {
//		t.Run(v.name, func(t *testing.T) {
//			data := UserInfoPhone{"123123"}
//			user := Response{
//				true,
//				data,
//			}
//
//			assert.NoError(t, user.Parse())
//		})
//	}
//}
