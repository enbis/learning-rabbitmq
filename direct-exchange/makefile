GOCMD=go
GOTEST=$(GOCMD) test
GOFLAGS :=

test/producer:
	cd ~/go/src/github.com/enbis/learning-rabbitmq/direct-exchange/producer \
	&& $(GOTEST) -run TestLaunchPublisher

test/consumer0:
	cd ~/go/src/github.com/enbis/learning-rabbitmq/direct-exchange/consumer \
	&& $(GOTEST) -run TestLaunchConsumer -qbinding 0 

test/consumer1:
	cd ~/go/src/github.com/enbis/learning-rabbitmq/direct-exchange/consumer \
	&& $(GOTEST) -run TestLaunchConsumer -qbinding 1	