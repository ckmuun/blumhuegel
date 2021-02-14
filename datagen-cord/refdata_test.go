package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var AssetSymbols = make([]AssetSymbol, 3)

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

	AssetSymbols[0] = symbol
	AssetSymbols[1] = symbol1
	AssetSymbols[2] = symbol2

}

func TestGetSymbols(t *testing.T) {

}

func Test_convertSymbolArr2Map(t *testing.T) {

	symbolMap, err := ConvertSymbolArr2Map(AssetSymbols)
	assert.NoError(t, err)

	assert.Exactly(t, symbolMap["CNTMF"], AssetSymbols[0])
	assert.Exactly(t, symbolMap["PLXXF"], AssetSymbols[1])
	assert.Exactly(t, symbolMap["ABQQ"], AssetSymbols[2])
}
