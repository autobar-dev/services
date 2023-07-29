package repositories

import (
	"encoding/json"
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

type MeiliProductsSearchOptions struct {
	Query           string
	HitsPerPage     int
	Page            int
	IncludeDisabled bool
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

func (mr MeiliRepository) SearchProducts(options *MeiliProductsSearchOptions) (*[]MeiliProduct, error) {
	filters := [][]string{}

	if !options.IncludeDisabled {
		filters = append(filters, []string{"enabled = true"})
	}

	result, err := mr.index.Search(options.Query, &meilisearch.SearchRequest{
		Query:       options.Query,
		HitsPerPage: int64(options.HitsPerPage),
		Page:        int64(options.Page),
		Filter:      filters,
	})
	if err != nil {
		return nil, err
	}

	hits := []MeiliProduct{}
	for _, h := range result.Hits {
		h_json_bytes, _ := json.Marshal(h)
		var mp MeiliProduct
		err = json.Unmarshal(h_json_bytes, &mp)

		if err != nil {
			fmt.Printf("failed to parse meili product = %+v\n", h)
		} else {
			hits = append(hits, mp)
		}
	}

	return &hits, nil
}
