# fair-dispatch

This exercise serves to compare two type of message dispatch.
To proceed with the comparison, two tests have been created: fair and unfair.

## unfair solution

* Terminal 0 -> Producer `go test -run TestPublishUnfair`
* Terminal 1 -> Consumer#0 `go test -run TestConsumeMsgUnfair`
* Terminal 2 -> Consumer#1 `go test -run TestConsumeMsgUnfair`

In the unfair solution the Producer send ten messages on the queue called unfair. Respectively odd with 3 seconds and even with 8 seconds to process the job effort. The two Consumers start porcessing it without the Qos value setted on the broker channel. That's what you can see as output on the terminal of both Consumer. One Consumer is constantly busy, the other has a lot less work to do. 

Terminal 1 -> first Consumer
```
Connection string  amqp://guest:guest@localhost:5672
Consuming unfair
message received and processed on Consumer: Message 1 #3 
message received and processed on Consumer: Message 3 #3 
message received and processed on Consumer: Message 5 #3 
message received and processed on Consumer: Message 7 #3 
message received and processed on Consumer: Message 9 #3
```

Terminal 2 -> second Consumer
```
Connection string  amqp://guest:guest@localhost:5672
Consuming unfair
message received and processed on Consumer: Message 0 #8 
message received and processed on Consumer: Message 2 #8 
message received and processed on Consumer: Message 4 #8 
message received and processed on Consumer: Message 6 #8 
message received and processed on Consumer: Message 8 #8 
```

What happens if we set Qos on the broker channel when the Consumer established the connection?

## fair solution

* Terminal 0 -> Producer `go test -run TestPublishFair`
* Terminal 1 -> Consumer#0 `go test -run TestConsumeMsgFair`
* Terminal 2 -> Consumer#1 `go test -run TestConsumeMsgFair`

In the unfair solution the Producer send ten messages on the queue called fair. Respectively odd with 3 seconds and even with 8 seconds to process the job effort, as well as the first solution but using another queue. The two Consumers start porcessing it with the Qos value setted on the broker channel, with the prefetch value to 1. That necessary telling the broker not to give more than one message to a Consumer at a time, don't dispatch a new message to a worker until it has processed the ack of the previous message.

That's what you can see as output on the terminal of both Consumer. The work between the two Consumers is divided much more equitably.

Terminal 1 -> first Consumer
```
Connection string  amqp://guest:guest@localhost:5672
Consuming fair
message received and processed on Consumer: Message 0 #8 
message received and processed on Consumer: Message 3 #3 
message received and processed on Consumer: Message 4 #8 
message received and processed on Consumer: Message 7 #3 
message received and processed on Consumer: Message 8 #8 
```

Terminal 2 -> second Consumer
```
Connection string  amqp://guest:guest@localhost:5672
Consuming fair
message received and processed on Consumer: Message 1 #3 
message received and processed on Consumer: Message 2 #8 
message received and processed on Consumer: Message 5 #3 
message received and processed on Consumer: Message 6 #8 
message received and processed on Consumer: Message 9 #3
```