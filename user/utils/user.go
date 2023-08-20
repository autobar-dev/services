package utils

import (

	"github.com/autobar-dev/services/user/repositories"
	"github.com/autobar-dev/services/user/types"
)

func PostgresUserToUser(pu repositories.PostgresUser) *types.User {
	return &types.User {
		Id:           pu.Id,
		Email: pu.Email,
		PhoneNumberCountryCode: pu.PhoneNumberCountryCode,
		PhoneNumber: pu.PhoneNumber,
		FirstName: pu.FirstName,
		LastName: pu.LastName,
		Locale: pu.Locale,
		Verified: pu.Verified,
		VerifiedAt: pu.VerifiedAt,
		CreatedAt: pu.CreatedAt,
		UpdatedAt: pu.UpdatedAt,
	}
}

func RedisUserToUser(ru repositories.RedisUser) *types.User {
	return &types.User {
		Id:           ru.Id,
		Email: ru.Email,
		PhoneNumberCountryCode: ru.PhoneNumberCountryCode,
		PhoneNumber: ru.PhoneNumber,
		FirstName: ru.FirstName,
		LastName: ru.LastName,
		Locale: ru.Locale,
		Verified: ru.Verified,
		VerifiedAt: ru.VerifiedAt,
		CreatedAt: ru.CreatedAt,
		UpdatedAt: ru.UpdatedAt,
	}
}


