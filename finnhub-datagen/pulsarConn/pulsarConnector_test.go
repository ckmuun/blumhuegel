package pulsarConn

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/rs/zerolog/log"
	"testing"
)

var topic = "blumhuegel/findata/test"

/*
	setup method initializing the pulsar client
*/
func init() {

	InitPulsarClientInstance("pulsar://192.168.178.8:30560")
}

func TestPulsarPubSubSimple(t *testing.T) {

	producer := *GetProducer(topic)
	defer producer.Close()

	log.Print("senidng 'hello' message")
	msgId, msgerr := producer.Send(
		context.Background(),
		&pulsar.ProducerMessage{
			Payload: []byte("hello my friendly friend"),
		},
	)
	if msgerr != nil {
		panic("message send to pulsar failed")
	}

	fmt.Print("message id ", msgId)

	consumer := GetConsumer(topic, "gopher-subscription")
	defer consumer.Close()

	msg, err := consumer.Receive(context.Background())
	if err != nil {
		log.Fatal()
	}

	fmt.Printf("Received message msgId: %#v -- content: '%s'\n",
		msg.ID(), string(msg.Payload()))
}

func TestPulsarPubSubSchema(t *testing.T) {
	log.Print("not implemented yet")
}
