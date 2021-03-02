package main

import (
	"encoding/json"
	"errors"
	"github.com/Finnhub-Stock-API/finnhub-go"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var once sync.Once

func getFinnhubAuth() context.Context {
	apikey := viper.Get("API_KEY")
	auth := context.WithValue(context.Background(), finnhub.ContextAPIKey, finnhub.APIKey{
		Key: apikey.(string),
	})
	return auth
}

/*
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
*/

// setup viper
func init() {

}

func GetFinnhubSymbolsAsArr(exchange string) []AssetSymbol {
	apikey := viper.Get("FINNHUB_API_KEY") //.(string)
	resp, err := http.Get("https://finnhub.io/api/v1/stock/symbol?exchange=" + exchange + "&token=" + apikey.(string))
	if err != nil {
		log.Println(err)
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

	return result
}

func GetSymbolsAsShorthandMap(exchange string) map[string]AssetSymbol {
	resultMap, err := ConvertSymbolArr2Map(GetFinnhubSymbolsAsArr(exchange))

	if err != nil {
		panic(err)
	}

	return resultMap
}

/*
	This function converts an array of model.Asset Symbols into a map with
	the symbol shorthand (f.e. 'AAPL' for Apple Inc.) as key and the full symbol as val.
*/
func ConvertSymbolArr2Map(symbols []AssetSymbol) (symbolMap map[string]AssetSymbol, err error) {
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
