############################### Docker ###############################
docker_down:
	docker-compose -f docker-compose.yaml down
	docker-compose -f ./payment_gateway/docker-compose.yaml down
	docker-compose -f ./payment_processor/docker-compose.yaml down
	docker network rm payment_gateway_network || true

docker_up_apps: docker_down
	docker network create payment_gateway_network 
	docker-compose -f docker-compose.yaml up -d
	docker-compose -f ./payment_gateway/docker-compose.yaml up --build --force-recreate -d
	docker-compose -f ./payment_processor/docker-compose.yaml up --build --force-recreate -d

docker_up_broker: docker_down
	docker-compose -f ./payment_broker/docker-compose.yaml up --build --force-recreate