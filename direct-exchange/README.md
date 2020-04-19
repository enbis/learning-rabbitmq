# direct-exchange

Receiving data selectively. The routing algorithm behind a direct exchange is simple: a message goes to the queues whose binding key exactly matches the routing key of the message. On the ohter side it can't do routing based on multiple criteria.

* A message queue binds to the exchange using a routing key, K.
* A publisher sends the Exchange a message with the routing key, R.
* The message is passed to the message queue if K equals to R.

## the solution

* Terminal 0 -> `make test/consumer/room0` -> It runs the Consumer#0, binding with all_rooms and first_room routing key. 
* Terminal 1 -> `make test/consumer/room1` -> It runs the Consumer#1, binding with all_rooms and second_room routing key 
* Terminal 2 -> `make test/producer` -> It runs the Producer, sending 3 messages respectively Bulb On, Bulb Off and Bulb Off. Each one destined to the three different routing key as shown below.

*Terminal 2 -> The Producer*
```
Connection string  amqp://guest:guest@localhost:5672

0: Light bulb On to routingKey all_rooms
1: Light bulb Off to routingKey first_room
2: Light bulb Off to routingKey second_room
```

*Terminal 0 -> The Consumer#0*
```
Connection string  amqp://guest:guest@localhost:5672

Binding to  all_rooms
Binding to  first_room
Waiting messages

message received: Light bulb On
message received: Light bulb Off
```

*Terminal 1 -> The Consumer#1*
```
Connection string  amqp://guest:guest@localhost:5672

Binding to all_rooms
Binding to second_room
Waiting messages

message received: Light bulb On
message received: Light bulb Off
```
