package pulsarConn

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/rs/zerolog/log"
	"sync"
)

var pulsarClientSingleton *pulsar.Client
var once sync.Once

/*
	This method shall create a singleton.
*/
func InitPulsarClientInstance() *pulsar.Client {
	log.Print("init pulsar client")
	once.Do(func() {
		pulsarClient, err := pulsar.NewClient(
			pulsar.ClientOptions{
				URL: "pulsar://95.121.107.34.bc.googleusercontent.com:6650",
			})

		if err != nil {
			panic("could not create pulsar client")
		}

		defer pulsarClient.Close()
		pulsarClientSingleton = &pulsarClient
	})
	return pulsarClientSingleton
}

/*
	Creat a pulsar producer
	@return pulsar.Producer
	@topic the topic to produce messages to.
	Note that the syntax is [persistent || non-persistent]://<tenant>/<namespace>/<topic>
	just <topic> will default to persistent://default/pulsar/<topic>
*/
func GetProducer(topic string) pulsar.Producer {
	log.Print("deref pulsar client to create producer")
	client := *pulsarClientSingleton
	log.Print("creating producer")
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})

	if err != nil {
		panic("cant creat prod")
	}
	//defer producer.Close()
	return producer
}

/*
	Create a pulsar consumer for a given topic and a given subscriptionName
*/
func GetConsumer(topic string, subscriptionName string) pulsar.Consumer {
	log.Print("deref pulsar client to create consumer e.g. subscription")
	client := *pulsarClientSingleton

	log.Print("creat consumer subscriber")
	consumer, err := client.Subscribe(
		pulsar.ConsumerOptions{
			Topic:            topic,
			SubscriptionName: subscriptionName,
			Type:             pulsar.Shared,
		})
	if err != nil {
		panic("could not create consumer")
	}

	//	defer consumer.Close()

	return consumer
}
