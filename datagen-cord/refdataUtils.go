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

	symbol.AssetType = s.TrimSpace(s.ToLower(s.ReplaceAll(symbol.AssetType, " ", "-")))
	symbol.DisplaySymbol = s.TrimSpace(s.ToLower(s.ReplaceAll(symbol.DisplaySymbol, " ", "-")))
	symbol.Description = s.TrimSpace(s.ToLower(s.ReplaceAll(symbol.Description, " ", "-")))
	symbol.Figi = s.TrimSpace(s.ToLower(s.ReplaceAll(symbol.Figi, " ", "-")))
	symbol.Mic = s.TrimSpace(s.ToLower(s.ReplaceAll(symbol.Mic, " ", "-")))
	symbol.Currency = s.TrimSpace(s.ToLower(s.ReplaceAll(symbol.Currency, " ", "-")))
	symbol.Symbol = s.TrimSpace(s.ToLower(s.ReplaceAll(symbol.Symbol, " ", "-")))

	return symbol
}
