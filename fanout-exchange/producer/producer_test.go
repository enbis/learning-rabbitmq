package producer

import (
	"fmt"
	"testing"
	"time"

	"github.com/enbis/learning-rabbitmq/global/connection"
	"github.com/enbis/learning-rabbitmq/global/models"
	"github.com/enbis/learning-rabbitmq/global/utils"
	"github.com/spf13/viper"
)

var configuration *models.Configurations

func TestLaunchPublisher(t *testing.T) {
	initConf()

	broker, err := connection.Connect(configuration.ConnString)
	if err != nil {
		panic(err.Error())
	}

	err = broker.SetExchange(configuration.ExchangeName, "fanout", true, false, false, false, nil)
	if err != nil {
		panic(err.Error())
	}

	for i := 0; i < 10; i++ {
		time.Sleep(5000 * time.Millisecond)
		m := fmt.Sprintf("Message %d %s", i, utils.RandomStr())
		fmt.Printf("%d: %s \n", i, m)
		Publish(m, *broker, configuration.ExchangeName)
	}

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
