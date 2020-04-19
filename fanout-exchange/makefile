GOCMD=go
GOTEST=$(GOCMD) test

test/producer:
	cd ~/go/src/github.com/enbis/learning-rabbitmq/fanout-exchange/producer \
	&& $(GOTEST) -run TestLaunchPublisher

test/consumer/room0:
	cd ~/go/src/github.com/enbis/learning-rabbitmq/fanout-exchange/consumer \
	&& $(GOTEST) -run TestLaunchConsumer -qname room0

test/consumer/room1:
	cd ~/go/src/github.com/enbis/learning-rabbitmq/fanout-exchange/consumer \
	&& $(GOTEST) -run TestLaunchConsumer -qname room1