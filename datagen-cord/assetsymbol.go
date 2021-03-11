package main

/*
	Struct for AssetSymbols.
	In Finnhub they are just called "symbols", we prefix it here to avoid confusion.
	Also, the model.AssetType field is called just "type" in finnhub, prefixed to distinguish between 'type' golang keyword
*/
type AssetSymbol struct {
	Currency      string `json:"currency"`
	Description   string `json:"description"`
	DisplaySymbol string `json:"displaySymbol"`
	Figi          string `json:"figi"`
	Mic           string `json:"mic"`
	Symbol        string `json:"symbol"`
	AssetType     string `json:"type"`
}
