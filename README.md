# level2practices

# 1. Run RabbitMQ on 5672 port.

```sudo docker-compose up -d rabbitmq```

>*if you have error: "docker: Error response from daemon: driver failed programming external connectivity on endpoint rabbitmq (7d9bcc13b9d657b8b27893e64f1b851a998fa403a2324d2ab39ad52e7d08c8bd): Error starting userland proxy: listen tcp4 0.0.0.0:15672: bind: address already in use."*

>Try:

>1.1. ```sudo lsof -i tcp:5672```

>1.2. ```sudo kill -9 [insert PID rabbitmq container from 1.1.]```

# 2. Run MongoDB.

```sudo docker-compose up -d mongodb```

# 3. Run Producer.

```sudo docker-compose up producer```

# 4. Run Consumer.

```sudo docker-compose up consumer```


