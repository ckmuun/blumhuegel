package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

var once sync.Once

/*
	Struct for model.Asset Symbols.
	In Finnhub they are just called "symbols", we prefix it here to avoid confusion.
	Also, the model.AssetType field is called just "type" in finnhub, prefixed to distinguish between type golang keyword
*/

func getApiKey() string {
	log.Println("loading api key")
	file, err := os.Open("./api-key.txt") // The "./" prefix is required

	if err != nil {
		log.Fatal()
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal()
		}
	}()

	scanner := bufio.NewScanner(file)

	apikey := scanner.Text()
	return apikey
}

func GetSymbols(exchange string) map[string]AssetSymbol {
	resp, err := http.Get("https://finnhub.io/api/v1/stock/symbol?exchange=" + exchange + "&token=" + getApiKey())
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	var result []AssetSymbol

	buf, _ := ioutil.ReadAll(resp.Body)

	unmarshalErr := json.Unmarshal(buf, &result)
	if unmarshalErr != nil {
		panic(unmarshalErr)
	}
	// clear ~5mb byte buffer
	buf = nil
	resultMap, err := convertSymbolArr2Map(result)

	if err != nil {
		panic(err)
	}

	fmt.Println(resultMap["AAPL"])

	log.Print("fetching stock symbols from finnhub")

	return resultMap
}

/*
	This function converts an array of model.Asset Symbols into a map with
	the symbol shorthand (f.e. 'AAPL' for Apple Inc.) as key and the full symbol as val.
*/
func convertSymbolArr2Map(symbols []AssetSymbol) (symbolMap map[string]AssetSymbol, err error) {
	symbolMap = make(map[string]AssetSymbol)
	err = nil

	defer func() {
		symbols = nil
	}()

	if len(symbols) == 0 {
		err = errors.New("error -- can not transform 0 length arr")
		return symbolMap, err
	}

	for _, symbol := range symbols {
		symbolMap[symbol.Symbol] = symbol
	}
	return symbolMap, err
}

func convertSymbolArr2MapWithGivenKey(fieldToUseAsKey string, symbols []AssetSymbol) (symbolMap map[string][]AssetSymbol, err error) {
	symbolMap = make(map[string][]AssetSymbol)
	err = nil

	defer func() {
		symbols = nil
	}()

	if len(symbols) == 0 {
		err = errors.New("error -- can not transform 0 length arr")
		return symbolMap, err
	}

	return symbolMap, err
}
