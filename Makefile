.PHONY: migration_fixture_create migration_fixture_up mockgen

include app.env

############################### DOCKER ###############################
docker_up:
	docker-compose up --build 
	
############################### MIGRATE ###############################
migration_fixture_create:
	migrate create -ext sql -dir adapter/repository/fixture/migration -seq $(NAME)

migration_fixture_up:
	migrate -path adapter/repository/fixture/migration -database "$(DB_SOURCE)" -verbose up

############################### MOCKGEN ###############################
REPOSITORY_PATH = domain/repository/*.go
REPOSITORY_PATH_MOCKGEN = domain/repository/mock

mockgen:
	@for file in $(REPOSITORY_PATH); do \
		mockgen -source=$$file -destination=$(REPOSITORY_PATH_MOCKGEN)/`basename $$file` ; \
	done