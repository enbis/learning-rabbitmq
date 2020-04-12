# start-queuing

* Producer, the same like `start-messaging`.
* Queue, acts like an interface between Producer and Consumer. This time will be used to distribute time-consuming tasks among multiple Consumers. 
* Consumers, more than one. 

We simulate a task through sending inside a string the time in seconds the program has to wait `time.sleep` before finish the job. Purpose of the program is understand how parallelize works.

You will notice that with two Consumer that both waiting for messages, RabbitMQ will send each message in sequence to the next Consumer. So basically RabbitMQ divides the total number of messages into equal parts to each Consumer. This is round-robin dealing technique. 

