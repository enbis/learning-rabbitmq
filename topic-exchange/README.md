# topic-exchange

Messages sent to a topic exchange have a routing_key composed by a list of words, delimited by dots. The words usually specify some features connected to the message. The logic behind the topic exchange is similar to a direct one, a message sent with a particular routing key will be delivered to all queues that are bound with a mataching binding key. There are two special cases for binding keys:

* `*` star substitutes exactly one word
* `#` hash substitutes for zero or more words. 

## the solution

* Terminal 0 -> `make test/consumer/lights` -> It runs the Consumer#0, that handles this routing key: `#.light`.
* Terminal 1 -> `make test/consumer/garage` -> It runs the Consumer#1, that handles this routing key: `*.garage.*`
* Terminal 2 -> `make test/producer` -> It runs the Producer, sending 5 messages.

### messages

house.room1.light -> direct to consumer/lights only
house.garage.light -> direct to both consumers
house.garage.door -> direct to consumer/garage only
house.backyard.irrigation -> neither of two consumers receive the the message
house.garage.light.desktopLamp -> neither of two consumers receive the the message

*Terminal 2 -> The Producer*
```
Connection string  amqp://guest:guest@localhost:5672

0: value 0 Action: On to routingKey house.room1.light
1: value 1 Action: On to routingKey house.garage.light
2: value 2 Action: Open to routingKey house.garage.door
3: value 3 Action: On to routingKey house.backyard.irrigation
4: value 4 Action: On to routingKey house.garage.light.desktopLamp
```

*Terminal 0 -> The Consumer#0*
```
qbinding  #.light
Connection string  amqp://guest:guest@localhost:5672

Waiting messages

Where: room1
What: light
0 Action: On
----------
Where: garage
What: light
1 Action: On
----------
```

*Terminal 1 -> The Consumer#1*
```
qbinding  *.garage.*
Connection string  amqp://guest:guest@localhost:5672

Waiting messages

Where: garage
What: light
1 Action: On
----------
Where: garage
What: door
2 Action: Open
----------
```
