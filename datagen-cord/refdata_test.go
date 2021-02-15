package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var assetSymbols = make([]AssetSymbol, 3)

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

	assetSymbols[0] = symbol
	assetSymbols[1] = symbol1
	assetSymbols[2] = symbol2

}

func TestGetSymbols(t *testing.T) {

}

func Test_convertSymbolArr2Map(t *testing.T) {

	symbolMap, err := ConvertSymbolArr2Map(assetSymbols)
	assert.NoError(t, err)

	assert.Exactly(t, symbolMap["CNTMF"], assetSymbols[0])
	assert.Exactly(t, symbolMap["PLXXF"], assetSymbols[1])
	assert.Exactly(t, symbolMap["ABQQ"], assetSymbols[2])
}
