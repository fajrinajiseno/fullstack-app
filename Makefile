DOCKER_COMPOSE = docker-compose

run-docker:
	$(DOCKER_COMPOSE) up --build

run-docker-detach:
	$(DOCKER_COMPOSE) up --build -d

stop-docker:
	$(DOCKER_COMPOSE) down -v
