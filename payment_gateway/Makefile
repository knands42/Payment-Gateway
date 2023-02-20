.PHONY: docker_down docker_up_app docker_up_deps migration_create migration_up migration_down migration_down_all seeder_create seeder_up seeder_down app_run app_build

NODE_ENV ?= local

############################### Docker ###############################
docker_down:
	docker-compose down

docker_up_app: docker_down
	docker-compose up --build --force-recreate

docker_up_deps: docker_down
	docker-compose up --build postgres -d
	docker-compose up --build zookeeper -d
	docker-compose up --build kafka-topics-generator -d
	docker-compose up --build control-center -d
	
############################### Migration ###############################
migration_create:
	npx sequelize-cli migration:generate --name $(NAME) --config ./dist/config/database/sequelize.config.js --migrations-path ./src/config/database/migrations --env $(NODE_ENV)

migration_up:
	npx sequelize-cli db:migrate --config ./dist/config/database/sequelize.config.js --migrations-path ./src/config/database/migrations --env $(NODE_ENV)

migration_down:
	npx sequelize-cli db:migrate:undo --config ./dist/config/database/sequelize.config.js --migrations-path ./src/config/database/migrations --env $(NODE_ENV)

migration_down_all:
	npx sequelize-cli db:migrate:undo:all --config ./dist/config/database/sequelize.config.js --migrations-path ./src/config/database/migrations --env $(NODE_ENV)

############################### Seeder ###############################
seeder_create:
	npx sequelize-cli seed:generate --name $(NAME) --config ./dist/config/database/sequelize.config.js --seeders-path ./src/config/database/seeders --env local

seeder_up:
	npx sequelize-cli db:seed:all --config ./dist/config/database/sequelize.config.js --seeders-path ./src/config/database/seeders --env local

seeder_down:
	npx sequelize-cli db:seed:undo:all --config ./dist/config/database/sequelize.config.js --seeders-path ./src/config/database/seeders --env local

############################### App ###############################
app_run:
	npm run start:dev

app_build:
	npm run build
