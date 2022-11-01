package service

import (
	"foodway/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegistration(t *testing.T) {
	funcs := []struct {
		name    string
		fn      func() *domain.UserInfoPhone
		IsValid bool
	}{
		{
			name: "valid",
			fn: func() *domain.UserInfoPhone {
				return &domain.UserInfoPhone{
					Phone: "89634239857",
				}
			},
			IsValid: true,
		},
		{
			name: "invalid",
			fn: func() *domain.UserInfoPhone {
				return &domain.UserInfoPhone{
					Phone: "89634239",
				}
			},
			IsValid: false,
		},
		{
			name: "invalid",
			fn: func() *domain.UserInfoPhone {
				return &domain.UserInfoPhone{
					Phone: "asdasd",
				}
			},
			IsValid: false,
		},
	}

	for _, v := range funcs {
		t.Run(v.name, func(t *testing.T) {
			if v.IsValid {
				assert.NoError(t, v.fn().Validation())
			} else {
				assert.Error(t, v.fn().Validation())
			}
		})
	}
}
