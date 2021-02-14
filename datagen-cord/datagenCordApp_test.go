package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "\"pong\"", w.Body.String())
}

func Test_main(t *testing.T) {

}

func Test_getAssetSymbolsShorthands(t *testing.T) {
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

	symbolmap := make(map[string]AssetSymbol)

	symbolmap["ABQQ"] = symbol
	symbolmap["PLXXF"] = symbol1
	symbolmap["CNTMF"] = symbol2

	arr := getAssetSymbolShorthands(symbolmap)

	assert.Exactly(t, arr[0], "ABQQ")
	assert.Exactly(t, arr[1], "CNTMF")
	assert.Exactly(t, arr[2], "PLXXF")

}
