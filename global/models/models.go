package models

type Configurations struct {
	ConnString   string `yaml:"ConnString"`
	QueueName    string `yaml:"QueueName"`
	ExchangeName string `yaml:"ExchangeName"`
}

var IntToString = map[int]string{
	0: "all",
	1: "first",
	2: "second",
}
