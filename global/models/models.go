package models

type Configurations struct {
	ConnString   string `yaml:"ConnString"`
	QueueName    string `yaml:"QueueName"`
	ExchangeName string `yaml:"ExchangeName"`
}

var IntToString = map[int]string{
	0: "all_rooms",
	1: "first_room",
	2: "second_room",
}
