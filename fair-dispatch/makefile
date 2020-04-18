GOCMD=go
GOTEST=$(GOCMD) test

test/producer/fair:
	cd ~/go/src/github.com/enbis/learning-rabbitmq/fair-dispatch/producer \
	&& $(GOTEST) -run TestPublishFair

test/consumer/fair:
	cd ~/go/src/github.com/enbis/learning-rabbitmq/fair-dispatch/consumer \
	&& $(GOTEST) -run TestConsumeMsgFair

test/producer/unfair:
	cd ~/go/src/github.com/enbis/learning-rabbitmq/fair-dispatch/producer \
	&& $(GOTEST) -run TestPublishUnfair

test/consumer/unfair:
	cd ~/go/src/github.com/enbis/learning-rabbitmq/fair-dispatch/consumer \
	&& $(GOTEST) -run TestConsumeMsgUnfair