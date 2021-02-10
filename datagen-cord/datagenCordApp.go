package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"sort"
)

var usSymbolsFull = make(map[string]assetSymbol)

// TODO add configuration for file-based port settings
func main() {

	log.Println("fetching us symbols")
	usSymbolsFull = GetSymbols("US")

	router := setupRouter()
	router.Run(":7077")
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	router.GET("/symbols/us", func(c *gin.Context) {
		c.JSON(200, usSymbolsFull)
	})

	router.GET("/symbols/us/short", func(c *gin.Context) {
		c.JSON(200, getAssetSymbolShorthands(usSymbolsFull))
	})
	return router
}

/*
	This just takes the map keys (the stock symbol shorthands) and puts them into an array.
*/
func getAssetSymbolShorthands(fullSymbols map[string]assetSymbol) []string {
	log.Println("converting full symbols map to array containing shorthands")
	log.Println("number of symbols:")
	log.Println(len(fullSymbols))

	keys := make([]string, len(fullSymbols))

	i := 0
	for k := range fullSymbols {
		keys[i] = k
		i++
	}
	// Everything in Go is evaluated lazily, but this runs in-place.
	log.Println("sorting shorthands")
	sort.Strings(keys)
	return keys
}
