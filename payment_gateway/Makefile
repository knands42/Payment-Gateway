.PHONY: docker_down docker_up_app docker_up_deps app_run app_build

NODE_ENV ?= default

############################### Docker ###############################
docker_down:
	docker-compose down

docker_up_app: docker_down
	docker-compose up --build --force-recreate

docker_up_deps:
	docker-compose up --build mysql
	
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
	npx sequelize-cli seed:generate --name $(NAME) --config ./dist/config/database/sequelize.config.js --seeders-path ./src/config/database/seeders --env $(NODE_ENV)

seeder_up:
	npx sequelize-cli db:seed:all --config ./dist/config/database/sequelize.config.js --seeders-path ./src/config/database/seeders --env $(NODE_ENV)

seeder_down:
	npx sequelize-cli db:seed:undo:all --config ./dist/config/database/sequelize.config.js --seeders-path ./src/config/database/seeders --env $(NODE_ENV)

############################### App ###############################
app_run:
	npm run start:dev

app_build:
	npm run build
