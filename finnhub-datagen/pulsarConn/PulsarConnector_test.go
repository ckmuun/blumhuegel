package pulsarConn

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/rs/zerolog/log"
	"testing"
)

var topic = "topico"

func init() {
	GetPulsarClientInstance()
}

func TestGetConsumer(t *testing.T) {

	producer := GetProducer(topic)
	defer producer.Close()

	log.Print("senidng 'hello' message")
	producer.Send(
		context.Background(),
		&pulsar.ProducerMessage{
			Payload: []byte("hello"),
		},
	)

	consumer := GetConsumer(topic, "my-subscription")
	defer consumer.Close()

	msg, err := consumer.Receive(context.Background())
	if err != nil {
		log.Fatal()
	}

	fmt.Printf("Received message msgId: %#v -- content: '%s'\n",
		msg.ID(), string(msg.Payload()))
}

func TestGetProducer(t *testing.T) {

}

func TestGetPulsarClientInstance(t *testing.T) {

}
