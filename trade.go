package bittrex

import "github.com/shopspring/decimal"

// Trade Used in getmarkethistory
type Trade struct {
	OrderUUID int64           `json:"Id"`
	Timestamp jTime           `json:"TimeStamp"`
	Quantity  decimal.Decimal `json:"Quantity"`
	Price     decimal.Decimal `json:"Price"`
	Total     decimal.Decimal `json:"Total"`
	FillType  string          `json:"FillType"`
	OrderType string          `json:"OrderType"`
}