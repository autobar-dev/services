package interfaces

type RemoteExchangeRateStoreRow struct {
	BaseCode        string
	DestinationCode string
	ConversionRate  float64
}

type RemoteExchangeRateStore interface {
	Get(string, string) (*RemoteExchangeRateStoreRow, error)
}
