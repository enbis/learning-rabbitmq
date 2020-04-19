# direct-exchange

The routing algorithm behind a direct exchange is simple: a message goes to the queues whose binding key exactly matches the routing key of the message.

* A message queue binds to the exchange using a routing key, K.
* A publisher sends the Exchange a message with the routing key, R.
* The message is passed to the message queue if K equals to R.

## the solution

* Terminal 0 -> `make test/consumer0` -> It runs the Consumer#0, binding with all and first routing key. 
* Terminal 1 -> `make test/consumer1` -> It runs the Consumer#1, binding with all and second routing key 
* Terminal 2 -> `make test/producer` -> It runs the Producer, sending 3 messages with random string. Each one destined to the three different routing key as shown below.

*Terminal 2 -> The Producer*
```
Connection string  amqp://guest:guest@localhost:5672

0: value Message 0 RädJrnCY to routingKey all 
1: value Message 1 GxCåWssr to routingKey first 
2: value Message 2 ljAGoBzK to routingKey second
```

*Terminal 0 -> The Consumer#0*
```
Connection string  amqp://guest:guest@localhost:5672

Binding to all and first routing keys
Waiting for messages

message received: Message 0 RädJrnCY
message received: Message 1 GxCåWssr
```

*Terminal 1 -> The Consumer#1*
```
Connection string  amqp://guest:guest@localhost:5672

Binding to all and second routing keys
Waiting for messages

message received: Message 0 RädJrnCY
message received: Message 2 ljAGoBzK
```
