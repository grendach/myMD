# Kafka commands:

* delete all unused data in docker
    ```
    $ docker system prune -a
    ```

* start kafka docker container
    ```
    $ sudo docker-compose up
    ```

* scale kafka 
    ```
    $ sudo docker-compose scale kafka=2 
    ```

* connecting to kafka console 
    ```
    $ sudo ./start-kafka-shell.sh 172.17.0.1 172.17.0.1:2181 
    ```
*  starting the topic on kafka with 2 partitions and 2 replication-factor
    ```
    $ KAFKA_HOME/bin/kafka-topics.sh --create --topic testtopic1 --partitions 2 --zookeeper $ZK --replication-factor 2 
    ```
* starting the topic on kafka with 1 partitions and 1 replication-factor
    ```
    $KAFKA_HOME/bin/kafka-topics.sh --create --topic kafkatestteam --partitions 1 --zookeeper $ZK --replication-factor 1 
    ```

*  posting on created topic a new posts as producer
    ```
    $ KAFKA_HOME/bin/kafka-console-producer.sh --topic=kafkatestteam --broker-list=`broker-list.sh` 
    ```
* reading a meesages from the topic as consumer
    ```
    $ KAFKA_HOME/bin/kafka-console-consumer.sh --topic=kafkatestteam --zookeeper=$ZK 
    ```

* check list of all kafka topics
    ```
    $ KAFKA_HOME/bin/kafka-topics.sh --list --zookeeper 172.17.0.1:2181
    ```

* check list of all entries for a topic 
    ```
    $ KAFKA_HOME/bin/kafka-console-consumer.sh --zookeeper $ZK --topic kafkatestteam --group my_grpoup --from-beginning 
    ```

* delete topic
    ```
    $KAFKA_HOME/bin/kafka-topics.sh --zookeeper 172.17.0.1:2181 --delete --topic topic
    ```
* list of kafka groups
    ```
    $ KAFKA_HOME/bin/kafka-consumer-groups.sh --zookeeper $ZK --list
    ```