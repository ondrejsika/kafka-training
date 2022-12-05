[Ondrej Sika (sika.io)](https://sika.io) | <ondrej@sika.io>

# Kafka Training

## Course

## Kafka Theoretical Introduction

### Kafka Components

- https://www.conduktor.io/kafka/kafka-fundamentals

![](./images/kafka_components.webp)

### Kafka Topics

- https://www.conduktor.io/kafka/kafka-topics

![](./images/kafka_topics.webp)

![](./images/kafka_topic_partitions.webp)

![](./images/kafka_topic_example.webp)

### Kafka Producers

- https://www.conduktor.io/kafka/kafka-producers

![](./images/kafka_producers.webp)

### Kafka Consumers

- https://www.conduktor.io/kafka/kafka-consumers

![](./images/kafka_consumers.webp)

### Kafka Brokers

- https://www.conduktor.io/kafka/kafka-brokers

![](./images/kafka_brokers.webp)

![](./images/kafka_brokers_partitions.webp)

### Kafka Topic Replications

- https://www.conduktor.io/kafka/kafka-topic-replication

![](./images/kafka_topic_replication.webp)

## Strimzi

> Kafka on Kubernetes in a few minutes

- https://strimzi.io/
- https://strimzi.io/quickstarts/

## Install Strimzi

```
kubectl create namespace kafka
```

```
kubectl create -f 'https://strimzi.io/install/latest?namespace=kafka' -n kafka
```

See all Strimzi API resources (CRDs)

```
kubectl api-resources | grep strimzi.io
```

See operator's pod:

```
kubectl get pod -n kafka
```

or watch

```
watch -n 0.3 kubectl get pod -n kafka
```

## Create Kafka Cluster using Strimzi

Apply:

```
kubectl apply -f examples/strimzi/kafka-1.yml
```

See pods:

```
kubectl get pod -n kafka
```

or watch

```
watch -n 0.3 kubectl get pod -n kafka
```

Get kafkas

```
kubectl get kafka
```

```
kubectl get k
```

Get Bootstrap servers

```
kubectl describe k kafka-1 | grep "Bootstrap Servers"
```

Create topic

```
kubectl apply -f examples/strimzi/kafka-1-topic-1.yml
```

Get topic

```
kubectl get -f examples/strimzi/kafka-1-topic-1.yml
```

Watch topic

```
watch -n 0.3 kubectl get -f examples/strimzi/kafka-1-topic-1.yml
```

## Thank you! & Questions?

That's it. Do you have any questions? **Let's go for a beer!**

### Ondrej Sika

- email: <ondrej@sika.io>
- web: <https://sika.io>
- twitter: [@ondrejsika](https://twitter.com/ondrejsika)
- linkedin: [/in/ondrejsika/](https://linkedin.com/in/ondrejsika/)
- Newsletter, Slack, Facebook & Linkedin Groups: <https://join.sika.io>

_Do you like the course? Write me recommendation on Twitter (with handle `@ondrejsika`) and LinkedIn (add me [/in/ondrejsika](https://www.linkedin.com/in/ondrejsika/) and I'll send you request for recommendation). **Thanks**._

Wanna to go for a beer or do some work together? Just [book me](https://book-me.sika.io) :)
