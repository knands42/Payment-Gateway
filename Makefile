.PHONY: migration_fixture_create migration_fixture_up mockgen

############################### DOCKER ###############################
docker_up:
	docker-compose up --build 
	
############################### MIGRATE ###############################
migration_create:
	migrate create -ext sql -dir adapter/repository/migration -seq $(NAME)

migration_up:
	migrate -path adapter/repository/migration -database "sqlite3://transaction.db" up

############################### MOCKGEN ###############################
REPOSITORY_PATH = domain/repository/*.go
REPOSITORY_PATH_MOCKGEN = domain/repository/mock
ADAPTER_BROKER_PATH = adapter/broker/*.go
ADAPTER_BROKER_PATH_MOCKGEN = adapter/broker/mock

mockgen:
	@for file in $(REPOSITORY_PATH); do \
		mockgen -source=$$file -destination=$(REPOSITORY_PATH_MOCKGEN)/`basename $$file` ; \
	done
	@for file in $(ADAPTER_BROKER_PATH); do \
		mockgen -source=$$file -destination=$(ADAPTER_BROKER_PATH_MOCKGEN)/`basename $$file` ; \
	done

############################### APP ###############################
app_run:
	go run cmd/main.go

app_build:
	go build -o main cmd/main.go
