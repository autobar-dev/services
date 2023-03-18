package types

import "github.com/autobar-dev/services/currency/types/interfaces"

type AppContext struct {
	AppLogger *interfaces.AppLogger
	Stores    *AppStores
}
