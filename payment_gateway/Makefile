.PHONY: docker_down docker_up_app docker_up_deps app_run app_build

############################### DOCKER ###############################
docker_down:
	docker-compose down

docker_up_app: docker_down
	docker-compose up --build --force-recreate

docker_up_deps:
	docker-compose up --build mysql
	
############################### APP ###############################
app_run:
	npm run start:dev

app_build:
	npm run build
