# fanout-exchange

* A message queue binds to the exchange with no arguments.
* The message sends by the Producer to the Exchange, is passed to the queues unconditionally.
* The fanout exchange broadcasts all the messages it receives to all the queues it knows.

