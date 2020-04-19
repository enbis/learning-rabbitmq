# fanout-exchange

 Fanout exchange is only capable of broadcasting. The fanout exchanges ignored the routing-key value of a binding.

* A message queue binds to the exchange with no arguments.
* The message sends by the Producer to the Exchange, is passed to the queues unconditionally.
* The fanout exchange broadcasts all the messages it receives to all the queues it knows.

## the solution

* Terminal 0 -> `make test/consumer` -> It runs the Consumer#0. 
* Terminal 1 -> `make test/consumer` -> It runs the Consumer#1. 
* Terminal 2 -> `make test/producer` -> It runs the Producer.

*Terminal 2 -> The Producer*
```
Connection string  amqp://guest:guest@localhost:5672

0: Message 0 KHMsxaRO 
1: Message 1 äFÄmHHyo 
2: Message 2 fnzhNJPh 
3: Message 3 kclPötäp 
4: Message 4 xvjxkZub
```

*Terminal 0 -> The Consumer#0*
```
Connection string  amqp://guest:guest@localhost:5672

Waiting messages
message received: Message 0 KHMsxaRO
message received: Message 1 äFÄmHHyo
message received: Message 2 fnzhNJPh
message received: Message 3 kclPötäp
message received: Message 4 xvjxkZub
```

*Terminal 1 -> The Consumer#1*
```
Connection string  amqp://guest:guest@localhost:5672

Waiting messages
message received: Message 0 KHMsxaRO
message received: Message 1 äFÄmHHyo
message received: Message 2 fnzhNJPh
message received: Message 3 kclPötäp
message received: Message 4 xvjxkZub
```