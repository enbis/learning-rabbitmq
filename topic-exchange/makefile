GOCMD=go
GOTEST=$(GOCMD) test
GOFLAGS :=

test/producer:
	cd ~/go/src/github.com/enbis/learning-rabbitmq/topic-exchange/producer \
	&& $(GOTEST) -run TestLaunchPublisher

test/consumer/lights:
	cd ~/go/src/github.com/enbis/learning-rabbitmq/topic-exchange/consumer \
	&& $(GOTEST) -run TestLaunchConsumer -qbinding '#.light' -qname queue4lights

test/consumer/garage:
	cd ~/go/src/github.com/enbis/learning-rabbitmq/topic-exchange/consumer \
	&& $(GOTEST) -run TestLaunchConsumer -qbinding '*.garage.*' -qname queue4garage	