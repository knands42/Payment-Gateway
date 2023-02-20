############################### Docker ###############################
docker_down:
	docker-compose -f ./payment_broker/docker-compose.yaml down
	docker-compose -f ./payment_gateway/docker-compose.yaml down
	docker-compose -f ./payment_processor/docker-compose.yaml down

docker_up_apps: docker_down
	docker-compose -f ./payment_broker/docker-compose.yaml up --build --force-recreate -d
	docker-compose -f ./payment_gateway/docker-compose.yaml up --build --force-recreate -d
	docker-compose -f ./payment_processor/docker-compose.yaml up --build --force-recreate -d

docker_up_broker: docker_down
	docker-compose -f ./payment_broker/docker-compose.yaml up --build --force-recreate