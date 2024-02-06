 # Variables
 DOCKER-COMPOSE=docker-compose

# Targets
.PHONY: build run stop

build:
	$(DOCKER-COMPOSE) build 

run:
	$(DOCKER-COMPOSE) up -d
	@npx tailwindcss -i ./css/input.css -o ./css/output.css --watch

stop:
	@$(DOCKER-COMPOSE) down

