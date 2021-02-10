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

	symbolmap := make(map[string]assetSymbol)

	symbolmap["ABQQ"] = symbol
	symbolmap["PLXXF"] = symbol1
	symbolmap["CNTMF"] = symbol2

	arr := getAssetSymbolShorthands(symbolmap)

	assert.Exactly(t, arr[0], "ABQQ")
	assert.Exactly(t, arr[1], "CNTMF")
	assert.Exactly(t, arr[2], "PLXXF")

}
