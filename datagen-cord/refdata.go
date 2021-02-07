package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var once sync.Once

/*
	Struct for Asset Symbols.
	In Finnhub they are just called "symbols", we prefix it here to avoid confusion.
	Also, the assetType field is called just "type" in finnhub, prefixed to distinguish between type golang keyword
*/
type assetSymbol struct {
	Currency      string `json:"currency"`
	Description   string `json:"description"`
	DisplaySymbol string `json:"displaySymbol"`
	Figi          string `json:"figi"`
	Mic           string `json:"mic"`
	Symbol        string `json:"symbol"`
	AssetType     string `json:"type"`
}

func GetSymbols(exchange string) map[string]assetSymbol {
	//client := http.Client{}
	resp, err := http.Get("https://finnhub.io/api/v1/stock/symbol?exchange=" + exchange + "&token=c07ianf48v6retjaflk0")
	//resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	var resultMap map[string]assetSymbol
	var result []assetSymbol

	buf, _ := ioutil.ReadAll(resp.Body)

	//jsonbuf := string(buf)

	fmt.Println("################################################################")
	//fmt.Println(jsonbuf)
	fmt.Println("################################################################")

	unmarshalErr := json.Unmarshal(buf, &result)
	if unmarshalErr != nil {
		panic(unmarshalErr)
	}
	// clear ~5mb byte buffer
	buf = nil

	fmt.Println(result)

	// buf is non empty, around 4M Bytes TODO how do we serdes the byte buffer into a json string and/ or a runtime struct / interface

	log.Print("fetching stock symbols from finnhub")
	/*	if err != nil {
			fmt.Println(err)
		}

		//	res := json.NewDecoder(resp.Body).Decode(&result)
	*/

	//resp, err := client.Do(request)
	//var result map[string]interface{}

	log.Print("res:")
	//	log.Print(&res)

	return resultMap
}

/*
	This function converts an array of Asset Symbols into a map with
	the symbol shorthand (f.e. 'AAPL' for Apple Inc.) as key and the full symbol as val.
*/
func convertSymbolArr2Map(symbols []assetSymbol) (symbolMap map[string]assetSymbol, err error) {
	symbolMap = make(map[string]assetSymbol)
	err = nil

	defer func() {
		symbols = nil
	}()

	if len(symbols) == 0 {
		err = errors.New("error -- can not transform 0 length arr")
		return symbolMap, err
	}

	for _, symbol := range symbols {
		log.Println(symbol)
		symbolMap[symbol.Symbol] = symbol
	}
	return symbolMap, err
}
