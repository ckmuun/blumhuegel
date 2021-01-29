package pulsarConn

import (
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"sync"
)

var pulsarClientSingletonInstance pulsar.Client
var once sync.Once
var pulsarClientCreationErr error

/*
	This method shall create a singleton.
	In the examples, these kinds of methods ususally return a pointer *
	TODO find out whether a) this also works by-value (e.g. just return the value from the NewClient func) or b) how we get the ref
*/
func GetPulsarClientInstance() pulsar.Client {
	once.Do(func() {
		pulsarClientSingletonInstance, pulsarClientCreationErr = pulsar.NewClient(
			pulsar.ClientOptions{
				URL: "pulsar://95.121.107.34.bc.googleusercontent.com:6650",
			})

		if pulsarClientCreationErr != nil {
			panic("could not create pulsar client")
		}
	})
	return pulsarClientSingletonInstance

}

func main() {

	fmt.Println("initializing finnhub datagen fetcher")

	pulsarClient, err := pulsar.NewClient(
		pulsar.ClientOptions{
			URL: "pulsar://95.121.107.34.bc.googleusercontent.com:6650",
		})

	if err != nil {
		panic("error during pulsar client init")
	}
	defer pulsarClient.Close()
}
