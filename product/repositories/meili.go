package repositories

import (
	"fmt"
	"time"

	"github.com/meilisearch/meilisearch-go"
)

type MeiliProduct struct {
	Id           string            `json:"id"`
	Names        map[string]string `json:"names"`
	Descriptions map[string]string `json:"descriptions"`
	Cover        string            `json:"cover"`
	Enabled      bool              `json:"enabled"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
}

type MeiliRepository struct {
	client *meilisearch.Client
	index  *meilisearch.Index
}

const MeiliProductsIndexName string = "products"

func NewMeiliRepository(client *meilisearch.Client) *MeiliRepository {
	index, _ := client.GetIndex(MeiliProductsIndexName)

	return &MeiliRepository{
		client: client,
		index:  index,
	}
}

func (mr MeiliRepository) AddProduct(id string, names map[string]string, descriptions map[string]string, cover string, enabled bool, created_at time.Time, updated_at time.Time) error {
	mp := &MeiliProduct{
		Id:           id,
		Names:        names,
		Descriptions: descriptions,
		Cover:        cover,
		Enabled:      enabled,
		CreatedAt:    created_at,
		UpdatedAt:    updated_at,
	}

	_, err := mr.index.AddDocuments(mp)
	if err != nil {
		fmt.Printf("failed to add product to meili: %v\n", err)
	}

	return err
}
