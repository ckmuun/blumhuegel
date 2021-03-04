package pulsarConn

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/rs/zerolog/log"
	"sync"
)

var pulsarClientSingleton *pulsar.Client
var pulsarProducers map[string]*pulsar.Producer

var onceCreatePulsarClient sync.Once

func init() {
	pulsarProducers = make(map[string]*pulsar.Producer)

}

/*
	This method shall create a singleton.
*/
func InitPulsarClientInstance(pulsarEndpoint string) *pulsar.Client {
	log.Print("init pulsar client")
	onceCreatePulsarClient.Do(func() {
		pulsarClient, err := pulsar.NewClient(
			pulsar.ClientOptions{
				URL: pulsarEndpoint,
				//URL: "pulsar://95.121.107.34.bc.googleusercontent.com:6650",
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
	just <topic> will default to "persistent://default/pulsar/<topic>"
*/
func GetProducer(persTenNsTopic string) *pulsar.Producer {
	log.Print("deref pulsar client to create producer")
	client := *pulsarClientSingleton

	producer := pulsarProducers[persTenNsTopic]

	if producer != nil {
		return producer
	}

	log.Print("creating producer")
	createdProducer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: persTenNsTopic,
	})

	if err != nil {
		panic("cant creat prod")
	}
	//defer producer.Close()

	pulsarProducers[persTenNsTopic] = &createdProducer

	return &createdProducer
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
