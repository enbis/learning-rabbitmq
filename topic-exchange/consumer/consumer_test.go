package consumer

import (
	"flag"
	"fmt"
	"strings"
	"testing"

	"github.com/enbis/learning-rabbitmq/global/connection"
	"github.com/enbis/learning-rabbitmq/global/models"
	"github.com/spf13/viper"
)

var configuration *models.Configurations
var qb = flag.String("qbinding", "", "binding selection")
var qname = flag.String("qname", "", "queue name")

func TestLaunchConsumer(t *testing.T) {
	initConf()

	fmt.Println("qbinding ", *qb)

	if *qb == "" || *qname == "" {
		panic("flags are not properly setted")
	}

	broker, err := connection.Connect(configuration.ConnString)
	if err != nil {
		panic(err.Error())
	}

	err = broker.SetExchange(configuration.ExchangeName, "topic", true, false, false, false, nil)
	if err != nil {
		panic(err.Error())
	}

	err = broker.SetQueue(*qname, false, false, true, false, nil)
	if err != nil {
		panic(err.Error())
	}

	err = broker.SetBinding(configuration.ExchangeName, *qb, *qname, false, nil)
	if err != nil {
		panic(err.Error())
	}

	msgs, err := ConsumeMsg(*broker, "")
	if err != nil {
		panic(err.Error())
	}
	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			s := strings.Split(string(msg.RoutingKey), ".")
			fmt.Println("Where: " + s[1])
			fmt.Println("What: " + s[2])
			fmt.Println(string(msg.Body))
			fmt.Println("----------")
		}

	}()

	fmt.Println("Waiting messages")
	fmt.Println("")
	<-forever
}

func initConf() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../global/config")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	viper.SetDefault("connstring", "")
	viper.SetDefault("queuename", "")

	configuration = &models.Configurations{}
	err = viper.Unmarshal(configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

}
