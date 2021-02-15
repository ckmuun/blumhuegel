package main

import (
	"github.com/stretchr/testify/assert"
	"log"
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
	var assetSymbolsNormalized = make([]AssetSymbol, 3)
	symbolNorm := AssetSymbol{
		Currency:      "usd",
		Description:   "ab-international-group-corp",
		DisplaySymbol: "abqq",
		Figi:          "bbg007b0z9j3",
		Mic:           "otcm",
		Symbol:        "abqq",
		AssetType:     "common-stock",
	}
	symbol1Norm := AssetSymbol{
		Currency:      "usd",
		Description:   "plexus-holdings-plc",
		DisplaySymbol: "plxxf",
		Figi:          "bbg003pmr1g8",
		Mic:           "ootc",
		Symbol:        "plxxf",
		AssetType:     "common-stock",
	}
	symbol2Norm := AssetSymbol{
		Currency:      "usd",
		Description:   "cansortium-inc",
		DisplaySymbol: "cntmf",
		Figi:          "bbg00nwmb6y2",
		Mic:           "otcm",
		Symbol:        "cntmf",
		AssetType:     "common-stock",
	}

	assetSymbolsNormalized[0] = symbolNorm
	assetSymbolsNormalized[1] = symbol1Norm
	assetSymbolsNormalized[2] = symbol2Norm

	normalized := NormalizeAssetSymbols(assetSymbols1)

	log.Println("conducting asserts")
	assert.Exactly(t, normalized[0], assetSymbolsNormalized[0])
	assert.Exactly(t, normalized[1], assetSymbolsNormalized[1])
	assert.Exactly(t, normalized[2], assetSymbolsNormalized[2])
}

func Test_trimToLowerCase(t *testing.T) {

	someString := " THIS is A test 227 from Testistan"
	expected := "this-is-a-test-227-from-testistan"

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
