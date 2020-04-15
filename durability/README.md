# durability

* Queue: needs to be configured as durable in order to survive a broker restart. Queue durable means the queue will survive, not the messages on it. To do that, set messages persisten.
* Producer: set message delivery mode to persistent `DeliveryMode: amqp.Persistent,`

Launch RabbitMQ broker `docker run --name rabbit_dev -p 5672:5672 -p 15672:15672 rabbitmq:3-management`

Launch producer_test `go test -v`, it will create some messages to forward to the Consumer. Open RabbitMQ Management, look at Queues tab. You will find the durable queue, with some messages on it waiting for che Consumer. Restart the docker container `docker restart rabbit_dev`. Back to the Management Queues, the queue it will be still there with the messages ready to be deliverd on it. The messages will be on hold until the Consumer processes them and send back the ack.
