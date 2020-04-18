GOCMD=go
GOTEST=$(GOCMD) test

test: test/consumer \ 
	test/consumer \ 
	test/producer

test/producer:
	cd ~/go/src/github.com/enbis/learning-rabbitmq/fanout-exchange/producer \
	&& $(GOTEST) -run TestLaunchPublisher

test/consumer:
	cd ~/go/src/github.com/enbis/learning-rabbitmq/fanout-exchange/consumer \
	&& $(GOTEST) -run TestLaunchConsumer