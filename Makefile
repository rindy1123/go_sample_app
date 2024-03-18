API_CONTAINER_NAME := $(shell docker ps | grep api | cut -d ' ' -f1)
DB_USER := $(shell docker exec $(API_CONTAINER_NAME) sh -c 'echo $$DB_USER')
DB_PASSWORD := $(shell docker exec $(API_CONTAINER_NAME) sh -c 'echo $$DB_PASSWORD')
DB_NAME := $(shell docker exec $(API_CONTAINER_NAME) sh -c 'echo $$DB_NAME')
DB_PORT := $(shell docker exec $(API_CONTAINER_NAME) sh -c 'echo $$DB_PORT')
DB_HOST := $(shell docker exec $(API_CONTAINER_NAME) sh -c 'echo $$DB_HOST')
DB_SSLMODE := $(shell docker exec $(API_CONTAINER_NAME) sh -c 'echo $$DB_SSLMODE')

migrate_diff:
	docker exec $(API_CONTAINER_NAME) atlas migrate diff --env gorm

migrate_apply:
	docker exec $(API_CONTAINER_NAME) atlas migrate apply --url postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE) --dir file://migrations 
