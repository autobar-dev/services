package types

import (
	"time"

	filerepository "github.com/autobar-dev/shared-libraries/go/file-repository"
)

type ProductBadgeType string

const (
	ProductBadgeTypePrimary   ProductBadgeType = "primary"
	ProductBadgeTypeSecondary ProductBadgeType = "secondary"
)

type ProductBadge struct {
	Type  ProductBadgeType `json:"type"`
	Label string           `json:"label"`
	Value *string          `json:"value"`
}

type Product struct {
	Id           string              `json:"id"`
	Names        map[string]string   `json:"names"`
	Descriptions map[string]string   `json:"descriptions"`
	Cover        filerepository.File `json:"cover"`
	Enabled      bool                `json:"enabled"`
	Badges       []ProductBadge      `json:"badges"`
	CreatedAt    time.Time           `json:"created_at"`
	UpdatedAt    time.Time           `json:"updated_at"`
}
