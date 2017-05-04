package kafkax
//
//import (
//	"github.com/Shopify/sarama"
//	"crypto/tls"
//	"crypto/x509"
//	"flag"
//	"fmt"
//	"io/ioutil"
//	"log"
//	"os"
//	"strings"
//)
//
//var (
//	addr = flag.String("addr", ":9090", "The address to bind to")
//	brokers = flag.String("brokers", "localhost:9092", "The Kafka brokers to connect to, as a comma separated list")
//	verbose = flag.Bool("verbose", false, "Turn on Sarama logging")
//	certFile = flag.String("certificate", "", "The optional certificate file for client authentication")
//	keyFile = flag.String("key", "", "The optional key file for client authentication")
//	caFile = flag.String("ca", "", "The optional certificate authority file for TLS client authentication")
//	verifySsl = flag.Bool("verify", false, "Optional verify ssl certificates chain")
//)
//
//func main() {
//	flag.Parse()
//
//	if *verbose {
//		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
//	}
//
//	if *brokers == "" {
//		flag.PrintDefaults()
//		os.Exit(1)
//	}
//
//	brokerList := strings.Split(*brokers, ",")
//	log.Printf("Kafka brokers: %s", strings.Join(brokerList, ", "))
//
//	// For the data collector, we are looking for strong consistency semantics.
//	// Because we don't change the flush settings, sarama will try to produce messages
//	// as fast as possible to keep latency low.
//	config := sarama.NewConfig()
//	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
//	config.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
//	tlsConfig := createTlsConfiguration()
//	if tlsConfig != nil {
//		config.Net.TLS.Config = tlsConfig
//		config.Net.TLS.Enable = true
//	}
//
//	// On the broker side, you may want to change the following settings to get
//	// stronger consistency guarantees:
//	// - For your broker, set `unclean.leader.election.enable` to false
//	// - For the topic, you could increase `min.insync.replicas`.
//
//	producer, err := sarama.NewSyncProducer(brokerList, config)
//	if err != nil {
//		log.Fatalln("Failed to start Sarama producer:", err)
//	}
//
//	partition, offset, err := producer.SendMessage(&sarama.ProducerMessage{
//		Topic: "test",
//		Value: sarama.StringEncoder("hello k."),
//	})
//
//	if err != nil {
//		//
//
//	} else {
//		fmt.Println("partition:", partition)
//		fmt.Println("offect:", offset)
//	}
//
//}
//
//func createTlsConfiguration() (t *tls.Config) {
//	if *certFile != "" && *keyFile != "" && *caFile != "" {
//		cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		caCert, err := ioutil.ReadFile(*caFile)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		caCertPool := x509.NewCertPool()
//		caCertPool.AppendCertsFromPEM(caCert)
//
//		t = &tls.Config{
//			Certificates:       []tls.Certificate{cert},
//			RootCAs:            caCertPool,
//			InsecureSkipVerify: *verifySsl,
//		}
//	}
//	// will be nil by default if nothing is provided
//	return t
//}
//
//
//
