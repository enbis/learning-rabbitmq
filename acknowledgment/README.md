# Message acknowledgment

* Producer, it produces the messages.
* Queue, acts like an interface between Producer and Consumer.
* Consumers, more than one with the purpose of process messages.

Working with the acknowledge (ack) in order to make sure a message is never lost. An ack is sent back by the Consumer after receiving and processing each message. The ack is used to RabbitMQ to delete the message from the queue.
If a Consumer dies without sending back the ack, RabbitMQ understand that the message wasn't fully processed and will re-queue it in order to redelive it to another available Consumer.
To enable manual ack, set false for `auto-ack` argument of the Queue and add proper method `message.Ack(false)` once the task done.

## How to test

I've prepared standard testing package for the two processes to run: consumer and producer.
From within each folder, launch `go test -v`. Do twice for the Consumer to open two channels, and once for the Producer (there is no needs to open more than one Producer). The RabbitMQ will forward the messages to the both Consumer following round-robin technique. During the procedure, just kill one Consumer and you will immediatley recognize that all the messages goes to the only Consumer alive.
At the end of the process all the 20 messages will be sent, some towards the first Consumer for as long as it was active and all the rest to the second Consumer. No messages will be lost, this thanks to the ack feature.
