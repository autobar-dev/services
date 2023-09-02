package utils

import (
	"github.com/autobar-dev/services/user/repositories"
	"github.com/autobar-dev/services/user/types"
	walletrepository "github.com/autobar-dev/shared-libraries/go/wallet-repository"
)

func PostgresUserToUser(pu repositories.PostgresUser) *types.User {
	return &types.User{
		Id:                         pu.Id,
		Email:                      pu.Email,
		FirstName:                  pu.FirstName,
		LastName:                   pu.LastName,
		DateOfBirth:                pu.DateOfBirth,
		Locale:                     pu.Locale,
		IdentityVerificationId:     pu.IdentityVerificationId,
		IdentityVerificationSource: pu.IdentityVerificationSource,
		CreatedAt:                  pu.CreatedAt,
		UpdatedAt:                  pu.UpdatedAt,
	}
}

func RedisUserToUser(ru repositories.RedisUser) *types.User {
	return &types.User{
		Id:                         ru.Id,
		Email:                      ru.Email,
		FirstName:                  ru.FirstName,
		LastName:                   ru.LastName,
		DateOfBirth:                ru.DateOfBirth,
		Locale:                     ru.Locale,
		IdentityVerificationId:     ru.IdentityVerificationId,
		IdentityVerificationSource: ru.IdentityVerificationSource,
		CreatedAt:                  ru.CreatedAt,
		UpdatedAt:                  ru.UpdatedAt,
	}
}

func UserToUserExtended(u types.User, w walletrepository.Wallet) *types.UserExtended {
	dob := TimeToDateString(u.DateOfBirth)

	return &types.UserExtended{
		Id:                         u.Id,
		Email:                      u.Email,
		FirstName:                  u.FirstName,
		LastName:                   u.LastName,
		DateOfBirth:                dob,
		Locale:                     u.Locale,
		Wallet:                     w,
		IdentityVerificationId:     u.IdentityVerificationId,
		IdentityVerificationSource: u.IdentityVerificationSource,
		CreatedAt:                  u.CreatedAt,
		UpdatedAt:                  u.UpdatedAt,
	}
}
