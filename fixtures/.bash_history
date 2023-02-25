kafka-topics --list --zookeeper payment_base_zookeeper:2181
kafka-topics --list 
kafka-topics --list --bootstrap-server=payment_base:9092
kafka-topics --list --bootstrap-server=payment_base:9092
kafka-topics --list --bootstrap-server=localhost:9092
kafka-topics --bootstrap-server=localhost:9092 --list | grep transactions
bash -c "kafka-topics --bootstrap-server=localhost:9092 --list | grep transactions"
