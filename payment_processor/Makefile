.PHONY: docker_up_app docker_up_deps docker_down migration_create migration_up mockgen app_run app_build

############################### DOCKER ###############################
docker_down:
	docker-compose down

docker_up_app: docker_down
	docker-compose up --build --force-recreate

docker_up_deps:
	docker-compose up --build postgres
	
############################### MIGRATE ###############################
migration_create:
	migrate create -ext sql -dir adapter/repository/migration -seq $(NAME)

migration_up:
	migrate -path adapter/repository/migration -database "sqlite3://transaction.db" up

############################### MOCKGEN ###############################
REPOSITORY_PATH_MOCKGEN=domain/repository/mock
ADAPTER_BROKER_PATH=adapter/broker/*.go
ADAPTER_BROKER_PATH_MOCKGEN=adapter/broker/mock

mockgen:
	mockgen -source=domain/repository/transaction_repository.go -destination=domain/repository/mock/transaction_repository.go ; \
	mockgen -source=adapter/broker/producer.go -destination=adapter/broker/mock/producer.go ; \

############################### APP ###############################
app_run:
	go run cmd/main.go

app_build:
	go build -o main cmd/main.go

############################### Kafka ###############################
kafka_publish:
	docker exec -it payment_base_kafka /bin/bash -c "kafka-console-producer --broker-list localhost:9092 --topic transactions --property parse.key=true --property key.separator=: < ./fixtures/success-transaction.txt"