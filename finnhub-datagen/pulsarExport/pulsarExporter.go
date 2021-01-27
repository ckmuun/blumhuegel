
package pulsarExport

import (
	"context"
	"fmt"
	finnhub "github.com/Finnhub-Stock-API/finnhub-go"
	pulsar "github.com/apache/pulsar-client-go/pulsar"
)

type finnhubRecord struct {
	timestamp int64
	content string
}

func setupClient() (pulsar.client, error)   {
	fmt.Print("dubap")

}


func pubToPulsar(record finnhubRecord) string  {
	return ""
}
