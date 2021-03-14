package main

import (
	"log"
	s "strings"
)

/*

 */
func NormalizeAssetSymbols(raw []AssetSymbol) []AssetSymbol {

	normalized := make([]AssetSymbol, len(raw))
	log.Println("normalizing fields of assetSymbols")
	for index := range raw {
		rawsym := raw[index]
		normalized[index] = trimToLowerCase(rawsym)
	}

	return normalized
}

// helper method
func trimToLowerCase(symbol AssetSymbol) AssetSymbol {

	symbol.AssetType = trimSingleString(symbol.AssetType)
	symbol.DisplaySymbol = trimSingleString(symbol.DisplaySymbol)
	symbol.Description = trimSingleString(symbol.Description)
	symbol.Figi = trimSingleString(symbol.Figi)
	symbol.Mic = trimSingleString(symbol.Mic)
	symbol.Currency = trimSingleString(symbol.Currency)
	symbol.Symbol = trimSingleString(symbol.Symbol)

	return symbol
}

func trimSingleString(stringToTrim string) string {

	stringToTrim = s.ToLower(stringToTrim)
	stringToTrim = s.TrimSpace(stringToTrim)
	return s.ReplaceAll(stringToTrim, " ", "-")
}
