############################### Docker ###############################
docker_down:
	docker-compose -f docker-compose.yaml down || true
	docker-compose -f ./payment_gateway/docker-compose.yaml down || true
	docker-compose -f ./payment_processor/docker-compose.yaml down || true
	docker network rm payment_gateway_network || true

docker_up_apps: docker_down
	docker network create payment_gateway_network 
	docker-compose -f docker-compose.yaml up -d
	docker-compose -f ./payment_gateway/docker-compose.yaml up --build --force-recreate -d
	docker-compose -f ./payment_processor/docker-compose.yaml up --build --force-recreate -d

############################### Kafka ###############################
kafka_publish:
	docker exec -it payment_base_kafka /bin/bash -c "kafka-console-producer --broker-list localhost:9092 --topic transactions --property parse.key=true --property key.separator=: < ./success-transaction.txt"
