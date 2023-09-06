package utils

import (
	"github.com/autobar-dev/services/module/repositories"
	"github.com/autobar-dev/services/module/types"
)

func RedisActivationSessionToActivationSession(ras repositories.RedisActivationSession) *types.ActivationSession {
	return &types.ActivationSession{
		Id:                ras.Id,
		UserId:            ras.UserId,
		SerialNumber:      ras.SerialNumber,
		ProductId:         ras.ProductId,
		Price:             ras.Price,
		AmountMillilitres: ras.AmountMillilitres,
		CreatedAt:         ras.CreatedAt,
		UpdatedAt:         ras.UpdatedAt,
	}
}

func ActivationSessionToRedisActivationSession(as types.ActivationSession) *repositories.RedisActivationSession {
	return &repositories.RedisActivationSession{
		Id:                as.Id,
		UserId:            as.UserId,
		SerialNumber:      as.SerialNumber,
		ProductId:         as.ProductId,
		Price:             as.Price,
		AmountMillilitres: as.AmountMillilitres,
		CreatedAt:         as.CreatedAt,
		UpdatedAt:         as.UpdatedAt,
	}
}
