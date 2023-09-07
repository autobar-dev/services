package types

import "github.com/autobar-dev/shared-libraries/go/product-repository"

type PrepareModuleData struct {
	Module  Module                    `json:"module"`
	Product productrepository.Product `json:"product"`
	Otk     string                    `json:"otk"`
}
