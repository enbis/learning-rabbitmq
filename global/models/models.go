package models

type Configurations struct {
	ConnString   string `yaml:"ConnString"`
	QueueName    string `yaml:"QueueName"`
	ExchangeName string `yaml:"ExchangeName"`
}
