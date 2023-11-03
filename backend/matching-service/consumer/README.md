# Matching Consumer

## Setting up environment

1. Create a `.env` file referencing from the `.env.sample` in the current directory

2. Here's an example of the `.env` file:

```
AMQP_SERVER_URL=amqp://guest:guest@localhost:5672/
COLLAB_URL=http://collaboration-service:8080
RMQ_QUEUE_URL=http://rabbitmq:15672/api/queues/%2f/
```
