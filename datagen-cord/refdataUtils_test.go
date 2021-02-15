package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var assetSymbols1 = make([]AssetSymbol, 3)

func init() {
	symbol := AssetSymbol{Currency: "USD",
		Description:   "AB INTERNATIONAL GROUP CORP",
		DisplaySymbol: "ABQQ",
		Figi:          "BBG007B0Z9J3",
		Mic:           "OTCM",
		Symbol:        "ABQQ",
		AssetType:     "Common Stock",
	}
	symbol1 := AssetSymbol{
		Currency:      "USD",
		Description:   "PLEXUS HOLDINGS PLC",
		DisplaySymbol: "PLXXF",
		Figi:          "BBG003PMR1G8",
		Mic:           "OOTC",
		Symbol:        "PLXXF",
		AssetType:     "Common Stock",
	}
	symbol2 := AssetSymbol{
		Currency:      "USD",
		Description:   "CANSORTIUM INC",
		DisplaySymbol: "CNTMF",
		Figi:          "BBG00NWMB6Y2",
		Mic:           "OTCM",
		Symbol:        "CNTMF",
		AssetType:     "Common Stock",
	}

	assetSymbols1[0] = symbol
	assetSymbols1[1] = symbol1
	assetSymbols1[2] = symbol2

}

func TestNormalizeAssetSymbols(t *testing.T) {

}

func Test_trimToLowerCase(t *testing.T) {

	someString := " THIS is A test from Testistan"
	expected := "thisisatestfromtestistan"

	symbol := AssetSymbol{
		Currency:      someString,
		Description:   someString,
		DisplaySymbol: someString,
		Figi:          someString,
		Mic:           someString,
		Symbol:        someString,
		AssetType:     someString,
	}

	symbol = trimToLowerCase(symbol)

	assert.Exactly(t, expected, symbol.Currency)
	assert.Exactly(t, expected, symbol.Description)
	assert.Exactly(t, expected, symbol.DisplaySymbol)
	assert.Exactly(t, expected, symbol.Figi)
	assert.Exactly(t, expected, symbol.Mic)
	assert.Exactly(t, expected, symbol.AssetType)
	assert.Exactly(t, expected, symbol.Symbol)
}
