package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSymbols(t *testing.T) {

}

func Test_convertSymbolArr2Map(t *testing.T) {

	symbol := assetSymbol{"USD",
		"AB INTERNATIONAL GROUP CORP",
		"ABQQ",
		"BBG007B0Z9J3",
		"OTCM",
		"ABQQ",
		"Common Stock",
	}
	symbol1 := assetSymbol{
		"USD",
		"PLEXUS HOLDINGS PLC",
		"PLXXF",
		"BBG003PMR1G8",
		"OOTC",
		"PLXXF",
		"Common Stock",
	}
	symbol2 := assetSymbol{
		"USD",
		"CANSORTIUM INC",
		"CNTMF",
		"BBG00NWMB6Y2",
		"OTCM",
		"CNTMF",
		"Common Stock",
	}

	arr := make([]assetSymbol, 3)

	arr[0] = symbol
	arr[1] = symbol1
	arr[2] = symbol2

	symbolMap, err := convertSymbolArr2Map(arr)
	assert.NoError(t, err)

	assert.Exactly(t, symbolMap["CNTMF"], symbol2)
	assert.Exactly(t, symbolMap["PLXXF"], symbol1)
	assert.Exactly(t, symbolMap["ABQQ"], symbol)
}
