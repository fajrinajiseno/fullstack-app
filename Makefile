DOCKER_COMPOSE = docker-compose

run-docker:
	$(DOCKER_COMPOSE) up --build

run-docker-detach:
	$(DOCKER_COMPOSE) up --build -d

run-backend-docker:
	$(DOCKER_COMPOSE) up backend --build

run-backend-docker-detach:
	$(DOCKER_COMPOSE) up backend --build -d

run-frontend-docker:
	$(DOCKER_COMPOSE) up frontend --build

run-frontend-docker-detach:
	$(DOCKER_COMPOSE) up frontend --build -d

stop-backend-docker:
	$(DOCKER_COMPOSE) stop backend

stop-frontend-docker:
	$(DOCKER_COMPOSE) stop frontend
