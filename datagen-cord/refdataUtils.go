package main

import (
	"log"
	s "strings"
)

/*

 */
func NormalizeAssetSymbols(raw []AssetSymbol) (normalized []AssetSymbol) {

	log.Println("normalizing fields of assetSymbols")
	for index := range raw {
		rawsym := raw[index]
		normalized[index] = trimToLowerCase(rawsym)
	}

	return normalized
}

// helper method
func trimToLowerCase(symbol AssetSymbol) AssetSymbol {

	symbol.AssetType = s.ToLower(s.ReplaceAll(symbol.AssetType, " ", ""))
	symbol.DisplaySymbol = s.ToLower(s.ReplaceAll(symbol.DisplaySymbol, " ", ""))
	symbol.Description = s.ToLower(s.ReplaceAll(symbol.Description, " ", ""))
	symbol.Figi = s.ToLower(s.ReplaceAll(symbol.Figi, " ", ""))
	symbol.Mic = s.ToLower(s.ReplaceAll(symbol.Mic, " ", ""))
	symbol.Currency = s.ToLower(s.ReplaceAll(symbol.Currency, " ", ""))
	symbol.Symbol = s.ToLower(s.ReplaceAll(symbol.Symbol, " ", ""))

	return symbol
}
