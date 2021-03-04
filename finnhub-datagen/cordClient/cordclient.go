package cordClient

import (
	"encoding/json"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var once sync.Once
var httpclient *http.Client

type assetSymbol struct {
	Currency      string `json:"currency"`
	Description   string `json:"description"`
	DisplaySymbol string `json:"displaySymbol"`
	Figi          string `json:"figi"`
	Mic           string `json:"mic"`
	Symbol        string `json:"symbol"`
	AssetType     string `json:"type"`
}

func init() {
	once.Do(func() {
		httpclient = &http.Client{
			Timeout: time.Second * 10,
		}
	})
}

func getHttpClient() *http.Client {
	return httpclient
}

func GetDemoSymbols() (symbols []string) {
	client := getHttpClient()

	resp, err := client.Get(viper.GetString("CORDSVC_URL") + "/demo")

	if err != nil {
		panic("error during cord svc query")
	}
	defer resp.Body.Close()

	buf, _ := ioutil.ReadAll(resp.Body)

	unmarshalErr := json.Unmarshal(buf, &symbols)
	if unmarshalErr != nil {
		panic(unmarshalErr)
	}
	buf = nil

	return symbols
}

func GetSymbolShorthandsToQuery() (symbols []string) {
	client := getHttpClient()

	resp, err := client.Get(viper.GetString("CORDSVC_URL") + "/selection")

	if err != nil {
		panic("error during cord svc query")
	}
	defer resp.Body.Close()

	buf, _ := ioutil.ReadAll(resp.Body)

	unmarshalErr := json.Unmarshal(buf, &symbols)
	if unmarshalErr != nil {
		panic(unmarshalErr)
	}
	// clear ~5mb byte buffer
	buf = nil

	return symbols
}
