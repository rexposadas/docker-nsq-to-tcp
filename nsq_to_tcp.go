package main

import (
	"flag"
	"fmt"
	"github.com/bitly/go-nsq"
	"github.com/bitly/nsq/util"
	"log"
	"net"
	"os"
	"os/signal"
)

var (
	topic            = flag.String("topic", "", "nsq topic")
	channel          = flag.String("channel", "", "nsq channel")
	outputTCPAddrs   = flag.String("output-tcp-address", "", "TCP address to send the message to")
	lookupdHTTPAddrs = util.StringArray{}
)

func init() {
	flag.Var(&lookupdHTTPAddrs, "lookupd-http-address", "lookupd HTTP address (may be given multiple times)")
}

type Handler struct {
}

func (h Handler) HandleMessage(message *nsq.Message) error {

	conn, err := net.Dial("tcp", *outputTCPAddrs)
	if err != nil {
		fmt.Printf("ERROR: could not connect to '%s' to send message '%s' \n\n",
			*outputTCPAddrs, message.Body)
		return nil
	}
	defer conn.Close()

	fmt.Printf("Sending: %s to address %+v \n\n", message.Body, *outputTCPAddrs)
	conn.Write(message.Body)

	return nil
}

func main() {

	flag.Parse()

	if *channel == "" {
		*channel = fmt.Sprintf("1")
	}

	if *topic == "" {
		log.Fatalf("--topic is required")
	}

	if len(lookupdHTTPAddrs) == 0 {
		log.Fatalf("-lookupd-http-address required")
	}

	if *outputTCPAddrs == "" {
		log.Fatalf("-output-tcp-address required")
	}

	fmt.Printf("using %s %s %s %+v \n", *channel, *topic, *outputTCPAddrs, lookupdHTTPAddrs)

	cfg := nsq.NewConfig()

	r, err := nsq.NewConsumer(*topic, *channel, cfg)
	if err != nil {
		log.Fatalf(err.Error())
	}
	h := &Handler{}

	r.SetHandler(h)

	for _, a := range lookupdHTTPAddrs {
		err = r.ConnectToNSQLookupd(a)
		if err != nil {
			log.Fatalf(err.Error())
		}
	}

	if err != nil {
		panic(err)
	}

	// Handle SIGINT and SIGTERM.
	finish := make(chan os.Signal)
	signal.Notify(finish, os.Interrupt, os.Kill)
	<-finish
}
