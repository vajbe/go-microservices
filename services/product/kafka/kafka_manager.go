package kafka

type KafkaManager struct {
	Consumer *KafkaConsumer
	Producer *KafkaProducer
}

var kManager KafkaManager

func SetKafkaManager(Consumer *KafkaConsumer, Producer *KafkaProducer) {
	kManager = KafkaManager{
		Consumer: Consumer,
		Producer: Producer,
	}
}

func GetKafkaManger() *KafkaManager {
	return &kManager
}
