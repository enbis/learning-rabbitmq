# fanout-exchange

 Fanout exchange is only capable of broadcasting. The fanout exchanges ignored the routing-key value of a binding.

* A message queue binds to the exchange with no arguments.
* The message sends by the Producer to the Exchange, is passed to the queues unconditionally.
* The fanout exchange broadcasts all the messages it receives to all the queues it knows.

## the solution

* Terminal 0 -> `make test/consumer/room0` -> It runs the Consumer#0, as room0.
* Terminal 1 -> `make test/consumer/room1` -> It runs the Consumer#1, as room1.
* Terminal 2 -> `make test/producer` -> It runs the Producer, that sends messages to switch on / off the light bulb on the rooms.

*Terminal 2 -> The Producer*
```
Connection string  amqp://guest:guest@localhost:5672

0: Light bulb On 
1: Light bulb Off 
2: Light bulb On 
3: Light bulb Off 
4: Light bulb On
```

*Terminal 0 -> The Consumer#0 as a Room0*
```
Connection string  amqp://guest:guest@localhost:5672

Waiting messages
message received: Light bulb On
message received: Light bulb Off
message received: Light bulb On
message received: Light bulb Off
message received: Light bulb On
```

*Terminal 1 -> The Consumer#1 as a Room1*
```
Connection string  amqp://guest:guest@localhost:5672

Waiting messages
message received: Light bulb On
message received: Light bulb Off
message received: Light bulb On
message received: Light bulb Off
message received: Light bulb On
```